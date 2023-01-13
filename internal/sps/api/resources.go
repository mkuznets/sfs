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

type CreateEpisodeRequest struct {
	FileId      string `json:"file_id" example:"file_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	Title       string `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=1"`
	Link        string `json:"link" example:"https://example.com" extensions:"x-order=2"`
	Authors     string `json:"authors" example:"The Owl" extensions:"x-order=3"`
	Description string `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=4"`
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

type EpisodeResponse struct {
	Id          string        `json:"id" example:"ep_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	File        *FileResponse `json:"file,omitempty" extensions:"x-order=1"`
	ChannelId   string        `json:"channel_id" example:"ch_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=2"`
	Title       string        `json:"title" example:"Bored Owls Online Radio" extensions:"x-order=3"`
	Link        string        `json:"link" example:"https://example.com" extensions:"x-order=4"`
	Authors     string        `json:"authors" example:"The Owl" extensions:"x-order=5"`
	Description string        `json:"description" example:"Bored owls talk about whatever happens to be on their minds" extensions:"x-order=6"`
	CreatedAt   types.Time    `json:"created_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=7"`
	UpdatedAt   types.Time    `json:"updated_at" swaggertype:"string" format:"date-time" example:"2023-01-01T01:02:03.456Z" extensions:"x-order=8"`
} // @name EpisodeResponse

type FileResponse struct {
	Id          string `json:"id,omitempty" example:"file_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
	Url         string `json:"url" example:"https://example.com/file.mp3" extensions:"x-order=1"`
	Size        int64  `json:"size" example:"123456" extensions:"x-order=2"`
	ContentType string `json:"content_type" example:"audio/mpeg" extensions:"x-order=3"`
}

// Only used in Swagger docs.
type _ herror.Response // @name ErrorResponse

type UploadResponse struct {
	Id string `json:"id" example:"file_2K9BWVNuo3sG4yM322fbP3mB6ls" extensions:"x-order=0"`
} // @name UploadResponse