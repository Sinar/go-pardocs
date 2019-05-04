package main

import (
	"flag"
	"fmt"

	"github.com/Sinar/go-pardocs/cmd"
)

func main() {
	//defer recoverFromPanic()
	fmt.Println("Welcome to go-pardocs!")

	flag.Parse()

	// Experiments only ..
	//cmd.SplitBukanLisanPDFsSpike()

	// Actual
	cmd.SplitBukanLisanPDFs()

}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ==>", r)
	}

}
