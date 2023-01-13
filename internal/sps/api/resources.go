package api

import (
	"mkuznets.com/go/sps/internal/herror"
	"mkuznets.com/go/sps/internal/types"
)

type CreateChannelRequest struct {
	Title       string `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=0"`
	Link        string `json:"link" example:"https://example.com" extensions:"x-order=1"`
	Authors     string `json:"authors" example:"The Owl"`
	Description string `json:"description" example:"Bored owls talk about whatever happens to be on their minds"`
} // @name CreateChannelRequest

type ChannelResponse struct {
	Id          string     `json:"id" example:"ch_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	Title       string     `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=1"`
	Link        string     `json:"link" example:"https://example.com" extensions:"x-order=3"`
	Authors     string     `json:"authors" example:"The Owl" extensions:"x-order=4"`
	Description string     `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=5"`
	CreatedAt   types.Time `json:"created_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=6"`
	UpdatedAt   types.Time `json:"updated_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=7"`
} // @name ChannelResponse

// Only used in Swagger docs.
type _ herror.Response // @name ErrorResponse

type UploadResponse struct {
	Id string `json:"id" example:"file_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
} // @name UploadResponse
