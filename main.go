package main

import (
	"flag"
	"fmt"

	"github.com/Sinar/go-pardocs/cmd"
)

func main() {
	fmt.Println("Welcome to go-pardocs!")

	flag.Parse()

	// Actual planning
	//cmd.SplitBukanLisanPDFs()
	// Test split
	cmd.ExecuteSplitPlan()

}
