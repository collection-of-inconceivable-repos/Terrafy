package search

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/collection-of-inconceivable-repos/terrafy/internal/app/terrafy/spotify"
)

func ExecSearchCmd(searchFlagSet *flag.FlagSet) {
	req := parseSearchFlags(searchFlagSet)
	api := spotify.New()
	result, err := api.SearchForTrack(req)

	if err != nil {
		log.Fatalf("Error searching for track: %s", err)
	}

	log.Printf("  Title: %s\n  Album: %s\nArtists: %s\n     ID: %s\n    URI: %s\n   Link: %s\n",
		result.Title,
		result.Album,
		strings.Join(result.Artists, ", "),
		string(result.ID),
		string(result.URI),
		string(result.Link),
	)
}

func parseSearchFlags(searchFlagSet *flag.FlagSet) spotify.SearchTrackRequest {
	searchFlagSet.Parse(os.Args[2:])
	planArgs := searchFlagSet.Args()

	if len(planArgs) < 1 {
		printPlanUsageAndExit()
	}

	title := planArgs[0]
	artists := planArgs[1:]

	return spotify.SearchTrackRequest{
		Title:   title,
		Artists: artists,
	}
}

func printPlanUsageAndExit() {
	fmt.Printf("Usage: %s search <title> [artists]\n", os.Args[0])
	os.Exit(1)
}
