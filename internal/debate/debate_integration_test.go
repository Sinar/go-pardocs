package debate_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Sinar/go-pardocs/internal/debate"
)

func TestNewDebateTOC(t *testing.T) {
	type args struct {
		sourcePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *debate.DebateTOC
		wantErr bool
	}{
		//{"Missing TOC", args{"testdata/Bad-DR-DewanSelangor.pdf"}, nil, true},
		{"badly formed", args{"testdata/DR-01072019.pdf"}, &debate.DebateTOC{}, false},
		{"normal #1", args{"testdata/DR-11042019.pdf"}, &debate.DebateTOC{}, false},
		{"normal #2", args{"testdata/DR-01072019_new.pdf"}, &debate.DebateTOC{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := debate.NewDebateTOC(tt.args.sourcePath)
			// If unexpected errors!?
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDebateTOC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Test all error scenarios ..
			if err != nil {
				// Check is correct error type! when have error!
				errNoTOC, ok := errors.Unwrap(err).(debate.ErrorNoTOCFound)
				if ok {
					fmt.Println("PASS: ", errNoTOC.Error())
					//spew.Dump(errNoTOC)
				} else {
					t.Fail()
				}
			}

			// Test TOC structure out ..
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDebateTOC() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeTOC(t *testing.T) {
	type args struct {
		pdfDoc *debate.PDFDocument
	}
	tests := []struct {
		name          string
		args          args
		wantStartPage int
		wantEndPage   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStartPage, gotEndPage := debate.RangeTOC(tt.args.pdfDoc)
			if gotStartPage != tt.wantStartPage {
				t.Errorf("RangeTOC() gotStartPage = %v, want %v", gotStartPage, tt.wantStartPage)
			}
			if gotEndPage != tt.wantEndPage {
				t.Errorf("RangeTOC() gotEndPage = %v, want %v", gotEndPage, tt.wantEndPage)
			}
		})
	}
}
