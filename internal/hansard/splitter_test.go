package hansard

import "testing"

func TestSplitPlan_ExecuteSplit(t *testing.T) {
	type fields struct {
		QuestionNum  string
		PageNumStart int
		PageNumEnd   int
	}
	tests := []struct {
		name   string
		label  string
		fields fields
	}{
		{"test #1", "par14-sesi1", fields{"3", 5, 6}},
	}
	// Output file: <Label>-soalan-<QuestionNum>.pdf
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := &SplitPlan{
				QuestionNum:  tt.fields.QuestionNum,
				PageNumStart: tt.fields.PageNumStart,
				PageNumEnd:   tt.fields.PageNumEnd,
			}
			sp.ExecuteSplit(tt.label)
		})
	}
}
