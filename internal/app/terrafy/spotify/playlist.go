package spotify

import (
	"context"
	"log"

	spotifyapi "github.com/zmb3/spotify/v2"
)

func (api SpotifyApi) ViewPlaylist(playlistID string) {
	ctx := context.Background()
	page, err := api.client.GetPlaylistItems(ctx, spotifyapi.ID(playlistID))
	if err != nil {
		log.Fatalf("failed to load playlist items page: %s\n", err)
	}

	for err == nil {
		err = api.client.NextPage(ctx, page)
	}

	if err != spotifyapi.ErrNoMorePages {
		log.Fatalf("failed to load playlist items page: %s\n", err)
	}
}
