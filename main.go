package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ciur/enward/parser"
	"github.com/ciur/enward/profiles"
)

var (
	flagConfig         = flag.String("c", ".enwardrc", "config file")
	flagUseProfileName = flag.String("n", "", "Profile to use")
)

func main() {
	flag.Parse()

	allProfiles, error := parser.LoadConfig(*flagConfig)

	if error != "" {
		fmt.Printf("Parsing error: %v\n", error)
		os.Exit(1)
	}
	profilesList := profiles.ProfilesList(allProfiles)
	found := profilesList.FindProfile(*flagUseProfileName)
	if found != nil {
		found.PrintExports()
	} else {
		fmt.Printf("Profile %s not found\n", *flagUseProfileName)
		os.Exit(1)
	}
}
