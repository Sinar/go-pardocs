package hansard_test

import (
	"reflect"
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

//func TestExtractPDF(t *testing.T) {
//	mypdfDoc := hansard.PDFDocument{}
//	type args struct {
//		pdfDoc  *hansard.PDFDocument
//		pdfPath string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		{"test #1", args{&mypdfDoc, ""}, true},
//		{"test #2", args{&mypdfDoc, ""}, true},
//		{"test #3", args{&mypdfDoc, ""}, true},
//		{"test #4", args{&mypdfDoc, ""}, true},
//		{"test #5", args{&mypdfDoc, ""}, true},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := hansard.ExtractPDF(tt.args.pdfDoc, tt.args.pdfPath); (err != nil) != tt.wantErr {
//				t.Errorf("ExtractPDF() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func TestNewPDFDoc(t *testing.T) {
	mypdfDoc := hansard.PDFDocument{}
	tests := []struct {
		name    string
		fixture string
		want    *hansard.PDFDocument
		wantErr bool
	}{
		{"test #1", "testdata/abc.pdf", &mypdfDoc, true},
		{"test #2", "testdata/abc.pdf", &mypdfDoc, true},
		{"test #3", "testdata/abc.pdf", &mypdfDoc, true},
		{"test #4", "testdata/abc.pdf", &mypdfDoc, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hansard.NewPDFDoc(tt.fixture)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPDFDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPDFDoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestPDFDocument_ExtractPDF(t *testing.T) {
//	type fields struct {
//		NumPages   int
//		Pages      []hansard.PDFPage
//		sourcePath string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//	}{
//		{"test #1", fields{}, true},
//		{"test #2", fields{}, true},
//		{"test #3", fields{}, true},
//		{"test #4", fields{}, true},
//		{"test #5", fields{}, true},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			pdfDoc := &hansard.PDFDocument{
//				NumPages: tt.fields.NumPages,
//				Pages:    tt.fields.Pages,
//			}
//			if err := pdfDoc.ExtractPDF(); (err != nil) != tt.wantErr {
//				t.Errorf("PDFDocument.ExtractPDF() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
