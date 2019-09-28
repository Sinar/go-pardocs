package debate_test

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/google/go-cmp/cmp"

	"github.com/sanity-io/litter"

	"github.com/Sinar/go-pardocs/internal/debate"
)

var update = flag.Bool("update", false, "update .golden files")
var updatePDF = flag.Bool("updatePDF", false, "update .fixture PDFs")

func TestNewDebateTOC(t *testing.T) {
	type args struct {
		fixtureLabel string
		sourcePath   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Missing TOC", args{"TOC-Bad-DR-DewanSelangor", "testdata/Bad-DR-DewanSelangor.pdf"}, true},
		{"TOC empty page2", args{"TOC-DR-01072019", "testdata/DR-01072019.pdf"}, false},
		{"TOC normal #1", args{"TOC-DR-11042019", "testdata/DR-11042019.pdf"}, false},
		{"TOC normal #2", args{"TOC-DR-01072019_new", "testdata/DR-01072019_new.pdf"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := debate.NewDebateTOCPDFContent(loadPDFFromFixture(t, tt.args.fixtureLabel, tt.args.sourcePath))
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

// Helper function to load from fixture; safe/update as per necessary?
func loadPDFFromFixture(t *testing.T, fixtureLabel string, sourcePath string) *debate.PDFDocument {
	// Mark as helper
	t.Helper()

	var pdfDoc *debate.PDFDocument
	// Read from cache; if not exist; complain that need to update
	fixture := filepath.Join("testdata", fixtureLabel+".fixture")
	if *updatePDF {
		// If run update; call the same function used by TOC to get the data
		pdfDoc, err := debate.NewPDFDocForTOC(sourcePath)
		if err != nil {
			t.Fatalf("NewPDFDocForTOC FAIL: %s", err.Error())
		}
		// Persist the data into the file
		w, werr := yaml.Marshal(pdfDoc)
		if werr != nil {
			t.Fatalf("Marshal FAIl: %s", werr.Error())
		}
		ioutil.WriteFile(fixture, w, 0644)
		return pdfDoc
	}
	// Normal path,read from fixture ,..
	want, rerr := ioutil.ReadFile(fixture)
	if rerr != nil {
		// Cannot proceed with one golden file update
		if os.IsNotExist(rerr) {
			t.Fatalf("Ensure run with -updatePDF flag first time! ERR: %s", rerr.Error())
		}
		t.Fatalf("Unexpected error: %s", rerr.Error())
	}
	pdfDoc = &debate.PDFDocument{}
	umerr := yaml.Unmarshal(want, pdfDoc)
	if umerr != nil {
		t.Fatalf("Unmarshal FAIl: %s", umerr.Error())
	}
	return pdfDoc
}
