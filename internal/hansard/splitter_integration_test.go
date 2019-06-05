package hansard_test

import (
	"reflect"
	"testing"

	"github.com/Sinar/go-pardocs/internal/hansard"
)

func TestNewSplitHansardDocument(t *testing.T) {
	type args struct {
		planFilename    string
		hansardType     hansard.HansardType
		sessionName     string
		originalPDFPath string
	}
	tests := []struct {
		name string
		args args
		want *hansard.SplitHansardDocument
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hansard.NewSplitHansardDocument(tt.args.planFilename, tt.args.hansardType, tt.args.sessionName, tt.args.originalPDFPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSplitHansardDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//shd := &hansard.SplitHansardDocument{
			//	HansardType:     tt.fields.HansardType,
			//	SessionName:     tt.fields.SessionName,
			//	OriginalPDFPath: tt.fields.OriginalPDFPath,
			//	DestSplitPDFs:   tt.fields.DestSplitPDFs,
			//	SplitPlans:      tt.fields.SplitPlans,
			//}
			//shd.PrepareExecuteSplit(tt.args.destSplitPDFs)
		})
	}
}

func TestSplitHansardDocument_PrepareSplit(t *testing.T) {
	type args struct {
		label             string
		currentworkingDir string
		planFilename      string
		hansardType       hansard.HansardType
		sessionName       string
		originalPDFPath   string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test #1", args{"par14sesi1", "/Users/mleow/GOMOD/go-pardocs/",
			"data/BukanLisan/split.yml", hansard.HANSARD_WRITTEN,
			"Pertanyaan Jawapan Bukan Lisan 22019_new",
			"raw/BukanLisan/Pertanyaan Jawapan Bukan Lisan 22019_new.pdf",
		}},
		//{"test #1", args{"./data/BukanLisan/split.yml", hansard.HANSARD_WRITTEN, "", "/Users/mleow/GOMOD/go-pardocs/raw/"}},
		//{"test #1", args{"./data/BukanLisan/split.yml", hansard.HANSARD_WRITTEN, "", "/Users/mleow/GOMOD/go-pardocs/raw/"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shd := hansard.NewSplitHansardDocument(tt.args.label, tt.args.currentworkingDir,
				tt.args.planFilename, tt.args.hansardType, tt.args.sessionName, tt.args.originalPDFPath)
			shd.SplitPlans = hansard.NewMockSplitPlan()
			shd.PrepareExecuteSplit()
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

//func TestSplitPlan_PrepareSplit(t *testing.T) {
//	type args struct {
//		originalFilename string
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		{"test #1", args{"./BukanLisan/Pertanyaan Jawapan Bukan Lisan 22019_new.pdf"}},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			sp := &hansard.SplitPlan{}
//			sp.PrepareSplit(tt.args.originalFilename)
//			// TODO: count number of split pages
//		})
//	}
//}
