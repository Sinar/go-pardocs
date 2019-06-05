package hansard

import (
	"fmt"
	"io/ioutil"

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
	DestSplitPDFs    string // Directory to store the final split items; default to ./data/<HansardType>/<SessionName>/
	SplitPlans       []SplitPlan
}

type SplitPlan struct {
	QuestionNum  string
	PageNumStart int
	PageNumEnd   int
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
	// TODO: Read the plan file

	b, rerr := ioutil.ReadFile(planFilename)
	if rerr != nil {
		panic(rerr)
	}
	umerr := yaml.Unmarshal(b, &splitHansardDocument.SplitPlans)
	if umerr != nil {
		panic(umerr)
	}
	return &splitHansardDocument
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

	proceedPrepare := rawDataFolderSetup(fmt.Sprintf("./raw/splitout/%s/%s/pages/", hansardType, shd.SessionName))
	if proceedPrepare {
		shd.PrepareSplit()
	}

	for _, sp := range shd.SplitPlans {
		sp.ExecuteSplit(shd.SessionName)
	}
}

// TODO: Refactor to local and group it out ..
func (shd *SplitHansardDocument) PrepareSplit() {
	// If no pdf, append PDF
	// check actual type via MIME? or ext?
	// Build out the full path ..
	//rawBasePath := "/Users/mleow/GOMOD/go-pardocs/raw/"
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
	cmd := papi.SplitCommand(shd.OriginalPDFPath, fmt.Sprintf("./raw/splitout/%s/%s/pages/", hansardType, shd.SessionName), 1, pdfcpu.NewDefaultConfiguration())
	o, perr := papi.Process(cmd)
	if perr != nil {
		panic(perr)
	}
	// What is the output??
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

func (sp *SplitPlan) ExecuteSplit(label string) {
	// TODO: Guardrail; check first that Prepare split is already there; full doc split out
	// into /tmp/split/<file_basename>/<file_basename>_<pagenum>/pdf

	outputFilename := fmt.Sprintf("%s-soalan-%s-bukanlisan", label, sp.QuestionNum)
	fmt.Println("====== ", outputFilename, " =======")

	var pagesToMerge []string

	for i := sp.PageNumStart; i <= sp.PageNumEnd; i++ {
		pagesToMerge = append(pagesToMerge, fmt.Sprintf("/tmp/BukanLisan/Pertanyaan Jawapan Bukan Lisan 22019_new_%d.pdf", i))
	}
	q.Q(pagesToMerge)

	cmd := papi.MergeCommand(pagesToMerge, "/tmp/BukanLisan/EXTRACT.pdf", pdfcpu.NewDefaultConfiguration())
	o, merr := papi.Process(cmd)
	if merr != nil {
		panic(merr)
	}
	q.Q(o)

}
