package hansard_test

import (
	"testing"

	"github.com/Sinar/go-pardocs/internal/hansard"
)

func TestSplitHansardDocument_PrepareExecuteSplit(t *testing.T) {
	type fields struct {
		HansardType     hansard.HansardType
		SessionName     string
		OriginalPDFPath string
		DestSplitPDFs   string
		SplitPlans      []hansard.SplitPlan
	}
	type args struct {
		destSplitPDFs string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shd := &hansard.SplitHansardDocument{
				HansardType:     tt.fields.HansardType,
				SessionName:     tt.fields.SessionName,
				OriginalPDFPath: tt.fields.OriginalPDFPath,
				DestSplitPDFs:   tt.fields.DestSplitPDFs,
				SplitPlans:      tt.fields.SplitPlans,
			}
			if err := shd.PrepareExecuteSplit(tt.args.destSplitPDFs); (err != nil) != tt.wantErr {
				t.Errorf("SplitHansardDocument.PrepareExecuteSplit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
			sp := &hansard.SplitPlan{
				QuestionNum:  tt.fields.QuestionNum,
				PageNumStart: tt.fields.PageNumStart,
				PageNumEnd:   tt.fields.PageNumEnd,
			}
			sp.ExecuteSplit(tt.label)
			// TODO: check the ordering and key output is correct
		})
	}
}

func TestSplitPlan_PrepareSplit(t *testing.T) {
	type args struct {
		originalFilename string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test #1", args{"./BukanLisan/Pertanyaan Jawapan Bukan Lisan 22019_new.pdf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := &hansard.SplitPlan{}
			sp.PrepareSplit(tt.args.originalFilename)
			// TODO: count number of split pages
		})
	}
}
