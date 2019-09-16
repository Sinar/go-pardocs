package debate_test

import (
	"fmt"
	"testing"

	"github.com/Sinar/go-pardocs/internal/debate"
	"github.com/davecgh/go-spew/spew"
)

func TestDebatePDFLoad(t *testing.T) {

	fmt.Println("Call TestDebatePDFLoad!!")

	tests := []struct {
		name        string
		testPDFPath string
	}{
		{"test #1", "testdata/DR-01072019_new.pdf"}, // Bad streaming data on page 2 (is empty page; no string)
		{"test #2", "testdata/DR-01042019.pdf"},
		{"test #3", "testdata/DR-02042019.pdf"},
		{"test #4", "testdata/DR-11042019.pdf"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Case: DR-01072019.pdf July 01, 2019 Session
			p, err := debate.NewPDFDoc(tt.testPDFPath)
			if err != nil {
				panic(err)
			}
			spew.Dump(p.Pages)
		})
	}

}
