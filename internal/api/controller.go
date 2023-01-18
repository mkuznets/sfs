package api

import (
	"context"
	"fmt"
	"io"
	"mkuznets.com/go/sps/internal/auth"
	"mkuznets.com/go/sps/internal/feed"
	"mkuznets.com/go/sps/internal/files"
	"mkuznets.com/go/sps/internal/store"
	"mkuznets.com/go/sps/internal/ytils/ycrypto"
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
	CreateUser(ctx context.Context) (*CreateUserResponse, error)
	Login(ctx context.Context, req *LoginRequest) (string, error)
}

type controllerImpl struct {
	fileStorage    files.Storage
	store          store.Store
	idService      IdService
	feedController feed.Controller
	authService    auth.Service
}

func NewController(store store.Store, fileStorage files.Storage, idService IdService, feedController feed.Controller, authService auth.Service) Controller {
	return &controllerImpl{
		store:          store,
		fileStorage:    fileStorage,
		idService:      idService,
		feedController: feedController,
		authService:    authService,
	}
}

// GetChannel returns the channel response for the given ID.
//
//	@ID			GetChannel
//	@Summary	Get channel by ID
//	@Tags		Channels
//	@Produce	json
//	@Param		id	path		string	true	"Channel ID"
//	@Success	200	{object}	ChannelResponse
//	@Failure	400	{object}	ErrorResponse
//	@Failure	404	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/channels/{id} [get]
//	@Security	Authentication
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

// CreateChannel creates a new channel with the given parameters and returns a response with the channel ID.
//
//	@ID			CreateChannel
//	@Summary	Create a new channel
//	@Tags		Channels
//	@Accept		json
//	@Produce	json
//	@Param		request	body		CreateChannelRequest	true	"CreateChannel request"
//	@Success	200		{object}	IdResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/channels [post]
//	@Security	Authentication
func (c *controllerImpl) CreateChannel(ctx context.Context, userId string, r CreateChannelRequest) (*IdResponse, error) {
	model := &store.Channel{
		Id:          c.idService.Channel(ctx),
		UserId:      userId,
		Title:       r.Title,
		Link:        r.Link,
		Authors:     r.Authors,
		Description: r.Description,
		CreatedAt:   ytime.Now(),
		UpdatedAt:   ytime.Now(),
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

// ListChannels returns a list of channels of the given user.
//
//	@ID			ListChannels
//	@Summary	List channels of the current user
//	@Tags		Channels
//	@Produce	json
//	@Success	200	{object}	IdResponse
//	@Failure	401	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/channels [get]
//	@Security	Authentication
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

// UploadFile uploads a new audio file and returns a response with the file ID.
//
//	@ID			UploadFile
//	@Summary	Uploads a new audio file
//	@Tags		Files
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		file	formData	file	true	"File to upload"
//	@Success	200		{object}	IdResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/files [post]
//	@Security	Authentication
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
		CreatedAt: ytime.Now(),
		UpdatedAt: ytime.Now(),
	}
	if err := c.store.CreateFile(ctx, model); err != nil {
		return nil, err
	}

	return &UploadResponse{
		Id: model.Id,
	}, nil
}

// CreateEpisode creates a new episode with the given parameters and returns a response with the new episode ID.
//
//	@ID			CreateEpisode
//	@Summary	Create a new episode
//	@Tags		Episodes
//	@Accept		json
//	@Produce	json
//	@Param		id		path		string					true	"Channel ID"
//	@Param		request	body		CreateEpisodeRequest	true	"CreateEpisode request"
//	@Success	200		{object}	IdResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	404		{object}	ErrorResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/channels/{id}/episodes [post]
//	@Security	Authentication
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
		CreatedAt:   ytime.Now(),
		UpdatedAt:   ytime.Now(),
	}

	if err := c.store.CreateEpisode(ctx, model); err != nil {
		return nil, err
	}

	return &IdResponse{
		Id: model.Id,
	}, nil
}

// ListEpisodes returns a list of episodes of the given channel.
//
//	@ID			ListEpisodes
//	@Summary	List episoded of the given channel
//	@Tags		Episodes
//	@Produce	json
//	@Param		id		path		string					true	"Channel ID"
//	@Success	200	{object}	IdResponse
//	@Failure	404	{object}	ErrorResponse
//	@Failure	401	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/channels/{id}/episodes [get]
//	@Security	Authentication
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

// CreateUser registers a new user and returns a response with the new account number.
//
//	@ID			CreateUser
//	@Summary	Register a new user
//	@Tags		Users
//	@Accept		json
//	@Produce	json
//	@Success	200		{object}	IdResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/users [post]
func (c *controllerImpl) CreateUser(ctx context.Context) (*CreateUserResponse, error) {
	id := c.idService.User(ctx)

	accountNumber := auth.RandomAccountNumber()
	accountNumberHashed, err := ycrypto.HashPassword(accountNumber, nil)
	if err != nil {
		return nil, err
	}

	model := &store.User{
		Id:            id,
		AccountNumber: accountNumberHashed,
		CreatedAt:     ytime.Now(),
		UpdatedAt:     ytime.Now(),
	}

	if err := c.store.CreateUser(ctx, model); err != nil {
		return nil, err
	}

	return &CreateUserResponse{
		AccountNumber: accountNumber,
	}, nil
}

func (c *controllerImpl) Login(ctx context.Context, req *LoginRequest) (string, error) {
	user, err := c.store.GetUserByAccountNumber(ctx, req.AccountNumber)
	if err != nil {
		return "", err
	}
	return c.authService.Token(user.Id)
}
