package hansard

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Structure copied from Dave Cheney
// https://dave.cheney.net/practical-go/presentations/gophercon-singapore-2019.html#_comparing_expected_an_actual
// TestSplit combines it all
func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"simple":       {input: loadSimple(), sep: "/", want: []string{"a", "b", "c"}},
		"wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		"no sep":       {input: "abc", sep: "/", want: []string{"abc"}},
		"trailing sep": {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
	}

	for name, tc := range tests {
		name := name // If want to run in parallel
		tc := tc     // If want to run in parallel

		t.Run(name, func(st *testing.T) {
			st.Parallel() // Is OK if we make above changes

			got := Split(tc.input, tc.sep)

			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				st.Fatalf(diff)
			}
		})
	}
}

func loadSimple() string {
	return "a/b/c"
}

func Test_detectPossibleQuestionNum(t *testing.T) {
	type args struct {
		linesExcerpt []string
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
			if err := detectPossibleQuestionNum(tt.args.linesExcerpt); (err != nil) != tt.wantErr {
				t.Errorf("detectPossibleQuestionNum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_detectPossibleSessionName(t *testing.T) {
	tests := []struct {
		name    string
		pdfPath string
		want    string
	}{
		{"test #1", "./raw/BukanLisan/bobo.pdf", "Bobo"},
		{"test #2", "./raw/BukanLisan/bobo with space.pdf", "BoboWithSpace"},
		{"test #3", "./raw/BukanLisan/bobo.pdf", "boboSession"},
		{"test #4", ".raw/JawatanKuasa/JKSTUPKK/rumusan-laporan-akhir-jawatankuasa-siasatan-tadbir-urus-perolehan-dan-kewangan-kerajaan-mengenai-projek-land-swap-di-bawah-kementerian-pertahanan_51-59.pdf",
			"boboSession"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectPossibleSessionName(tt.pdfPath); got != tt.want {
				t.Errorf("detectPossibleSessionName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHansardDocument(t *testing.T) {
	type args struct {
		pdfPath string
	}
	tests := []struct {
		name    string
		args    args
		want    *HansardDocument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHansardDocument(tt.args.pdfPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHansardDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHansardDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHansardDocument_ProcessLinesExcerpt(t *testing.T) {
	type args struct {
		linesExcerpt []string
	}
	tests := []struct {
		name    string
		hd      *HansardDocument
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.hd.ProcessLinesExcerpt(tt.args.linesExcerpt); (err != nil) != tt.wantErr {
				t.Errorf("HansardDocument.ProcessLinesExcerpt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHansardDocument_String(t *testing.T) {
	tests := []struct {
		name string
		hd   *HansardDocument
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hd.String()
		})
	}
}

func TestHansardDocument_SplitPDFByQuestions(t *testing.T) {
	tests := []struct {
		name    string
		hd      *HansardDocument
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.hd.SplitPDFByQuestions(); (err != nil) != tt.wantErr {
				t.Errorf("HansardDocument.SplitPDFByQuestions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
