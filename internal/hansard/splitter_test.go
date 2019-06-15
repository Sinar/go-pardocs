package hansard

import (
	"reflect"
	"testing"
)

func Test_detectSessionName(t *testing.T) {
	type args struct {
		hansardType        HansardType
		sourcePDFFileName  string
		contentofFirstPage []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectSessionName(tt.args.hansardType, tt.args.sourcePDFFileName, tt.args.contentofFirstPage); got != tt.want {
				t.Errorf("detectSessionName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractCoverPage(t *testing.T) {
	type args struct {
		hansardType     HansardType
		originalPDFPath string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractCoverPage(tt.args.hansardType, tt.args.originalPDFPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractCoverPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prepareSplit(t *testing.T) {
	type args struct {
		sessionName   string
		hansardType   string
		workingDir    string
		sourcePDFPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prepareSplit(tt.args.sessionName, tt.args.hansardType, tt.args.workingDir, tt.args.sourcePDFPath)
		})
	}
}

func TestSplitHansardDocumentPlan_Setup(t *testing.T) {
	type fields struct {
		sessionName   string
		hansardType   string
		workingDir    string
		sourcePDFPath string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hsdp := &SplitHansardDocumentPlan{
				sessionName:   tt.fields.sessionName,
				hansardType:   tt.fields.hansardType,
				workingDir:    tt.fields.workingDir,
				sourcePDFPath: tt.fields.sourcePDFPath,
			}
			hsdp.Setup()
		})
	}
}

func TestSplitHansardDocument_PrepareExecuteSplit(t *testing.T) {
	type fields struct {
		Label            string
		HansardType      HansardType
		SessionName      string
		WorkingDirectory string
		OriginalPDFPath  string
		DestSplitPDFs    string
		SplitPlans       []SplitPlan
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shd := &SplitHansardDocument{
				Label:            tt.fields.Label,
				HansardType:      tt.fields.HansardType,
				SessionName:      tt.fields.SessionName,
				WorkingDirectory: tt.fields.WorkingDirectory,
				OriginalPDFPath:  tt.fields.OriginalPDFPath,
				DestSplitPDFs:    tt.fields.DestSplitPDFs,
				SplitPlans:       tt.fields.SplitPlans,
			}
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
	type args struct {
		currentWorkingDir string
		hansardType       string
		sessionName       string
		label             string
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
			sp := &SplitPlan{
				QuestionNum:  tt.fields.QuestionNum,
				PageNumStart: tt.fields.PageNumStart,
				PageNumEnd:   tt.fields.PageNumEnd,
			}
			sp.ExecuteSplit(tt.args.currentWorkingDir, tt.args.hansardType, tt.args.sessionName, tt.args.label)
		})
	}
}

func TestSplitHansardDocumentPlan_ExecuteSplit(t *testing.T) {
	type fields struct {
		sessionName   string
		hansardType   string
		workingDir    string
		sourcePDFPath string
	}
	type args struct {
		label string
		hq    HansardQuestion
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
			shdp := &SplitHansardDocumentPlan{
				sessionName:   tt.fields.sessionName,
				hansardType:   tt.fields.hansardType,
				workingDir:    tt.fields.workingDir,
				sourcePDFPath: tt.fields.sourcePDFPath,
			}
			shdp.ExecuteSplit(tt.args.label, tt.args.hq)
		})
	}
}

func Test_getParliamentDocMetadata(t *testing.T) {
	type args struct {
		pdfPath string
		ht      HansardType
	}
	tests := []struct {
		name            string
		args            args
		wantSessionName string
		wantHansardType string
	}{
		{"test #1", args{"./raw/BukanLisan/Pertanyaan Jawapan Bukan Lisan 22019_new.pdf",
			HANSARD_WRITTEN},
			"Pertanyaan Jawapan Bukan Lisan 22019_new", "BukanLisan"},
		{"test #2", args{"/tmp/Parlimen_Lisan_22019.pdf",
			HANSARD_SPOKEN}, "Parlimen_Lisan_22019",
			"Lisan"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSessionName, gotHansardType := getParliamentDocMetadata(tt.args.pdfPath, tt.args.ht)
			if gotSessionName != tt.wantSessionName {
				t.Errorf("getParliamentDocMetadata() gotSessionName = %v, want %v", gotSessionName, tt.wantSessionName)
			}
			if gotHansardType != tt.wantHansardType {
				t.Errorf("getParliamentDocMetadata() gotHansardType = %v, want %v", gotHansardType, tt.wantHansardType)
			}
		})
	}
}

func TestSavePlan(t *testing.T) {
	type args struct {
		confHansardType HansardType
		workingDir      string
		sourcePDFPath   string
		hansardDoc      *HansardDocument
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SavePlan(tt.args.confHansardType, tt.args.workingDir, tt.args.sourcePDFPath, tt.args.hansardDoc)
		})
	}
}

func TestSplitHansardDocument_PrepareSplit(t *testing.T) {
	type fields struct {
		Label            string
		HansardType      HansardType
		SessionName      string
		WorkingDirectory string
		OriginalPDFPath  string
		DestSplitPDFs    string
		SplitPlans       []SplitPlan
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shd := &SplitHansardDocument{
				Label:            tt.fields.Label,
				HansardType:      tt.fields.HansardType,
				SessionName:      tt.fields.SessionName,
				WorkingDirectory: tt.fields.WorkingDirectory,
				OriginalPDFPath:  tt.fields.OriginalPDFPath,
				DestSplitPDFs:    tt.fields.DestSplitPDFs,
				SplitPlans:       tt.fields.SplitPlans,
			}
			shd.PrepareSplit()
		})
	}
}
