package pardocs // import "github.com/Sinar/go-pardocs"
import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/Sinar/go-pardocs/internal/hansard"
)

type ParliamentDocs struct {
	Conf Configuration
}

// CommandMode specifies the operation being executed.
type CommandMode int

// The available commands.
const (
	PLAN CommandMode = iota
	SPLIT
	RESET
)

// Configuration of a Context.
type Configuration struct {

	// Parliament Session Label
	ParliamentSession string

	// Hansard Type
	HansardType hansard.HansardType

	// ./raw + ./data folders are assumed to be relative to this dir
	WorkingDir string

	// Source PDF can be anywhere; maybe make it a Reader to be read direct from S3?
	SourcePDFPath string

	// Command being executed.
	Cmd CommandMode
}

func getParliamentDocMetadata(pdfPath string, ht hansard.HansardType) (sessionName string, hansardType string) {
	baseFilename := filepath.Base(pdfPath)
	sessionName = strings.Split(baseFilename, ".")[0]
	switch ht {
	case hansard.HANSARD_SPOKEN:
		hansardType = "Lisan"
	case hansard.HANSARD_WRITTEN:
		hansardType = "BukanLisan"
	default:
		panic(fmt.Errorf("INVALUED TYPE!!!"))
	}

	return sessionName, hansardType
}

func (pd *ParliamentDocs) Plan() {
	log.Println("In Plan ..")
	pdfPath := pd.Conf.SourcePDFPath
	// Extract out hansard.MaxLineProcessed lines from each page to be analyzed
	pdfDoc, err := hansard.NewPDFDoc(pdfPath)
	if err != nil {
		log.Fatal(err)
	}
	// Sanity check before proceeding ..
	if len(pdfDoc.Pages) < 1 {
		log.Fatal("Could NOT find any pages!")
	}
	// Analyze the Hansard Document to find the question split
	hansardDoc, _ := hansard.NewHansardDocument(pdfPath)
	for _, p := range pdfDoc.Pages {
		//log.Println("PAGE:", p.PageNo)
		// Detect question
		dterr := hansardDoc.ProcessLinesExcerpt(p.PageNo, p.PDFTxtSameLines)
		if dterr != nil {
			log.Fatal(dterr)
		}
	}
	// Wrap up processing; what if there is no pages?
	hansardDoc.Finalize()
	// TODO: Better refactoring somewhere else? looks like a bit of a hack ..
	hansardDoc.ParliamentSession = pd.Conf.ParliamentSession // Mis-naming? is this the right place to place this?
	hansardDoc.HansardType = pd.Conf.HansardType
	// Persist the  plan
	sessionName, hansardType := getParliamentDocMetadata(pdfPath, pd.Conf.HansardType)
	hansardDoc.PersistForSplit(fmt.Sprintf("%s/data/%s/%s", pd.Conf.WorkingDir, hansardType, sessionName))
}

func (pd *ParliamentDocs) Split() {
	log.Println("In Split ..")
}

func (pd *ParliamentDocs) Reset() {
	log.Println("In Reset ...")
	// Clean up plan
	// Clean up split pages folder
	// Clean up merged pages location
}
