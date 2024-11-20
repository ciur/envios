package main

import (
	"bufio"
	"envios/parser"
	"flag"
	"fmt"
	"os"
)

type Profile struct {
	name string
}

var (
	flagConfig = flag.String("c", ".enviousrc", "config file")
)

func main() {
	flag.Parse()

	file, err := os.Open(*flagConfig)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", *flagConfig, err)
		os.Exit(1)
	}

	defer file.Close()

	loadProfiles(file)
}

func loadProfiles(file *os.File) []Profile {
	loadTokens(file)
	return make([]Profile, 0)
}

func loadTokens(file *os.File) []parser.Token {
	var tokens []parser.Token

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parser.ParseProfileLine(line)
		fmt.Print(line)
	}

	return tokens
}
