package main

import (
	"fmt"

	"github.com/Sinar/go-pardocs/internal/debate"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("go-pardebate ..")

	p, err := debate.NewPDFDoc("./internal/debate/testdata/DR-11042019.pdf")
	if err != nil {
		panic(err)
	}

	spew.Dump(p)

}
