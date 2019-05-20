package hansard_test

import (
	"testing"

	"github.com/Sinar/go-pardocs/internal/hansard"
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

			got := hansard.Split(tc.input, tc.sep)

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

func TestHansardDocument_SplitPDFByQuestions(t *testing.T) {
	type fields struct {
		HansardType      hansard.HansardType
		SessionName      string
		HansardQuestions []hansard.HansardQuestion
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test case #1", fields{hansard.HANSARD_WRITTEN, "PAR14-April-2019", hansard.SetupHansardQuestion("37", 4, 5)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := &hansard.HansardDocument{
				HansardType:      tt.fields.HansardType,
				SessionName:      tt.fields.SessionName,
				HansardQuestions: tt.fields.HansardQuestions,
			}
			if err := hd.SplitPDFByQuestions(); (err != nil) != tt.wantErr {
				t.Errorf("HansardDocument.SplitPDFByQuestions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
