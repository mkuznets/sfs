package sps

type Channel struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Authors     string `json:"authors"`
	Description string `json:"description"`
}

type Episode struct {
	Id        string `json:"id"`
	ChannelId string `json:"channel_id"`
	Title     string `json:"title"`
}
