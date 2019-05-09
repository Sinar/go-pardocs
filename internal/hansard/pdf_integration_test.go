package hansard_test

import (
	"testing"

	"github.com/Sinar/go-pardocs/internal/hansard"
	"github.com/google/go-cmp/cmp"
)

// Structure copied from Dave Cheney
// https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_comparing_expected_an_actual
// TestSplit combines it all
func TestPDFTxtStyles(t *testing.T) {
	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: loadSimply(), sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for name, tc := range tests {
		name := name // If want to run in parallel
		tc := tc     // If want to run in parallel
		t.Run(name, func(st *testing.T) {
			st.Parallel() // Is OK if we make above changes

			got := hansard.Split(tc.input, tc.sep)

			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				st.Fatalf(diff)
			}
		})
	}
}

func loadSimply() string {
	return "a/b/c"
}

func TestExtractPDF(t *testing.T) {
	type args struct {
		pdfDoc  *hansard.PDFDocument
		pdfPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := hansard.ExtractPDF(tt.args.pdfDoc, tt.args.pdfPath); (err != nil) != tt.wantErr {
				t.Errorf("ExtractPDF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
