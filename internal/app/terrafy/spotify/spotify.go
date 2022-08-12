package spotify

import (
	"context"
	"log"
	"os"

	"github.com/collection-of-inconceivable-repos/terrafy/internal/app/terrafy/util"
	spotifylib "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

type SpotifyApi struct {
	client *spotifylib.Client
}

type Uri string

func New() SpotifyApi {
	accessToken := os.Getenv("SPOTIFY_ACCESS_TOKEN")
	if util.IsBlank(accessToken) {
		log.Fatalf("Please set the SPOTIFY_ACCESS_TOKEN environment variable")
	}

	ctx := context.Background()
	oauthToken := oauth2.Token{AccessToken: accessToken, TokenType: "Bearer"}
	httpClient := spotifyauth.New().Client(ctx, &oauthToken)

	var api SpotifyApi
	api.client = spotifylib.New(httpClient)

	return api
}
