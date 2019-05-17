package hansard

import "github.com/davecgh/go-spew/spew"

type HansardPage struct {
	pageNo                    int
	plainTextContent          string
	isPossibleStartofQuestion bool
	possibleQuestionNum       string
	linesExcerpt              []string
}

type HansardQuestion struct {
	questionNum           string
	questionNumberSnippet string
	pageNumStart          int
	pageNumEnd            int
	pages                 []HansardPage
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
	currentHansardPages   []HansardPage
}

func NewHansardDocument(pdfPath string) (*HansardDocument, error) {
	hansardDoc := HansardDocument{
		HansardType:     HANSARD_WRITTEN,
		SessionName:     detectPossibleSessionName(pdfPath),
		originalPDFPath: pdfPath,
		splitterState:   splitterState{0, "", make(map[string]QuestionStatus, 0), nil},
	}
	return &hansardDoc, nil
}

func detectPossibleSessionName(pdfPath string) string {

	return "defaultSession"
}

func detectPossibleQuestionNum(linesExcerpt []string) (possibleQuestionNum string, derr error) {
	for _, line := range linesExcerpt {
		// If have "SOALAN" in there; pull out the regexp for digit
		// can be SOALAN NO <digit>
		// or NO SOALAN <digit>
		// \w* SOALAN \w* <digit>
		if line == "SOALAN" {
			// Extract out the number; as string ..
			return line, nil
		}
	}
	// Empty means did not find possibleQuestionNum
	return "", nil
}

func NewHansardQuestion(possibleQuestionNum string) (*HansardQuestion, error) {
	hansardQuestion := HansardQuestion{
		possibleQuestionNum,
		possibleQuestionNum,
		0,
		0,
		make([]HansardPage, 0),
	}
	return &hansardQuestion, nil
}

// ProcessLinesExcerpt takes the extracted excerpt; and pull out all the metadata
func (hd *HansardDocument) ProcessLinesExcerpt(pageNum int, linesExcerpt []string) error {

	// ALWAYS Create the HansardPage struct for this page ..
	newPage := HansardPage{pageNo: pageNum, linesExcerpt: linesExcerpt}
	// If detect possible Question Num; check metadata state
	possibleQuestionNum, err := detectPossibleQuestionNum(linesExcerpt)
	if err != nil {
		return err
	}
	// if empty; just initialize it with the question mapped
	if possibleQuestionNum != hd.splitterState.lastMarkedQuestionNum {
		hansardQuestion, nhqerr := NewHansardQuestion(possibleQuestionNum)
		if nhqerr != nil {
			return nhqerr
		}

		// Case first one ..
		if hd.splitterState.currentHansardPages != nil {
			// Finalize from previous run; careful about off by one
			hansardQuestion.pages = hd.splitterState.currentHansardPages
			// Append the question AFTER appending the page!
			hd.HansardQuestions = append(hd.HansardQuestions, *hansardQuestion)
		}

		// Reset state
		newPage.isPossibleStartofQuestion = true
		newPage.possibleQuestionNum = possibleQuestionNum
		newPage.plainTextContent = "" // WHY is this here? Remove?
		// New one gets overwritten; will it get lost? or copied?
		hd.splitterState.currentHansardPages = []HansardPage{newPage}

	} else {
		// just append the pages; with know metadata info ..
		newPage.isPossibleStartofQuestion = false
		newPage.possibleQuestionNum = hd.splitterState.lastMarkedQuestionNum
		newPage.plainTextContent = "" // WHY is this here? Remove?
		// Track current page; attach to existing one already ..
		hd.splitterState.currentHansardPages = append(hd.splitterState.currentHansardPages, newPage)
	}

	// If found a new question; create a new HansardQuestion
	// else attach the HansardPage to it and update metadata

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
