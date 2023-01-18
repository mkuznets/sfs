package api_test

import (
	"mkuznets.com/go/sps/api"
	"mkuznets.com/go/sps/api/client/channels"
	"mkuznets.com/go/sps/api/models"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU1MTg1MjUsImlhdCI6MTY3Mzk4MjUyNSwidWlkIjoidXNyXzJLU3l3OHRMU0c1ZjQ1bHVodllSOGMyZFh5VSJ9.rTeiDNQtofA_d9d31Y3Gu0UkspO4EY_qkNLiE1ZIHwWqNMuoYL_y9mdIneq0b8nRzrimOHEnNJ5koY6u8pmeLnKsNcoFBM5rQnZJKB_eRdf7vrtTFjpXRBo-z9gkiyFisgGYu59950J3qzNyXolmKPmIn7jQumIz09Z41Du31U37ihzDcw9d1DHD1UT-9uEgeGhmYsBnjHSZE5UYMuGtDFBJhldyqpoo2h7rZOU2qtjc0nCGIkDz6ErLLk84VytqMU4FywHidvZ3rh_rW1OxJyjix7d_jCty8RPXD723rHtIDp28wPl4I7YYb56dVhcn7vuecAgNB00lwtgyWoDtYA"
	sps, err := api.New("http://localhost:8080/api", token)
	if err != nil {
		t.Fatal(err)
	}

	params := channels.NewCreateChannelParams().
		WithRequest(&models.CreateChannelRequest{
			Authors:     "A, B, C",
			Description: "foo",
			Link:        "https://example.com",
			Title:       "bara",
		})

	resp, err := sps.Channels.CreateChannel(params, sps.AuthInfo())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.GetPayload())

	resp2, err := sps.Channels.GetChannel(channels.NewGetChannelParams().WithID(resp.GetPayload().ID), sps.AuthInfo())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp2.GetPayload())
	t.Log(time.Time(resp2.GetPayload().CreatedAt))

}
