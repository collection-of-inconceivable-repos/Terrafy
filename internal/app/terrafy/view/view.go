package view

import (
	"flag"
	"fmt"
	"os"
)

func ViewPlanCmd(viewFlagSet *flag.FlagSet) {
	playlistId := parseViewFlags(viewFlagSet)

	// api = spotify.New()

	fmt.Printf("%s", playlistId)
}

func parseViewFlags(viewFlagSet *flag.FlagSet) string {
	viewFlagSet.Parse(os.Args[:2])
	viewArgs := viewFlagSet.Args()

	if len(viewArgs) != 1 {
		printViewUsageAndExit()
	}

	return viewArgs[0]
}

func printViewUsageAndExit() {
	fmt.Printf("Usage: %s view <playlist-id>\n", os.Args[0])
	os.Exit(1)
}
