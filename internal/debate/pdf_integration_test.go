package debate_test

import (
	"fmt"
	"testing"

	"github.com/Sinar/go-pardocs/internal/debate"
	"github.com/davecgh/go-spew/spew"
)

func TestDebatePDFLoad(t *testing.T) {

	fmt.Println("Call TestDebatePDFLoad!!")

	p, err := debate.NewPDFDoc("testdata/abc.pdf")
	if err != nil {
		panic(err)
	}

	spew.Dump(p)

}
