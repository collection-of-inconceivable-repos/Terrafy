package terrafy

import (
	"fmt"
	"os"
)

func PrintUsageAndExit() {
	fmt.Printf("Usage: %s plan|apply [<args>]\n", os.Args[0])
	os.Exit(1)
}
