package api_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"mkuznets.com/go/sps/internal/sps/api"
	"mkuznets.com/go/sps/internal/types"
	"testing"
	"time"
)

func mockedStore(t *testing.T) (api.Store, sqlmock.Sqlmock) {
	sqldb, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	db := bun.NewDB(sqldb, sqlitedialect.New())
	return api.NewStore(db), sqlMock
}

var testTime = types.NewTime(time.Date(2022, 1, 1, 11, 12, 13, 14000, time.UTC))

func Test_storeImpl_UpdateChannelFeeds(t *testing.T) {
	store, sqlMock := mockedStore(t)

	sqlMock.
		ExpectExec(`WITH "_data" ("id", "feed_content", "feed_published_at", "feed_url") ` +
			`AS (VALUES ` +
			`('ch_123', X'3c786d6c3e3c2f786d6c3e', 1641035533000, 'https://example.com/feed1.xml'), ` +
			`('ch_456', X'3c786d6c3e3c2f786d6c3e', 1641035533000, 'https://example.com/feed2.xml')) ` +
			`UPDATE "channels" AS "ch" SET feed_content = _data.feed_content, ` +
			`feed_published_at = _data.feed_published_at, ` +
			`feed_url = _data.feed_url ` +
			`FROM _data ` +
			`WHERE (ch.id = _data.id)`).
		WillReturnResult(sqlmock.NewResult(0, 0))

	channels := []*api.Channel{
		{
			Id: "ch_123",
			Feed: api.Feed{
				Content:     []byte("<xml></xml>"),
				Url:         "https://example.com/feed1.xml",
				PublishedAt: testTime,
			},
		},
		{
			Id: "ch_456",
			Feed: api.Feed{
				Content:     []byte("<xml></xml>"),
				Url:         "https://example.com/feed2.xml",
				PublishedAt: testTime,
			},
		},
	}

	err := store.UpdateChannelFeeds(context.Background(), channels)
	require.NoError(t, err)

	assert.NoError(t, sqlMock.ExpectationsWereMet())

	//event := hook.Events[0]
	//assert.Equal(t, query, event.Query)
}
