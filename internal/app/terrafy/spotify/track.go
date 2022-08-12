package spotify

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/collection-of-inconceivable-repos/terrafy/internal/app/terrafy/util"
	spotifylib "github.com/zmb3/spotify/v2"
)

type TrackId string
type Track struct {
	Title   string
	Album   string
	Artists []string
	Link    string
	URI     Uri
	ID      TrackId
}

type SearchTrackRequest struct {
	Title   string
	Album   string
	Artists []string
}

func (track Track) String() string {
	return fmt.Sprintf("%s | %s by %s", track.Title, track.Album, strings.Join(track.Artists, ", "))
}

func (api SpotifyApi) SearchForTrack(request SearchTrackRequest) (*Track, error) {
	ctx := context.Background()
	query := buildSearchQuery(request)

	results, err := api.client.Search(ctx, query,
		spotifylib.SearchTypeTrack,
		spotifylib.Market("US"),
		spotifylib.Limit(1))

	if err != nil {
		log.Printf("Search for track with query '%s' failed\n%w", query, err)
		return nil, errors.New("spotify search API error")
	}

	tracks := results.Tracks.Tracks
	if len(tracks) == 0 {
		return nil, errors.New("no matching tracks found")
	}

	track := tracks[0]
	normalizedTrack := &Track{
		ID:      TrackId(track.ID),
		URI:     Uri(track.URI),
		Link:    "https://open.spotify.com/track/" + string(track.ID),
		Title:   track.Name,
		Album:   track.Album.Name,
		Artists: extractArtistNames(track.Artists),
	}

	return normalizedTrack, nil
}

func extractArtistNames(artists []spotifylib.SimpleArtist) []string {
	mapFn := func(artist spotifylib.SimpleArtist) string {
		return artist.Name
	}
	return util.MapSlice(artists, mapFn)
}

func buildSearchQuery(request SearchTrackRequest) string {
	searchParams := make([]string, 0)

	if !util.IsBlank(request.Title) {
		searchParams = append(searchParams, `track:"`+request.Title+`"`)
	}

	for _, artist := range request.Artists {
		if !util.IsBlank(artist) {
			searchParams = append(searchParams, `artist:"`+artist+`"`)
		}
	}

	if !util.IsBlank(request.Album) {
		searchParams = append(searchParams, `album:"`+request.Album+`"`)
	}

	return strings.Join(searchParams, " ")
}
