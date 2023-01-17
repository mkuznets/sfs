package api

import (
	"context"
	"fmt"
	"io"
	"mkuznets.com/go/sps/internal/feed"
	"mkuznets.com/go/sps/internal/files"
	"mkuznets.com/go/sps/internal/store"
	"mkuznets.com/go/sps/internal/ytils/yerr"
	"mkuznets.com/go/sps/internal/ytils/yrand"
	"mkuznets.com/go/sps/internal/ytils/ytime"
)

type Controller interface {
	GetChannel(ctx context.Context, id string) (*ChannelResponse, error)
	CreateChannel(ctx context.Context, userId string, r CreateChannelRequest) (*IdResponse, error)
	ListChannels(ctx context.Context, userId string) ([]*ChannelResponse, error)
	UploadFile(ctx context.Context, userId string, f io.ReadSeeker) (*UploadResponse, error)
	CreateEpisode(ctx context.Context, userId, channelId string, r *CreateEpisodeRequest) (*IdResponse, error)
	ListEpisodes(ctx context.Context, userId, channelId string) ([]*EpisodeResponse, error)
}

type controllerImpl struct {
	fileStorage    files.Storage
	store          store.Store
	idService      IdService
	feedController feed.Controller
}

func NewController(store store.Store, fileStorage files.Storage, idService IdService, feedController feed.Controller) Controller {
	return &controllerImpl{
		store:          store,
		fileStorage:    fileStorage,
		idService:      idService,
		feedController: feedController,
	}
}

// GetChannel godoc
//
//	@Summary	Get channel by ID
//	@Tags		Channels
//	@Produce	json
//	@Param		id	path		string	true	"Channel ID"
//	@Success	200	{array}		ChannelResponse
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/channels/{id} [get]
func (c *controllerImpl) GetChannel(ctx context.Context, id string) (*ChannelResponse, error) {
	model, err := c.store.GetChannel(ctx, id)
	if err != nil {
		return nil, err
	}

	response := &ChannelResponse{
		Id:          model.Id,
		Title:       model.Title,
		Link:        model.Link,
		Authors:     model.Authors,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		FeedUrl:     model.Feed.Url,
	}

	return response, nil
}

func (c *controllerImpl) CreateChannel(ctx context.Context, userId string, r CreateChannelRequest) (*IdResponse, error) {
	model := &store.Channel{
		Id:          c.idService.Channel(ctx),
		UserId:      userId,
		Title:       r.Title,
		Link:        r.Link,
		Authors:     r.Authors,
		Description: r.Description,
		CreatedAt:   ytime.NewTimeNow(),
		UpdatedAt:   ytime.NewTimeNow(),
	}

	if err := c.store.CreateChannel(ctx, model); err != nil {
		return nil, err
	}
	if err := c.feedController.Update(ctx, model.Id); err != nil {
		return nil, err
	}

	response := &IdResponse{
		Id: model.Id,
	}

	return response, nil
}

func (c *controllerImpl) ListChannels(ctx context.Context, userId string) ([]*ChannelResponse, error) {
	channels, err := c.store.ListChannels(ctx, userId)
	if err != nil {
		return nil, err
	}

	response := make([]*ChannelResponse, 0)
	for _, channel := range channels {
		response = append(response, &ChannelResponse{
			Id:          channel.Id,
			Title:       channel.Title,
			Link:        channel.Link,
			Authors:     channel.Authors,
			Description: channel.Description,
			CreatedAt:   channel.CreatedAt,
			UpdatedAt:   channel.UpdatedAt,
		})
	}

	return response, nil
}

func (c *controllerImpl) UploadFile(ctx context.Context, userId string, f io.ReadSeeker) (*UploadResponse, error) {
	info, err := files.Info(f)
	if err != nil {
		return nil, yerr.Internal("failed to get file info").WithError(err)
	}
	if info.Mime.Type != "audio" {
		return nil, yerr.Validation("unsupported file type: %s", info.Mime.Value)
	}

	fileId := c.idService.File(ctx)

	path := fmt.Sprintf("file/%s/%s/%s_%s.%s", info.Hash.Digest[:2], info.Hash.Digest[2:4], yrand.Base62(15), fileId, info.Extension)
	upload, err := c.fileStorage.Upload(ctx, path, f)
	if err != nil {
		return nil, err
	}

	model := &store.File{
		Id:        fileId,
		UserId:    userId,
		UploadUrl: upload.Url,
		UploadId:  upload.Id,
		Size:      info.Size,
		Hash:      info.Hash.String(),
		MimeType:  info.Mime.Value,
		CreatedAt: ytime.NewTimeNow(),
		UpdatedAt: ytime.NewTimeNow(),
	}
	if err := c.store.CreateFile(ctx, model); err != nil {
		return nil, err
	}

	return &UploadResponse{
		Id: model.Id,
	}, nil
}

func (c *controllerImpl) CreateEpisode(ctx context.Context, userId, channelId string, r *CreateEpisodeRequest) (*IdResponse, error) {
	channel, err := c.store.GetChannel(ctx, channelId)
	if err != nil {
		return nil, err
	}
	if channel.UserId != userId {
		return nil, yerr.NotFound("channel not found")
	}

	f, err := c.store.GetFile(ctx, r.FileId)
	if err != nil {
		return nil, err
	}
	if f.UserId != userId {
		return nil, yerr.NotFound("file not found")
	}

	model := &store.Episode{
		Id:          c.idService.Episode(ctx),
		ChannelId:   channelId,
		Title:       r.Title,
		Link:        r.Link,
		Authors:     r.Authors,
		Description: r.Description,
		FileId:      r.FileId,
		CreatedAt:   ytime.NewTimeNow(),
		UpdatedAt:   ytime.NewTimeNow(),
	}

	if err := c.store.CreateEpisode(ctx, model); err != nil {
		return nil, err
	}

	return &IdResponse{
		Id: model.Id,
	}, nil
}

func (c *controllerImpl) ListEpisodes(ctx context.Context, userId, channelId string) ([]*EpisodeResponse, error) {
	channel, err := c.store.GetChannel(ctx, channelId)
	if err != nil {
		return nil, err
	}
	if channel.UserId != userId {
		return nil, yerr.NotFound("channel not found")
	}

	episodes, err := c.store.ListEpisodesWithFiles(ctx, channelId)
	if err != nil {
		return nil, err
	}

	response := make([]*EpisodeResponse, 0)
	for _, episode := range episodes {
		response = append(response, &EpisodeResponse{
			Id: episode.Id,
			File: &FileResponse{
				Id:          episode.File.Id,
				Url:         episode.File.UploadUrl,
				Size:        episode.File.Size,
				ContentType: episode.File.MimeType,
			},
			ChannelId:   episode.ChannelId,
			Title:       episode.Title,
			Link:        episode.Link,
			Authors:     episode.Authors,
			Description: episode.Description,
			CreatedAt:   episode.CreatedAt,
			UpdatedAt:   episode.UpdatedAt,
		})
	}

	return response, nil
}
