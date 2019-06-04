package hansard

import (
	"fmt"

	papi "github.com/hhrutter/pdfcpu/pkg/api"
	"github.com/hhrutter/pdfcpu/pkg/pdfcpu"
	"github.com/y0ssar1an/q"
)

type SplitPlan struct {
	QuestionNum  string
	PageNumStart int
	PageNumEnd   int
}

// NewSplitPlan will use a Reader (better!) to extract out the plan
func NewSplitPlan(planFilename string) []SplitPlan {
	// TODO: Read the plan file
	return nil
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
	outputFilename := fmt.Sprintf("%s-soalan-%s-bukanlisan", label, sp.QuestionNum)
	fmt.Println("====== ", outputFilename, " =======")

	pageSelection := fmt.Sprintf("%d-%d", sp.PageNumStart, sp.PageNumEnd)

	conf := pdfcpu.NewDefaultConfiguration()
	//conf.Cmd = pdfcpu.CommandMode(pdfcpu.SPLIT)

	// Use UNIX EOL?
	//wctx := pdfcpu.NewWriteContext(pdfcpu.EolLF)
	//wctx.DirName = "/tmp/"
	//wctx.FileName = "bob.pdf"
	//wctx.SelectedPages = pdfcpu.IntSet{pageSelection}
	//cmd := papi.SplitCommand("filenamein", "/tmp", 5, conf)

	baseRawDir := "/Users/mleow/GOMOD/go-pardocs/raw/BukanLisan"
	inputFileName := "Pertanyaan Jawapan Bukan Lisan 22019_new.pdf"
	fullInputPath := baseRawDir + "/" + inputFileName

	// For each question + folder combo
	// Extract out the only file there?
	// and put path into merge string; put in correct order ..

	cmd := papi.ExtractPagesCommand(fullInputPath, fmt.Sprintf("/tmp/go-pardocs/%d", 1), []string{pageSelection}, conf)
	o, perr := papi.Process(cmd)
	if perr != nil {
		panic(perr)
	}
	q.Q(o)

}
