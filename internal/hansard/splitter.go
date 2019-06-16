package hansard

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	papi "github.com/hhrutter/pdfcpu/pkg/api"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/y0ssar1an/q"
)

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

func SavePlan(confHansardType HansardType, workingDir string, sourcePDFPath string, hansardDoc *HansardDocument) {

	sessionName, hansardType := getParliamentDocMetadata(sourcePDFPath, confHansardType)
	hansardDoc.PersistForSplit(fmt.Sprintf("%s/data/%s/%s", workingDir, hansardType, sessionName))

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

func LoadSplitHansardDocPlanFromFile(confHansardType HansardType, workingDir string, sourcePDFPath string) *HansardDocument {
	sessionName, hansardType := getParliamentDocMetadata(sourcePDFPath, confHansardType)
	splitPlanPath := fmt.Sprintf("%s/data/%s/%s/split.yml", workingDir, hansardType, sessionName)
	return loadSplitHansardDocPlan(splitPlanPath)
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

//func (hsdp *SplitHansardDocumentPlan) Setup() {
//
//	// Prepare final location
//	proceedPrepare := rawDataFolderSetup(fmt.Sprintf("%s/raw/splitout/%s/%s/pages/", hsdp.workingDir, hsdp.hansardType, hsdp.sessionName))
//	if proceedPrepare {
//		fmt.Println("OK; splitting!!")
//		prepareSplit(hsdp.sessionName, hsdp.hansardType, hsdp.workingDir, hsdp.sourcePDFPath)
//	}
//
//}

// Use the CLI method which is tested to work ..
func prepareSplit(sessionName string, hansardType string, workingDir string, sourcePDFPath string) {
	// pdfcpu split <sourcePDF> <destDir>
	// e.g.
	// $ pdfcpu split ./raw/Lisan/JDR12032019.pdf /var/folders/_p/qk0rf40514b4sgy16r5qyxs40000gn/T/pardocs819115744/raw/splitout/Lisan/JDR12032019/pages

	// Assumes  user has pdfcpu  in the normal go bin dir ..
	u, uerr := user.Current()
	if uerr != nil {
		panic(uerr)
	}
	cmdPath := fmt.Sprintf("%s/go/bin/pdfcpu", u.HomeDir)
	destPDFDir := fmt.Sprintf("%s/raw/splitout/%s/%s/pages/", workingDir, hansardType, sessionName)

	cmd := exec.Command(cmdPath, "split", sourcePDFPath, destPDFDir)
	// DEBUG
	//spew.Dump(cmd)
	o, exerr := cmd.CombinedOutput()
	if exerr != nil {
		panic(exerr)
	}
	//  DEBUG ..
	q.Q(string(o))
}

//  Cannot use via API  until https://github.com/hhrutter/pdfcpu/issues/87 resolved
func prepareSplitAPI(sessionName string, hansardType string, workingDir string, sourcePDFPath string) {

	// Relax validation  --> https://github.com/hhrutter/pdfcpu/issues/80
	conf := pdfcpu.NewDefaultConfiguration()
	// Not needed as it is actually relaxed by default :(
	conf.ValidationMode = pdfcpu.ValidationRelaxed
	//  DEBUG
	//	fmt.Println("VALIDATION: ", conf.ValidationModeString())
	destPDFDir := fmt.Sprintf("%s/raw/splitout/%s/%s/pages/", workingDir, hansardType, sessionName)
	cmd := papi.SplitCommand(sourcePDFPath, destPDFDir,
		1, conf)
	o, perr := papi.Process(cmd)
	if perr != nil {
		panic(perr)
	}
	// DEBUG
	q.Q(o)
}

func (shdp *SplitHansardDocumentPlan) ExecuteSplit(label string, hq HansardQuestion) {
	// TODO: Guardrail; check first that Prepare split is already there; full doc split out
	// into /tmp/split/<file_basename>/<file_basename>_<pagenum>/pdf
	// Setup local variables; need sto  be  better?
	currentWorkingDir := shdp.workingDir
	hansardType := shdp.hansardType
	sessionName := shdp.sessionName
	// Prepare final location
	proceedPrepare := rawDataFolderSetup(fmt.Sprintf("%s/raw/splitout/%s/%s/pages/", currentWorkingDir,
		hansardType, sessionName))
	if proceedPrepare {
		fmt.Println("OK; splitting!!")
		prepareSplit(sessionName, hansardType, currentWorkingDir, shdp.sourcePDFPath)
	}

	// Pre-reqs are done; now can start the split itself ..
	var pagesToMerge []string

	for i := hq.PageNumStart; i <= hq.PageNumEnd; i++ {
		sourcePDFPath := fmt.Sprintf("%s/raw/splitout/%s/%s/pages/%s_%d.pdf", currentWorkingDir, hansardType, sessionName, sessionName, i)
		pagesToMerge = append(pagesToMerge, sourcePDFPath)
	}
	// DEBUG
	//q.Q(pagesToMerge)

	// Ensure the merged directory is there ..
	rawDataFolderSetup(fmt.Sprintf("%s/splitout", currentWorkingDir))
	finalMergedPDFPath := fmt.Sprintf("%s/splitout/%s-soalan-%s-%s.pdf", currentWorkingDir, label, hansardType, hq.QuestionNum)
	fmt.Println(">>>=========== Merged file at: ", finalMergedPDFPath, " ==============<<<<<<")

	// Relax validation  --> https://github.com/hhrutter/pdfcpu/issues/80
	// Real-life data are pretty broken ..
	conf := pdfcpu.NewDefaultConfiguration()
	// Not needed
	//conf.ValidationMode = pdfcpu.ValidationRelaxed
	cmd := papi.MergeCommand(pagesToMerge, finalMergedPDFPath, conf)
	o, merr := papi.Process(cmd)
	if merr != nil {
		panic(merr)
	}
	// DEBUG
	q.Q(o)

}

// Helper Functions for Split Testing
func SetupSplitPlanFixture(testDir string, fixtureDir string, scenarioDir string, sourcePDFPath string, ht HansardType) error {
	// Read the fixture data
	b, rerr := ioutil.ReadFile(fmt.Sprintf("%s/%s/split.yml", fixtureDir, scenarioDir))
	if rerr != nil {
		return rerr
	}
	// Copy fixcture data over to testDir
	sessionName, hansardType := getParliamentDocMetadata(sourcePDFPath, ht)
	splitPlanPath := fmt.Sprintf("%s/data/%s/%s", testDir, hansardType, sessionName)

	// Create dir that does not exist
	rawDataFolderSetup(splitPlanPath)
	werr := ioutil.WriteFile(splitPlanPath+"/split.yml", b, 0644)
	if werr != nil {
		return werr
	}

	return nil
}

// Helper function
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
