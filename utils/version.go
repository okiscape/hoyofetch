package utils

import "fmt"

var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

func PrintVersion() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Build: %s\n", BuildTime)
	fmt.Printf("Commit: %s\n", GitCommit)
}
