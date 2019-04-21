package main

import (
	"flag"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Welcome to go-pardocs!")

	flag.Parse()

	spew.Println("bob")
}
