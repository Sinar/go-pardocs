package hansard

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"

	papi "github.com/hhrutter/pdfcpu/pkg/api"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/y0ssar1an/q"
)

type SplitHansardDocument struct {
	Label            string
	HansardType      HansardType
	SessionName      string // Get this from the front page cover .. or the reference lookup ..
	WorkingDirectory string // defaults to cwd if not defined ..
	OriginalPDFPath  string // Used for split later ..
	DestSplitPDFs    string // Directory to store the final split items; default to ./data/<HansardType>/<ParliamentSession>/
	SplitPlans       []SplitPlan
}

type SplitPlan struct {
	QuestionNum  string
	PageNumStart int
	PageNumEnd   int
}

type SplitHansardDocumentPlan struct {
	sessionName   string
	hansardType   string
	workingDir    string
	sourcePDFPath string
}

func NewSplitHansardDocumentPlan(confHansardType HansardType, workingDir string, sourcePDFPath string) *SplitHansardDocumentPlan {
	sessionName, hansardType := getParliamentDocMetadata(sourcePDFPath, confHansardType)

	splitHansardDocPlan := SplitHansardDocumentPlan{
		sessionName:   sessionName,
		hansardType:   hansardType,
		workingDir:    workingDir,
		sourcePDFPath: sourcePDFPath,
	}
	return &splitHansardDocPlan
}

func LoadSplitHansardDocPlanFromFile(confHansardType HansardType, workingDir string, sourcePDFPath string) *HansardDocument {
	sessionName, hansardType := getParliamentDocMetadata(sourcePDFPath, confHansardType)
	splitPlanPath := fmt.Sprintf("%s/data/%s/%s/split.yml", workingDir, hansardType, sessionName)
	return loadSplitHansardDocPlan(splitPlanPath)
}

func loadSplitHansardDocPlan(splitPlanPath string) *HansardDocument {
	splitHansardDocPlan := HansardDocument{}

	// Read the plan file
	b, rerr := ioutil.ReadFile(splitPlanPath)
	if rerr != nil {
		panic(rerr)
	}
	umerr := yaml.Unmarshal(b, &splitHansardDocPlan)
	if umerr != nil {
		panic(umerr)
	}

	return &splitHansardDocPlan
}

// NewSplitPlan will use a Reader (better!) to extract out the plan
func NewSplitHansardDocument(label string, currentWorkingDir string, planFilename string, hansardType HansardType, sessionName string, originalPDFPath string) *SplitHansardDocument {
	splitHansardDocument := SplitHansardDocument{
		Label:            label,
		HansardType:      hansardType,
		SessionName:      sessionName,
		WorkingDirectory: currentWorkingDir,
		OriginalPDFPath:  originalPDFPath,
	}
	// Use mock for simpler cases ..
	splitHansardDocument.SplitPlans = NewMockSplitPlan()

	// Read the plan file
	b, rerr := ioutil.ReadFile(fmt.Sprintf("%s/%s", currentWorkingDir, planFilename))
	if rerr != nil {
		panic(rerr)
	}
	umerr := yaml.Unmarshal(b, &splitHansardDocument.SplitPlans)
	if umerr != nil {
		panic(umerr)
	}
	return &splitHansardDocument
}

func getParliamentDocMetadata(pdfPath string, ht HansardType) (sessionName string, hansardType string) {
	baseFilename := filepath.Base(pdfPath)
	sessionName = strings.Split(baseFilename, ".")[0]
	switch ht {
	case HANSARD_SPOKEN:
		hansardType = "Lisan"
	case HANSARD_WRITTEN:
		hansardType = "BukanLisan"
	default:
		panic(fmt.Errorf("INVALID TYPE!!!"))
	}

	return sessionName, hansardType
}

func SavePlan(confHansardType HansardType, workingDir string, sourcePDFPath string, hansardDoc *HansardDocument) {

	sessionName, hansardType := getParliamentDocMetadata(sourcePDFPath, confHansardType)
	hansardDoc.PersistForSplit(fmt.Sprintf("%s/data/%s/%s", workingDir, hansardType, sessionName))

}

func (shd *SplitHansardDocument) PrepareExecuteSplit() {
	// If the split document does not exist; call the Prepare
	var hansardType string
	switch shd.HansardType {
	case HANSARD_WRITTEN:
		hansardType = "BukanLisan"

	case HANSARD_SPOKEN:
		hansardType = "Lisan"

	default:
		panic(fmt.Errorf("Incorrect TYPE: %#v", shd.HansardType))
	}

	proceedPrepare := rawDataFolderSetup(fmt.Sprintf("%s/raw/splitout/%s/%s/pages/", shd.WorkingDirectory, hansardType, shd.SessionName))
	if proceedPrepare {
		fmt.Println("OK; splitting!!")
		shd.PrepareSplit()
	}

	for _, sp := range shd.SplitPlans {
		sp.ExecuteSplit(shd.WorkingDirectory, hansardType, shd.SessionName, shd.Label)
	}
}

// TODO: Refactor to local and group it out ..
func (shd *SplitHansardDocument) PrepareSplit() {
	fmt.Println("In shd.PrepareSplit ..")
	// If no pdf, append PDF
	// check actual type via MIME? or ext?
	// Build out the full path ..
	//rawBasePath := "/Users/mleow/GOMOD/go-go-pardocs/raw/"
	//
	var hansardType string
	switch shd.HansardType {
	case HANSARD_WRITTEN:
		hansardType = "BukanLisan"

	case HANSARD_SPOKEN:
		hansardType = "Lisan"

	default:
		panic(fmt.Errorf("Incorrect TYPE: %#v", shd.HansardType))
	}
	// Assumes created ?
	cmd := papi.SplitCommand(fmt.Sprintf("%s/%s", shd.WorkingDirectory, shd.OriginalPDFPath),
		fmt.Sprintf("%s/raw/splitout/%s/%s/pages/", shd.WorkingDirectory, hansardType, shd.SessionName),
		1, pdfcpu.NewDefaultConfiguration())
	o, perr := papi.Process(cmd)
	if perr != nil {
		panic(perr)
	}
	// What is the output??
	// DEBUG
	q.Q(o)
}

func extractCoverPage(hansardType HansardType, originalPDFPath string) []string {
	var contentCoverPage []string

	// Based on the type of document; might have different interpretations of Cover Page
	// e.g state assembly might have Attendance Sheet? and/or proper ToC?

	return contentCoverPage
}

func detectSessionName(hansardType HansardType, sourcePDFFileName string, contentCoverPage []string) string {
	// Guardrail; check no pdf; and is not some sort of basefilename?

	// Look for common pattern based on the type ..

	// Can;t find aything; just return sourcePDFFileName! with [^[:alphanum:]] replaced with single -
	return sourcePDFFileName
}

// NewMockSplitPlan returns
// 	portion of actual test case file PDF
func NewMockSplitPlan() []SplitPlan {
	return []SplitPlan{
		{"1", 2, 3},
		{"2", 4, 4},
		{"3", 5, 6},
		{"4", 7, 8},
		{"5", 9, 9},
		{"6", 10, 12},
	}
}

func (sp *SplitPlan) ExecuteSplit(currentWorkingDir string, hansardType string, sessionName string, label string) {
	// TODO: Guardrail; check first that Prepare split is already there; full doc split out
	// into /tmp/split/<file_basename>/<file_basename>_<pagenum>/pdf

	var pagesToMerge []string

	for i := sp.PageNumStart; i <= sp.PageNumEnd; i++ {
		sourcePDFPath := fmt.Sprintf("%s/raw/splitout/%s/%s/pages/%s_%d.pdf", currentWorkingDir, hansardType, sessionName, sessionName, i)
		pagesToMerge = append(pagesToMerge, sourcePDFPath)
	}
	// DEBUG
	//q.Q(pagesToMerge)

	finalMergedPDFPath := fmt.Sprintf("%s/splitout/%s-soalan-%s-%s.pdf", currentWorkingDir, label, hansardType, sp.QuestionNum)
	fmt.Println(">>>=========== Merged file at: ", finalMergedPDFPath, " ==============<<<<<<")
	cmd := papi.MergeCommand(pagesToMerge, finalMergedPDFPath, pdfcpu.NewDefaultConfiguration())
	o, merr := papi.Process(cmd)
	if merr != nil {
		panic(merr)
	}
	// DEBUG
	q.Q(o)

}

func (hsdp *SplitHansardDocumentPlan) Setup() {

	// Prepare final location
	proceedPrepare := rawDataFolderSetup(fmt.Sprintf("%s/raw/splitout/%s/%s/pages/", hsdp.workingDir, hsdp.hansardType, hsdp.sessionName))
	if proceedPrepare {
		fmt.Println("OK; splitting!!")
		prepareSplit(hsdp.sessionName, hsdp.hansardType, hsdp.workingDir, hsdp.sourcePDFPath)
	}

}

func prepareSplit(sessionName string, hansardType string, workingDir string, sourcePDFPath string) {

	cmd := papi.SplitCommand(sourcePDFPath,
		fmt.Sprintf("%s/raw/splitout/%s/%s/pages/", workingDir, hansardType, sessionName),
		1, pdfcpu.NewDefaultConfiguration())
	o, perr := papi.Process(cmd)
	if perr != nil {
		panic(perr)
	}
	// DEBUG
	q.Q(o)
}

func (shdp *SplitHansardDocumentPlan) ExecuteSplit(label string, hq HansardQuestion) {
	currentWorkingDir := shdp.workingDir
	hansardType := shdp.hansardType
	sessionName := shdp.sessionName

	// TODO: Guardrail; check first that Prepare split is already there; full doc split out
	// into /tmp/split/<file_basename>/<file_basename>_<pagenum>/pdf

	var pagesToMerge []string

	for i := hq.PageNumStart; i <= hq.PageNumEnd; i++ {
		sourcePDFPath := fmt.Sprintf("%s/raw/splitout/%s/%s/pages/%s_%d.pdf", currentWorkingDir, hansardType, sessionName, sessionName, i)
		pagesToMerge = append(pagesToMerge, sourcePDFPath)
	}
	// DEBUG
	//q.Q(pagesToMerge)

	finalMergedPDFPath := fmt.Sprintf("%s/splitout/%s-soalan-%s-%s.pdf", currentWorkingDir, label, hansardType, hq.QuestionNum)
	fmt.Println(">>>=========== Merged file at: ", finalMergedPDFPath, " ==============<<<<<<")
	cmd := papi.MergeCommand(pagesToMerge, finalMergedPDFPath, pdfcpu.NewDefaultConfiguration())
	o, merr := papi.Process(cmd)
	if merr != nil {
		panic(merr)
	}
	// DEBUG
	q.Q(o)

}
