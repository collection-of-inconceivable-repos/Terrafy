package main

import (
	"flag"
	"log"
	"os"

	"github.com/collection-of-inconceivable-repos/terrafy/internal/app/terrafy"
	"github.com/collection-of-inconceivable-repos/terrafy/internal/app/terrafy/plan"
	"github.com/collection-of-inconceivable-repos/terrafy/internal/app/terrafy/search"
)

const callback string = "http://localhost:8888/callback"

func main() {
	setupLogger()
	parseArguments()
}

func parseArguments() {
	planFlagSet := flag.NewFlagSet("plan", flag.ExitOnError)
	searchFlagSet := flag.NewFlagSet("search", flag.ExitOnError)

	if len(os.Args) < 2 {
		terrafy.PrintUsageAndExit()
	}

	switch os.Args[1] {
	case "plan":
		plan.ExecPlanCmd(planFlagSet)
	case "search":
		search.ExecSearchCmd(searchFlagSet)
	case "apply":
		panic("Implement apply cmd")
	default:
		terrafy.PrintUsageAndExit()
	}
}

func setupLogger() {
	log.SetFlags(0)
}
