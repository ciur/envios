package main

import (
	"flag"
	"fmt"

	"github.com/ciur/enward/parser"
)

var (
	flagConfig = flag.String("c", ".enwardrc", "config file")
)

func main() {
	flag.Parse()

	profiles, error := parser.LoadConfig(*flagConfig)
	fmt.Printf("%v %s\n", profiles, error)

}
