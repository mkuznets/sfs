package ui

import (
	"bytes"
	"context"
	"mkuznets.com/go/sps/internal/api"
	"mkuznets.com/go/sps/internal/templates"
	"text/template"
)

type Controller interface {
	Page(ctx context.Context, page Page) (string, error)

	GetLoginContent(ctx context.Context) (string, error)
	GetLoginPage(ctx context.Context) (string, error)

	GetPodcastsPage(ctx context.Context, userId string) (string, error)
	GetPodcastsContent(ctx context.Context, userId string) (string, error)
	GetPodcastsHeader(ctx context.Context) (string, error)
	GetPodcastsBody(ctx context.Context, userId string) (string, error)

	GetEpisodesPage(ctx context.Context, userId, channelId string) (string, error)
	GetEpisodesContent(ctx context.Context, userId, channelId string) (string, error)
	GetEpisodesHeader(ctx context.Context, channelId string) (string, error)
	GetEpisodesBody(ctx context.Context, userId, channelId string) (string, error)
}

type controllerImpl struct {
	ac  api.Controller
	tpl *template.Template
}

func NewController(ac api.Controller) Controller {
	tpl := template.Must(template.ParseFS(templates.Templates, "*.html"))
	return &controllerImpl{
		ac:  ac,
		tpl: tpl,
	}
}

func (c *controllerImpl) Page(ctx context.Context, page Page) (string, error) {
	var buf bytes.Buffer
	if err := c.tpl.ExecuteTemplate(&buf, "index.html", page); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *controllerImpl) GetLoginContent(ctx context.Context) (string, error) {
	var buf bytes.Buffer
	err := c.tpl.ExecuteTemplate(&buf, "login.html", nil)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *controllerImpl) GetLoginPage(ctx context.Context) (string, error) {
	var buf bytes.Buffer

	content, err := c.GetLoginContent(ctx)
	if err != nil {
		return "", err
	}

	err = c.tpl.ExecuteTemplate(&buf, "index.html", map[string]interface{}{
		"Content": content,
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *controllerImpl) GetPodcastsPage(ctx context.Context, userId string) (string, error) {
	var out bytes.Buffer

	content, err := c.GetPodcastsContent(ctx, userId)
	if err != nil {
		return "", err
	}

	if err := c.tpl.ExecuteTemplate(&out, "index.html", &Page{Content: content}); err != nil {
		return "", err
	}

	return out.String(), nil
}

func (c *controllerImpl) GetPodcastsContent(ctx context.Context, userId string) (string, error) {
	var out bytes.Buffer

	bc, err := c.GetPodcastsHeader(ctx)
	if err != nil {
		return "", err
	}

	body, err := c.GetPodcastsBody(ctx, userId)
	if err != nil {
		return "", err
	}

	err = c.tpl.ExecuteTemplate(&out, "content.html", &Content{
		Header: bc,
		Body:   body,
	})
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func (c *controllerImpl) GetPodcastsHeader(ctx context.Context) (string, error) {
	var out bytes.Buffer

	data := []Header{}
	if err := c.tpl.ExecuteTemplate(&out, "content_header.html", data); err != nil {
		return "", err
	}

	return out.String(), nil
}

func (c *controllerImpl) GetPodcastsBody(ctx context.Context, userId string) (string, error) {
	var out bytes.Buffer

	channels, err := c.ac.ListChannels(ctx, userId)
	if err != nil {
		return "", err
	}
	err = c.tpl.ExecuteTemplate(&out, "content_body_podcasts.html", channels)
	if err != nil {
		return "", err
	}

	return out.String(), err
}

func (c *controllerImpl) GetEpisodesHeader(ctx context.Context, channelId string) (string, error) {
	channel, err := c.ac.GetChannel(ctx, channelId)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer

	data := []Header{
		{Icon: "fa-solid fa-rss", Text: channel.Title},
	}
	if err := c.tpl.ExecuteTemplate(&out, "content_header.html", data); err != nil {
		return "", err
	}

	return out.String(), nil

}

func (c *controllerImpl) GetEpisodesBody(ctx context.Context, userId, channelId string) (string, error) {
	var out bytes.Buffer

	episodes, err := c.ac.ListEpisodes(ctx, userId, channelId)
	if err != nil {
		return "", err
	}
	err = c.tpl.ExecuteTemplate(&out, "content_body_episodes.html", episodes)
	if err != nil {
		return "", err
	}

	return out.String(), err
}

func (c *controllerImpl) GetEpisodesContent(ctx context.Context, userId, channelId string) (string, error) {
	var out bytes.Buffer

	bc, err := c.GetEpisodesHeader(ctx, channelId)
	if err != nil {
		return "", err
	}

	body, err := c.GetEpisodesBody(ctx, userId, channelId)
	if err != nil {
		return "", err
	}

	err = c.tpl.ExecuteTemplate(&out, "content.html", &Content{
		Header: bc,
		Body:   body,
	})
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func (c *controllerImpl) GetEpisodesPage(ctx context.Context, userId, channelId string) (string, error) {
	var out bytes.Buffer

	content, err := c.GetEpisodesContent(ctx, userId, channelId)
	if err != nil {
		return "", err
	}

	if err := c.tpl.ExecuteTemplate(&out, "index.html", &Page{Content: content}); err != nil {
		return "", err
	}

	return out.String(), nil
}
