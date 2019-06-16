package pardocs_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Sinar/go-pardocs"
	"github.com/Sinar/go-pardocs/internal/hansard"
)

func TestParliamentDocs_Plan(t *testing.T) {
	tests := []struct {
		name string
		pd   *pardocs.ParliamentDocs
	}{
		// Too huge!!
		{"test #1", &pardocs.ParliamentDocs{pardocs.Configuration{
			"par14sesi1", hansard.HANSARD_SPOKEN,
			".", "./raw/Lisan/JDR12032019.pdf", pardocs.PLAN}}},
		//{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := ioutil.TempDir("", "pardocs")
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dir)
			log.Println("Dir is ", dir)
			tt.pd.Conf.WorkingDir = dir
			tt.pd.Plan()

			// Let's check
			plan := hansard.LoadSplitHansardDocPlanFromFile(tt.pd.Conf.HansardType, tt.pd.Conf.WorkingDir, tt.pd.Conf.SourcePDFPath)
			if plan.ParliamentSession != tt.pd.Conf.ParliamentSession {
				t.Fail()
			}
			if plan.HansardType != tt.pd.Conf.HansardType {
				t.Fail()
			}
			for _, q := range plan.HansardQuestions {
				// DEBUG
				//spew.Dump(q)
				// Check start
				if q.QuestionNum == "1" {
					if q.PageNumStart != 2 || q.PageNumEnd != 5 {
						log.Println("Wrong Q1")
						t.Fail()
					}
				}
				// Check middle
				if q.QuestionNum == "20" {
					if q.PageNumStart != 65 || q.PageNumEnd != 68 {
						log.Println("Wrong Q20")
						t.Fail()
					}
				}
				// Check end
				if q.QuestionNum == "82" {
					if q.PageNumStart != 208 || q.PageNumEnd != 210 {
						log.Println("Wrong Q82")
						t.Fail()
					}
				}
			}
			//plan.HansardQuestions
		})
	}
}

func TestParliamentDocs_Split(t *testing.T) {
	type fields struct {
		scenarioDir string
		Conf        pardocs.Configuration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"test #1", fields{"happy1", pardocs.Configuration{"par14sesi1", hansard.HANSARD_SPOKEN, ".", "./raw/Lisan/JDR12032019.pdf", pardocs.SPLIT}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := &pardocs.ParliamentDocs{
				Conf: tt.fields.Conf,
			}
			// Setup the temp testing Dir ..
			dir, err := ioutil.TempDir("", "pardocs")
			if err != nil {
				log.Fatal(err)
			}
			// Comment out below if need to see the output in dir
			defer os.RemoveAll(dir)
			log.Println("Dir is ", dir)

			// Setup the  workingDir patch
			pd.Conf.WorkingDir = dir
			// Setup fixture from testdata folder
			absSourcePDFPath, _ := filepath.Abs(tt.fields.Conf.SourcePDFPath)
			// Bad  hack ...
			pd.Conf.SourcePDFPath = absSourcePDFPath
			absFixtureDirPath, _ := filepath.Abs("./testdata")
			ferr := hansard.SetupSplitPlanFixture(dir, absFixtureDirPath, tt.fields.scenarioDir,
				absSourcePDFPath, tt.fields.Conf.HansardType)
			if ferr != nil {
				t.Fatal(ferr)
			}

			// Execute split
			pd.Split()

			// Check the split is correct
			// First question
			// mid question
			// last question
		})
	}
}

func TestParliamentDocs_Reset(t *testing.T) {
	type fields struct {
		Conf pardocs.Configuration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pd := &pardocs.ParliamentDocs{
				Conf: tt.fields.Conf,
			}
			pd.Reset()
		})
	}
}

// Helper functions
//func getParliamentDocMetadata(pdfPath string, ht hansard.HansardType) (sessionName string, hansardType string) {
//	baseFilename := filepath.Base(pdfPath)
//	sessionName = strings.Split(baseFilename, ".")[0]
//	switch ht {
//	case hansard.HANSARD_SPOKEN:
//		hansardType = "Lisan"
//	case hansard.HANSARD_WRITTEN:
//		hansardType = "BukanLisan"
//	default:
//		panic(fmt.Errorf("INVALID TYPE!!!"))
//	}
//
//	return sessionName, hansardType
//}
