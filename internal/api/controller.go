package api

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"mkuznets.com/go/sfs/internal/files"
	"mkuznets.com/go/sfs/internal/rss"
	"mkuznets.com/go/sfs/internal/store"
	"mkuznets.com/go/sfs/internal/user"
	"mkuznets.com/go/ytils/yslice"
	"mkuznets.com/go/ytils/ytime"
)

type Controller interface {
	GetFeeds(ctx context.Context, req *GetFeedsRequest, usr user.User) (*GetFeedsResponse, error)
	CreateFeeds(ctx context.Context, req *CreateFeedsRequest, usr user.User) (*CreateFeedsResponse, error)
	GetItems(ctx context.Context, req *GetItemsRequest, usr user.User) (*GetItemsResponse, error)
	CreateItems(ctx context.Context, req *CreateItemsRequest, usr user.User) (*CreateItemsResponse, error)
	UploadFiles(ctx context.Context, fs []multipart.File, usr user.User) (*UploadFilesResponse, error)
	GetRssUrl(ctx context.Context, feedId string) (string, error)
}

type controllerImpl struct {
	fileStorage   files.Storage
	store         store.Store
	idService     IdService
	rssController rss.Controller
}

func NewController(store store.Store, fileStorage files.Storage, idService IdService, feedController rss.Controller) Controller {
	return &controllerImpl{
		store:         store,
		fileStorage:   fileStorage,
		idService:     idService,
		rssController: feedController,
	}
}

// GetFeeds returns a list of feeds matching the given parameters.
//
//	@ID			GetFeeds
//	@Summary	Get feeds matching the given parameters
//	@Tags		Feeds
//	@Accept		json
//	@Produce	json
//	@Param		request	body		GetFeedsRequest	true	"Parameters for filtering feeds"
//	@Success	200		{object}	GetFeedsResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/feeds/get [post]
//	@Security	Authentication
func (c *controllerImpl) GetFeeds(ctx context.Context, req *GetFeedsRequest, usr user.User) (*GetFeedsResponse, error) {
	filter := store.FeedFilter{
		Ids:     req.Ids,
		UserIds: []string{usr.Id()},
	}

	feeds, err := c.store.GetFeeds(ctx, &filter)
	if err != nil {
		return nil, fmt.Errorf("HTTP 500: get feeds: %w", err)
	}

	results := yslice.Map(feeds, func(f *store.Feed) *FeedResource {
		return &FeedResource{
			Id:          f.Id,
			RssUrl:      f.RssUrl,
			Title:       f.Title,
			Link:        f.Link,
			Authors:     f.Authors,
			Description: f.Description,
			CreatedAt:   f.CreatedAt,
			UpdatedAt:   f.UpdatedAt,
		}
	})

	return &GetFeedsResponse{Data: results}, nil
}

// CreateFeeds creates new feeds with the given parameters.
//
//	@ID			CreateFeeds
//	@Summary	Create new feeds
//	@Tags		Feeds
//	@Accept		json
//	@Produce	json
//	@Param		request	body		CreateFeedsRequest	true	"CreateFeeds request"
//	@Success	200		{object}	CreateFeedsResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/feeds/create [post]
//	@Security	Authentication
func (c *controllerImpl) CreateFeeds(ctx context.Context, r *CreateFeedsRequest, usr user.User) (*CreateFeedsResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, fmt.Errorf("HTTP 400: %w", err)
	}

	feeds := make([]*store.Feed, 0)
	for _, i := range r.Data {
		feed := &store.Feed{
			Id:          c.idService.Feed(ctx),
			UserId:      usr.Id(),
			Type:        "podcast",
			Title:       i.Title,
			Link:        i.Link,
			Authors:     i.Authors,
			Description: i.Description,
			CreatedAt:   ytime.Now(),
			UpdatedAt:   ytime.Now(),
		}
		if err := c.rssController.BuildRss(ctx, feed); err != nil {
			return nil, fmt.Errorf("HTTP 500: build rss: %w", err)
		}
		feeds = append(feeds, feed)
	}

	if err := c.store.CreateFeeds(ctx, feeds); err != nil {
		return nil, fmt.Errorf("HTTP 500: create feeds: %w", err)
	}

	items := make([]*CreateFeedsResultResource, 0)
	for _, model := range feeds {
		items = append(items, &CreateFeedsResultResource{Id: model.Id})
	}

	return &CreateFeedsResponse{Data: items}, nil
}

// GetItems returns a list of items matching the given parameters.
//
//	@ID			GetItems
//	@Summary	Get items matching the given parameters
//	@Tags		Items
//	@Accept		json
//	@Produce	json
//	@Param		request	body		GetItemsRequest	true	"Parameters for filtering items"
//	@Success	200		{object}	GetItemsResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/items/get [post]
//	@Security	Authentication
func (c *controllerImpl) GetItems(ctx context.Context, req *GetItemsRequest, usr user.User) (*GetItemsResponse, error) {
	filter := store.ItemFilter{
		Ids:     req.Ids,
		FeedIds: req.FeedIds,
		UserIds: []string{usr.Id()},
	}
	if len(filter.UserIds) == 0 {
		filter.UserIds = []string{usr.Id()}
	}

	items, err := c.store.GetItems(ctx, &filter)
	if err != nil {
		return nil, fmt.Errorf("HTTP 500: get items: %w", err)
	}

	results := yslice.Map(items, func(i *store.Item) *ItemResource {
		return &ItemResource{
			Id: i.Id,
			File: &ItemFileResource{
				Id:          i.File.Id,
				Url:         i.File.UploadUrl,
				Size:        i.File.Size,
				ContentType: i.File.MimeType,
			},
			FeedId:      i.FeedId,
			Title:       i.Title,
			Link:        i.Link,
			Authors:     i.Authors,
			Description: i.Description,
			CreatedAt:   i.CreatedAt,
			UpdatedAt:   i.UpdatedAt,
			PublishedAt: i.PublishedAt,
		}
	})

	return &GetItemsResponse{Data: results}, nil
}

// CreateItems creates new items and returns a response with their IDs.
//
//	@ID			CreateItems
//	@Summary	Create new items and returns a response with their IDs
//	@Tags		Items
//	@Accept		json
//	@Produce	json
//	@Param		request	body		CreateItemsRequest	true	"CreateItems request"
//	@Success	200		{object}	CreateItemsResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/items/create [post]
//	@Security	Authentication
func (c *controllerImpl) CreateItems(ctx context.Context, r *CreateItemsRequest, usr user.User) (*CreateItemsResponse, error) {
	if err := r.Validate(); err != nil {
		return nil, fmt.Errorf("HTTP 400: %w", err)
	}

	items := make([]*store.Item, 0)

	err := c.store.Tx(ctx, func(ctxT context.Context) error {
		var fs []*store.File

		feeds, err := c.store.GetFeeds(ctxT, &store.FeedFilter{
			Ids: yslice.UniqueMap(r.Data, func(v *CreateItemsResource) string { return v.FeedId }),
		})
		if err != nil {
			return fmt.Errorf("HTTP 500: get feeds: %w", err)
		}
		feedsById := yslice.MapByKey(feeds, func(v *store.Feed) string { return v.Id })

		for _, i := range r.Data {
			if _, ok := feedsById[i.FeedId]; !ok {
				return fmt.Errorf("HTTP 404: no feed %s", i.FeedId)
			}

			file, err := c.store.GetFileById(ctxT, i.FileId)
			if err != nil {
				if err == store.ErrNotFound {
					return fmt.Errorf("HTTP 404: no file %s", i.FileId)
				}
				return fmt.Errorf("HTTP 500: get file: %w", err)
			}
			if file.ItemId != nil {
				return fmt.Errorf("HTTP 400: file already used")
			}

			publishedAt := i.PublishedAt
			if publishedAt.IsZero() {
				publishedAt = ytime.Now()
			}

			item := &store.Item{
				Id:          c.idService.Item(ctxT),
				FeedId:      i.FeedId,
				UserId:      usr.Id(),
				Title:       i.Title,
				Link:        i.Link,
				Authors:     i.Authors,
				Description: i.Description,
				FileId:      i.FileId,
				CreatedAt:   ytime.Now(),
				UpdatedAt:   ytime.Now(),
				PublishedAt: publishedAt,
			}
			items = append(items, item)

			file.ItemId = &item.Id
			fs = append(fs, file)
		}

		if err := c.store.CreateItems(ctxT, items); err != nil {
			return fmt.Errorf("HTTP 500: create items: %w", err)
		}
		if err := c.store.UpdateFiles(ctxT, fs, "item_id"); err != nil {
			return fmt.Errorf("HTTP 500: update files: %w", err)
		}

		if err := c.rssController.UpdateFeeds(ctxT, feeds); err != nil {
			return fmt.Errorf("HTTP 500: update feeds: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	results := make([]*CreateItemResultResource, 0)
	for _, item := range items {
		results = append(results, &CreateItemResultResource{Id: item.Id})
	}

	return &CreateItemsResponse{Data: results}, nil
}

// UploadFiles uploads new audio files and returns a response with the file IDs.
//
//	@ID			UploadFiles
//	@Summary	Upload new audio files
//	@Tags		Files
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		file	formData	file	true	"File to upload (can be repeated multiple times)"
//	@Success	200		{object}	UploadFilesResponse
//	@Failure	400		{object}	ErrorResponse
//	@Failure	401		{object}	ErrorResponse
//	@Failure	500		{object}	ErrorResponse
//	@Router		/files/upload [post]
//	@Security	Authentication
func (c *controllerImpl) UploadFiles(ctx context.Context, fs []multipart.File, usr user.User) (*UploadFilesResponse, error) {
	results := make([]*UploadFileResultResource, 0)

	for _, f := range fs {
		result := &UploadFileResultResource{}
		results = append(results, result)

		model, err := c.uploadFile(ctx, f, usr)
		if err != nil {
			result.Error = err.Error()
			continue
		}
		if err := c.store.CreateFile(ctx, model); err != nil {
			result.Error = err.Error()
			continue
		}

		result.Id = model.Id
	}

	return &UploadFilesResponse{Data: results}, nil
}

func (c *controllerImpl) uploadFile(ctx context.Context, f io.ReadSeeker, usr user.User) (*store.File, error) {
	info, err := files.Info(f)
	if err != nil {
		return nil, fmt.Errorf("read filetype: %w", err)
	}
	if info.Mime.Type != "audio" {
		return nil, fmt.Errorf("unsupported file type: %s", info.Mime.Value)
	}

	fileId := c.idService.File(ctx)

	path := fmt.Sprintf("files/%s/%s/%s.%s", info.Hash.Digest[:2], info.Hash.Digest[2:4], fileId, info.Extension)
	upload, err := c.fileStorage.Upload(ctx, path, f)
	if err != nil {
		return nil, fmt.Errorf("upload file: %w", err)
	}

	return &store.File{
		Id:        fileId,
		UserId:    usr.Id(),
		UploadUrl: upload.Url,
		UploadId:  upload.Id,
		Size:      info.Size,
		Hash:      info.Hash.String(),
		MimeType:  info.Mime.Value,
		CreatedAt: ytime.Now(),
		UpdatedAt: ytime.Now(),
	}, nil
}

func (c *controllerImpl) GetRssUrl(ctx context.Context, feedId string) (string, error) {
	filter := &store.FeedFilter{
		Ids: []string{feedId},
	}

	feeds, err := c.store.GetFeeds(ctx, filter)
	if err != nil {
		return "", fmt.Errorf("HTTP 500: get feed: %w", err)
	}
	if len(feeds) == 0 {
		return "", fmt.Errorf("HTTP 404: no feed %s", feedId)
	}

	return feeds[0].RssUrl, nil
}
