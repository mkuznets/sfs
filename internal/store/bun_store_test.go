package store_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"

	"mkuznets.com/go/sfs/internal/store"
)

func mockedStore(t *testing.T) (store.Store, sqlmock.Sqlmock) {
	sqldb, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	db := bun.NewDB(sqldb, sqlitedialect.New())
	return store.NewBunStore(db), sqlMock
}

// var testTime = ytime.New(time.Date(2022, 1, 1, 11, 12, 13, 14000, time.UTC))

//func Test_storeImpl_UpdateChannelFeeds(t *testing.T) {
//	st, sqlMock := mockedStore(t)
//
//	sqlMock.
//		ExpectExec(`WITH "_data" ("id", "feed_published_at", "feed_url") ` +
//			`AS (VALUES ` +
//			`('ch_123', 1641035533000, 'https://example.com/feed1.xml'), ` +
//			`('ch_456', 1641035533000, 'https://example.com/feed2.xml')) ` +
//			`UPDATE "channels" AS "ch" SET feed_published_at = _data.feed_published_at, ` +
//			`feed_url = _data.feed_url ` +
//			`FROM _data ` +
//			`WHERE (ch.id = _data.id)`).
//		WillReturnResult(sqlmock.NewResult(0, 0))
//
//	channels := []*store.Feed{
//		{
//			Id: "ch_123",
//			Feed: store.Feed{
//				Url:         "https://example.com/feed1.xml",
//				PublishedAt: testTime,
//			},
//		},
//		{
//			Id: "ch_456",
//			Feed: store.Feed{
//				Url:         "https://example.com/feed2.xml",
//				PublishedAt: testTime,
//			},
//		},
//	}
//
//	err := st.UpdateChannelFeeds(context.Background(), channels)
//	require.NoError(t, err)
//
//	assert.NoError(t, sqlMock.ExpectationsWereMet())
//
//	//event := bunHook.Events[0]
//	//assert.Equal(t, query, event.Query)
//}
