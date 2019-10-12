package main

import (
	"errors"
	"fmt"

	"github.com/Sinar/go-pardocs/internal/debate"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("go-pardebate ..")

	//p, err := debate.NewPDFDoc("./internal/debate/testdata/DR-11042019.pdf")
	//if err != nil {
	//	panic(err)
	//}

	t, err := debate.NewDebateTOC("./internal/debate/testdata/DR-11042019.pdf")
	if err != nil {
		errNoTOC, ok := errors.Unwrap(err).(debate.ErrorNoTOCFound)
		if ok {
			fmt.Println("TO_HANDLE: ", errNoTOC.Error())
			//spew.Dump(errNoTOC)
		} else {
			panic(err)
		}
	}

	spew.Dump(t)

}
