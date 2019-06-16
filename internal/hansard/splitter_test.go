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

//func TestSplitHansardDocumentPlan_Setup(t *testing.T) {
//	type fields struct {
//		sessionName   string
//		hansardType   string
//		workingDir    string
//		sourcePDFPath string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			hsdp := &SplitHansardDocumentPlan{
//				sessionName:   tt.fields.sessionName,
//				hansardType:   tt.fields.hansardType,
//				workingDir:    tt.fields.workingDir,
//				sourcePDFPath: tt.fields.sourcePDFPath,
//			}
//			hsdp.Setup()
//		})
//	}
//}

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
			// TODO: Check the plan is there as expected; also have failure scenario to catch errors?
		})
	}
}

func TestSetupSplitPlanFixture(t *testing.T) {
	type args struct {
		testDir       string
		fixtureDir    string
		scenarioDir   string
		sourcePDFPath string
		ht            HansardType
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
			if err := SetupSplitPlanFixture(tt.args.testDir, tt.args.fixtureDir, tt.args.scenarioDir, tt.args.sourcePDFPath, tt.args.ht); (err != nil) != tt.wantErr {
				t.Errorf("SetupSplitPlanFixture() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
