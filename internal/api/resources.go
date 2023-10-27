package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"mkuznets.com/go/sfs/internal/mtime"
	"mkuznets.com/go/sfs/internal/rss"
)

type GetFeedsRequest struct {
	Ids []string `json:"ids" extensions:"x-order=0"`
} // @name GetFeedsRequest

func (r *GetFeedsRequest) Validate() error {
	return nil
}

type GetFeedsResponse struct {
	Data []*FeedResource `json:"data"`
} // @name GetFeedsResponse

type FeedResource struct {
	Id          string     `json:"id" example:"feed_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	RssUrl      string     `json:"rss_url" example:"https://example.com/feed.rss" extensions:"x-order=1"`
	Title       string     `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=2"`
	Link        string     `json:"link" example:"https://example.com" extensions:"x-order=3"`
	Authors     string     `json:"authors" example:"The Owl" extensions:"x-order=4"`
	Description string     `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=5"`
	CreatedAt   mtime.Time `json:"created_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=6"`
	UpdatedAt   mtime.Time `json:"updated_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=7"`
} // @name FeedResource

type CreateFeedsRequest struct {
	Data []*CreateFeedsResource `json:"data"`
} // @name CreateFeedsRequest

func (r *CreateFeedsRequest) Validate() error {
	return validation.ValidateStruct(
		r, validation.Field(&r.Data, validation.Required, validation.Length(1, 100)),
	)
}

type CreateFeedsResponse struct {
	Data []*CreateFeedsResultResource `json:"data"`
} // @name CreateFeedsResponse

type CreateFeedsResource struct {
	Title       string `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=0"`
	Link        string `json:"link" example:"https://example.com" extensions:"x-order=1"`
	Authors     string `json:"authors" example:"The Owl" extensions:"x-order=2"`
	Description string `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=3"`
} // @name CreateFeedsResource

func (r *CreateFeedsResource) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.Title, validation.Required, validation.Length(1, 1024)),
		validation.Field(&r.Link, is.URL, validation.Length(0, 5*1024)),
		validation.Field(&r.Authors, validation.Length(0, 1024)),
		validation.Field(&r.Description, validation.Length(0, 64*1024)),
	)
}

type CreateFeedsResultResource struct {
	Id string `json:"id" example:"feed_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
} // @name CreateFeedsResultResource

type GetItemsRequest struct {
	Ids     []string `json:"ids" extensions:"x-order=0"`
	FeedIds []string `json:"feed_ids" extensions:"x-order=1"`
} // @name GetItemsRequest

type GetItemsResponse struct {
	Data []*ItemResource `json:"data"`
} // @name GetItemsResponse

type ItemResource struct {
	Id          string            `json:"id" example:"item_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	File        *ItemFileResource `json:"file,omitempty" extensions:"x-order=1"`
	FeedId      string            `json:"feed_id" example:"feed_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=2"`
	Title       string            `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=3"`
	Link        string            `json:"link" example:"https://example.com" extensions:"x-order=4"`
	Authors     string            `json:"authors" example:"The Owl" extensions:"x-order=5"`
	Description string            `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=6"`
	CreatedAt   mtime.Time        `json:"created_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=7"`
	UpdatedAt   mtime.Time        `json:"updated_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=8"`
	PublishedAt mtime.Time        `json:"published_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=9"`
} // @name ItemResource

type ItemFileResource struct {
	Id          string `json:"id,omitempty" example:"file_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	Url         string `json:"url" example:"https://example.com/file.mp3" extensions:"x-order=1"`
	Size        int64  `json:"size" example:"123456" extensions:"x-order=2"`
	ContentType string `json:"content_type" example:"audio/mpeg" extensions:"x-order=3"`
} // @name ItemFileResource

type CreateItemsRequest struct {
	Data []*CreateItemsResource `json:"data"`
} // @name CreateItemsRequest

func (r *CreateItemsRequest) Validate() error {
	return validation.ValidateStruct(
		r, validation.Field(&r.Data, validation.Required, validation.Length(1, 10000)),
	)
}

type CreateItemsResource struct {
	FileId      string     `json:"file_id" example:"file_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	FeedId      string     `json:"feed_id" example:"feed_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=1"`
	Title       string     `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=2"`
	Link        string     `json:"link" example:"https://example.com" extensions:"x-order=3"`
	Authors     string     `json:"authors" example:"The Owl" extensions:"x-order=4"`
	Description string     `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=5"`
	PublishedAt mtime.Time `json:"published_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=6"`
} // @name CreateItemsResource

func (r *CreateItemsResource) Validate() error {
	return validation.ValidateStruct(
		r,
		validation.Field(&r.FileId, validation.Required, validation.Length(1, 1024)),
		validation.Field(&r.FeedId, validation.Required, validation.Length(1, 1024)),
		validation.Field(&r.Title, validation.Required, validation.Length(1, 1024)),
		validation.Field(&r.Link, is.URL, validation.Length(0, 5*1024)),
		validation.Field(&r.Authors, validation.Length(0, 1024)),
		validation.Field(&r.Description, validation.Length(0, 64*1024)),
	)
}

type CreateItemsResponse struct {
	Data []*CreateItemResultResource `json:"data"`
} // @name CreateItemsResponse

type CreateItemResultResource struct {
	Id string `json:"id" example:"item_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
} // @name CreateItemResultResource

type UploadFilesResponse struct {
	Data []*UploadFileResultResource `json:"data"`
} // @name UploadFilesResponse

type UploadFileResultResource struct {
	Id    string `json:"id,omitempty" example:"file_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	Error string `json:"error,omitempty" example:"invalid file format" extensions:"x-order=1"`
} // @name UploadFileResultResource

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
} // @name ErrorResponse

// Only used in Swagger docs.
type _ rss.Podcast // @name Podcast
