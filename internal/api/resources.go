package api

import (
	"mkuznets.com/go/sps/internal/ytils/yerr"
	"mkuznets.com/go/sps/internal/ytils/ytime"
)

type GetFeedsRequest struct {
	Ids     []string `json:"ids" extensions:"x-order=0"`
	UserIds []string
} // @name GetFeedsRequest

type GetFeedsResponse struct {
	Data []*FeedResource `json:"data"`
} // @name GetFeedsResponse

type FeedResource struct {
	Id          string     `json:"id" example:"feed_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	FeedUrl     string     `json:"feed_url" extensions:"x-order=1"`
	Title       string     `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=2"`
	Link        string     `json:"link" example:"https://example.com" extensions:"x-order=3"`
	Authors     string     `json:"authors" example:"The Owl" extensions:"x-order=4"`
	Description string     `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=5"`
	CreatedAt   ytime.Time `json:"created_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=6"`
	UpdatedAt   ytime.Time `json:"updated_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=7"`
} // @name FeedResource

type CreateFeedsRequest struct {
	Data []*CreateFeedsResource `json:"data"`
} // @name CreateFeedsRequest

type CreateFeedsResponse struct {
	Data []*CreateFeedsResultResource `json:"data"`
} // @name CreateFeedsResponse

type CreateFeedsResource struct {
	Title       string `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=0"`
	Link        string `json:"link" example:"https://example.com" extensions:"x-order=1"`
	Authors     string `json:"authors" example:"The Owl" extensions:"x-order=2"`
	Description string `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=3"`
} // @name CreateFeedsResource

type CreateFeedsResultResource struct {
	Id string `json:"id" example:"feed_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
} // @name CreateFeedsResultResource

type GetItemsRequest struct {
	Ids     []string `json:"ids" extensions:"x-order=0"`
	FeedIds []string `json:"feed_ids" extensions:"x-order=1"`
	UserIds []string
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
	CreatedAt   ytime.Time        `json:"created_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=7"`
	UpdatedAt   ytime.Time        `json:"updated_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=8"`
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

type CreateItemsResponse struct {
	Data []*CreateItemResultResource `json:"data"`
} // @name CreateItemsResponse

type CreateItemsResource struct {
	FileId      string `json:"file_id" example:"file_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	FeedId      string `json:"feed_id" example:"feed_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=1"`
	Title       string `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=2"`
	Link        string `json:"link" example:"https://example.com" extensions:"x-order=3"`
	Authors     string `json:"authors" example:"The Owl" extensions:"x-order=4"`
	Description string `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=5"`
} // @name CreateItemsResource

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

// Only used in Swagger docs.
type _ yerr.Response // @name ErrorResponse
