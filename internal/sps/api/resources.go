package api

import "mkuznets.com/go/sps/internal/types"

type CreateChannelRequest struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Authors     string `json:"authors"`
	Description string `json:"description"`
}

type ChannelResponse struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Link        string     `json:"link"`
	Authors     string     `json:"authors"`
	Description string     `json:"description"`
	CreatedAt   types.Time `json:"created_at"`
	UpdatedAt   types.Time `json:"updated_at"`
}
