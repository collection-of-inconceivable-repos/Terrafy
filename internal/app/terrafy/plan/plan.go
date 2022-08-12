package plan

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/collection-of-inconceivable-repos/terrafy/internal/app/terrafy/spotify"
	"gopkg.in/yaml.v3"
)

func ExecPlanCmd(planFlagSet *flag.FlagSet) {
	planFileName := parsePlanFlags(planFlagSet)
	fmt.Println(planFileName)
	loadFile(planFileName)

	api := spotify.New()

	req := spotify.SearchTrackRequest{
		Title:   "Under Pressure",
		Artists: []string{"Queen"},
	}

	result, err := api.SearchForTrack(req)
	if err != nil {
		log.Fatalf("Error while searching for track: %s\n", err)
	}

	fmt.Printf("%s [%s]\n", result, result.ID)
}

func parsePlanFlags(planFlagSet *flag.FlagSet) string {
	planFlagSet.Parse(os.Args[2:])
	planArgs := planFlagSet.Args()

	if len(planArgs) != 1 {
		printPlanUsageAndExit()
	}

	return planArgs[0]
}

func loadFile(planFileName string) {
	planFileContents, err := ioutil.ReadFile(planFileName)
	if err != nil {
		log.Fatalf("Failed to read file '%s'\n", planFileName)
	}

	var playlistDefinition PlaylistDefinition
	err = yaml.Unmarshal(planFileContents, &playlistDefinition)
	if err != nil {
		log.Fatalf("Failed to parse YAML file '%s'\n", planFileName)
	}

	// fmt.Printf("%#v\n", playlistDefinition)

	// for i, track := range playlistDefinition.Tracks {
	// 	fmt.Printf("%d - %s - %t\n", i, track.Link, track.HasSpotifyLink())
	// }
}

func printPlanUsageAndExit() {
	fmt.Printf("Usage: %s plan <file>\n", os.Args[0])
	os.Exit(1)
}
