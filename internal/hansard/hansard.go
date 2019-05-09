package hansard

import "github.com/davecgh/go-spew/spew"

type HansardPage struct {
	pageNo                    int
	plainTextContent          string
	isPossibleStartofQuestion bool
	possibleQuestionNum       string
	questionNumberSnippet     string
}

type HansardQuestion struct {
	questionNum  string
	pageNumStart int
	pageNumEnd   int
	pages        []HansardPage
}

type HansardDocument struct {
	HansardType      HansardType
	SessionName      string // Get this from the front page cover .. or the reference lookup ..
	HansardQuestions []HansardQuestion
	splitterState    splitterState
	originalPDFPath  string // Used for split later ..
}

type HansardType int

const (
	HANSARD_SPOKEN HansardType = iota
	HANSARD_WRITTEN
)

type QuestionStatus int

const (
	QUESTION_NOT_SEEN QuestionStatus = iota
	QUESTION_SEEN
	QUESTION_EXTRACTED
)

type splitterState struct {
	lastMarkedPage        int
	lastMarkedQuestionNum string
	questionsStatus       map[string]QuestionStatus
}

func NewHansardDocument(pdfPath string) (*HansardDocument, error) {
	hansardDoc := HansardDocument{
		HansardType:     HANSARD_WRITTEN,
		SessionName:     detectPossibleSessionName(pdfPath),
		originalPDFPath: pdfPath,
	}
	return &hansardDoc, nil
}

func detectPossibleSessionName(pdfPath string) string {

	return "defaultSession"
}

func detectPossibleQuestionNum(linesExcerpt []string) error {
	return nil
}

// ProcessLinesExcerpt takes the extracted excerpt; and pull out all the metadata
func (hd *HansardDocument) ProcessLinesExcerpt(linesExcerpt []string) error {
	return nil
}

// String to dump out the structure we have derived; ready to output to pdfcpu to split Command!
func (hd *HansardDocument) String() {
	spew.Dump(hd)
}

// SplitPDFByQuestions output to actual PDF based on derived data
func (hd *HansardDocument) SplitPDFByQuestions() error {
	// Guard checks; whgat if got nothing; check length ..
	return nil
}

func Split(t string, c string) []string {
	return []string{"bob"}
}
