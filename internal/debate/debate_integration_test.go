package debate_test

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/sanity-io/litter"

	"github.com/Sinar/go-pardocs/internal/debate"
)

var update = flag.Bool("update", false, "update .golden files")

func TestNewDebateTOC(t *testing.T) {
	type args struct {
		sourcePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		//{"Missing TOC", args{"testdata/Bad-DR-DewanSelangor.pdf"}, nil, true},
		{"empty page2", args{"testdata/DR-01072019.pdf"}, false},
		{"normal #1", args{"testdata/DR-11042019.pdf"}, false},
		{"normal #2", args{"testdata/DR-01072019_new.pdf"}, false},
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
			} else {
				// Use Goldenfile pattern
				sq := litter.Options{
					HidePrivateFields: false,
				}
				actual := []byte(sq.Sdump(got))
				golden := filepath.Join("testdata", tt.name+".golden")
				if *update {
					ioutil.WriteFile(golden, actual, 0644)
				}
				want, rerr := ioutil.ReadFile(golden)
				if rerr != nil {
					// Cannot proceed with one golden file update
					if os.IsNotExist(rerr) {
						t.Fatalf("Ensure run with -update flag first time! ERR: %s", rerr.Error())
					}
					// Below is one way to do; but above is backwards compatible
					//if errors.Is(rerr, os.ErrNotExist) {
					//	fmt.Println("Missing file?")
					//	t.Fatal()
					//}
					t.Fatalf("Unexpected error: %s", rerr.Error())
				}
				// Test TOC structure out ..
				//if !reflect.DeepEqual(actual, want) {
				//	t.Errorf("NewDebateTOC() got = %v, want %v", got, want)
				//}

				if diff := cmp.Diff(actual, want); diff != "" {
					t.Errorf("NewDebateTOC() mismatch (-actual +want):\n%s", diff)
				}
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

// Helper function to get PDF raw content
func loadGoldenDebateTOC(sourcePath string) *debate.DebateTOC {
	// Strip to baseline and load the .golden version
	return nil
}
