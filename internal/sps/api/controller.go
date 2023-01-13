package api

import (
	"context"
	"io"
	"mkuznets.com/go/sps/internal/feed"
	"mkuznets.com/go/sps/internal/herror"
	"mkuznets.com/go/sps/internal/types"
	"os"
	"sort"
	"time"
)

type Controller interface {
	GetChannel(ctx context.Context, id string) (*ChannelResponse, error)
	GetFeed(ctx context.Context, userId, channelId string) (*feed.Podcast, error)
	CreateChannel(ctx context.Context, userId string, r CreateChannelRequest) (*ChannelResponse, error)
	ListChannels(ctx context.Context, userId string) ([]*ChannelResponse, error)
	UploadFile(ctx context.Context, userId string, f io.ReadSeeker) (*UploadResponse, error)
	CreateEpisode(ctx context.Context, userId, channelId string, r *CreateEpisodeRequest) (*EpisodeResponse, error)
	ListEpisodes(ctx context.Context, userId, channelId string) ([]*EpisodeResponse, error)
}

type controllerImpl struct {
	uploader Uploader
	store    Store
}

func NewController(store Store, uploader Uploader) Controller {
	return &controllerImpl{
		store:    store,
		uploader: uploader,
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
	}

	return response, nil
}

func (c *controllerImpl) GetFeed(ctx context.Context, userId, channelId string) (*feed.Podcast, error) {
	channel, err := c.store.GetChannel(ctx, channelId)
	if err != nil {
		return nil, err
	}
	if channel.UserId != userId {
		return nil, herror.NotFound("channel not found")
	}

	episodes, err := c.store.ListEpisodesWithFiles(ctx, channelId)
	if err != nil {
		return nil, err
	}
	if len(episodes) == 0 {
		return nil, herror.Validation("no episodes")
	}

	sort.Slice(episodes, func(i, j int) bool {
		return episodes[i].CreatedAt.After(episodes[j].CreatedAt.Time)
	})

	publishedAt := episodes[0].CreatedAt.Time

	podcast := &feed.Podcast{
		Version: "2.0",
		Itunes:  "http://www.itunes.com/dtds/podcast-1.0.dtd",
		Channel: &feed.Channel{
			Title:         channel.Title,
			Link:          channel.Link,
			Description:   channel.Description,
			LastBuildDate: time.Now().Format(time.RFC1123Z),
			PubDate:       publishedAt.Format(time.RFC1123Z),
			IAuthor:       channel.Authors,
		},
	}

	var items []*feed.Item
	for _, episode := range episodes {
		items = append(items, &feed.Item{
			Guid: feed.Guid{
				IsPermaLink: false,
				Text:        episode.Id,
			},
			PubDate: episode.CreatedAt.Format(time.RFC1123Z),
			Title:   episode.Title,
			Link:    episode.Link,
			Description: &feed.Description{
				Text: episode.Description,
			},
			IAuthor: episode.Authors,
			Enclosure: &feed.Enclosure{
				URL:    episode.File.Url,
				Length: episode.File.Size,
				Type:   episode.File.ContentType,
			},
		})
	}
	podcast.Channel.Items = items

	return podcast, nil
}

func (c *controllerImpl) CreateChannel(ctx context.Context, userId string, r CreateChannelRequest) (*ChannelResponse, error) {
	model := &Channel{
		Id:          RandomChannelId(),
		UserId:      userId,
		Title:       r.Title,
		Link:        r.Link,
		Authors:     r.Authors,
		Description: r.Description,
		CreatedAt:   types.NewTime(time.Now()),
		UpdatedAt:   types.NewTime(time.Now()),
	}

	if err := c.store.CreateChannel(ctx, model); err != nil {
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
	tmpFile, err := os.CreateTemp("", "tmpUpload")
	defer func() {
		_ = tmpFile.Close()
		_ = os.Remove(tmpFile.Name())
	}()

	if err != nil {
		return nil, err
	}

	info, err := c.uploader.Upload(ctx, f)
	if err != nil {
		return nil, err
	}

	model := &File{
		Id:          RandomFileId(),
		UserId:      userId,
		Url:         info.Url,
		Size:        info.Size,
		ContentType: info.ContentType,
		CreatedAt:   types.NewTime(time.Now()),
		UpdatedAt:   types.NewTime(time.Now()),
	}
	if err := c.store.CreateFile(ctx, model); err != nil {
		return nil, err
	}

	return &UploadResponse{
		Id: model.Id,
	}, nil
}

func (c *controllerImpl) CreateEpisode(ctx context.Context, userId, channelId string, r *CreateEpisodeRequest) (*EpisodeResponse, error) {
	channel, err := c.store.GetChannel(ctx, channelId)
	if err != nil {
		return nil, err
	}
	if channel.UserId != userId {
		return nil, herror.NotFound("channel not found")
	}

	file, err := c.store.GetFile(ctx, r.FileId)
	if err != nil {
		return nil, err
	}
	if file.UserId != userId {
		return nil, herror.NotFound("file not found")
	}

	model := &Episode{
		Id:          RandomEpisodeId(),
		ChannelId:   channelId,
		Title:       r.Title,
		Link:        r.Link,
		Authors:     r.Authors,
		Description: r.Description,
		FileId:      r.FileId,
		CreatedAt:   types.NewTime(time.Now()),
		UpdatedAt:   types.NewTime(time.Now()),
	}

	if err := c.store.CreateEpisode(ctx, model); err != nil {
		return nil, err
	}

	response := &EpisodeResponse{
		Id: model.Id,
		File: &FileResponse{
			Id:          file.Id,
			Url:         file.Url,
			Size:        file.Size,
			ContentType: file.ContentType,
		},
		ChannelId:   model.ChannelId,
		Title:       model.Title,
		Link:        model.Link,
		Authors:     model.Authors,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}

	return response, nil
}

func (c *controllerImpl) ListEpisodes(ctx context.Context, userId, channelId string) ([]*EpisodeResponse, error) {
	channel, err := c.store.GetChannel(ctx, channelId)
	if err != nil {
		return nil, err
	}
	if channel.UserId != userId {
		return nil, herror.NotFound("channel not found")
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
				Url:         episode.File.Url,
				Size:        episode.File.Size,
				ContentType: episode.File.ContentType,
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
