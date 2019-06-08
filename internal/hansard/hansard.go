package hansard

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/davecgh/go-spew/spew"
	"github.com/y0ssar1an/q"
	"gopkg.in/yaml.v2"
)

type HansardPage struct {
	pageNo                    int
	plainTextContent          string
	isPossibleStartofQuestion bool
	possibleQuestionNum       string
	linesExcerpt              []string
}

type HansardQuestion struct {
	QuestionNum           string
	questionNumberSnippet string
	PageNumStart          int
	PageNumEnd            int
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
	lastMarkedPage         int
	lastMarkedQuestionNum  string
	questionsStatus        map[string]QuestionStatus
	currentHansardQuestion *HansardQuestion
	currentHansardPages    []HansardPage
}

func NewHansardDocument(pdfPath string) (*HansardDocument, error) {
	hansardDoc := HansardDocument{
		HansardType:     HANSARD_WRITTEN,
		SessionName:     detectPossibleSessionName(pdfPath),
		originalPDFPath: pdfPath,
		splitterState:   splitterState{0, "", make(map[string]QuestionStatus, 0), nil, nil},
	}
	return &hansardDoc, nil
}

func detectPossibleSessionName(pdfPath string) string {

	return "defaultSession"
}

func detectPossibleQuestionNum(linesExcerpt []string) (possibleQuestionNum string, derr error) {
	// Setup regexp once
	re := regexp.MustCompile(`(?i)^.*soalan.*\s+(\d+).*$`)

	for _, line := range linesExcerpt {

		sm := re.FindStringSubmatch(line)
		// DEBUG:
		//fmt.Println("LINE: ", line)
		//spew.Dump(sm)

		// If have "SOALAN" in there; pull out the regexp for digit
		// can be SOALAN NO <digit>
		// or NO SOALAN <digit>
		// \w* SOALAN \w* <digit>
		if sm != nil {
			// Extract out the number; as string ..
			// DEBUG:
			//fmt.Println("FOUND NUM: ", sm[1])
			return sm[1], nil
		}
	}
	// Empty means did not find possibleQuestionNum
	return "", nil
}

func NewHansardQuestion(pageNumStart int, possibleQuestionNum string) (*HansardQuestion, error) {
	// Guard rail
	//PageNumStart, err := strconv.Atoi(possibleQuestionNum)
	//if err != nil {
	//	return nil, err
	//}
	if pageNumStart < 1 {
		return nil, fmt.Errorf("Page Num %d invalid!", pageNumStart)
	}
	// Start and end can be the same page for a one pager?
	hansardQuestion := HansardQuestion{
		possibleQuestionNum,
		possibleQuestionNum,
		pageNumStart,
		pageNumStart,
		nil,
	}
	return &hansardQuestion, nil
}

// ProcessLinesExcerpt takes the extracted excerpt; and pull out all the metadata
func (hd *HansardDocument) ProcessLinesExcerpt(pageNum int, linesExcerpt []string) error {
	// DEBUG: INput
	//spew.Dump(linesExcerpt)

	// If found a new question; create a new HansardQuestion
	// else attach the HansardPage to it and update metadata

	// ALWAYS Create the HansardPage struct for this page ..
	newPage := HansardPage{pageNo: pageNum, linesExcerpt: linesExcerpt}
	// If detect possible Question Num; check metadata state
	possibleQuestionNum, err := detectPossibleQuestionNum(linesExcerpt)
	if err != nil {
		return err
	}
	// DEBUG
	//fmt.Println("STATE: ", hd.splitterState)

	// if found a genuine QuestionNum NOT matching lastMarkedQ; then execute; too narrow; there are  other scenarios
	if possibleQuestionNum != "" && possibleQuestionNum != hd.splitterState.lastMarkedQuestionNum {
		// Avoids special case for first iteration ..
		if hd.splitterState.lastMarkedQuestionNum != "" {
			// DEBUG:
			//fmt.Println("PAGE", pageNum, " POSSIBLE Q: ", possibleQuestionNum)
			// NOTE: SCOPED; for temp use
			hansardQuestion := hd.splitterState.currentHansardQuestion
			// Finalize from previous run; careful about off by one
			hansardQuestion.pages = hd.splitterState.currentHansardPages
			// DEBUG:
			//spew.Dump(hansardQuestion.pages)
			// Append the question AFTER appending the page!
			hd.HansardQuestions = append(hd.HansardQuestions, *hansardQuestion)
		}

		// Create new Question struct and attach it for the current page
		hansardQuestion, nhqerr := NewHansardQuestion(pageNum, possibleQuestionNum)
		if nhqerr != nil {
			return fmt.Errorf("NewHansardQuestion: FAILED: %v", nhqerr)
		}

		// Reset state
		newPage.isPossibleStartofQuestion = true
		newPage.possibleQuestionNum = possibleQuestionNum
		newPage.plainTextContent = "" // WHY is this here? Remove?
		// New one gets overwritten; will it get lost? or copied?
		// reset to new state
		hd.splitterState.lastMarkedQuestionNum = possibleQuestionNum
		hd.splitterState.currentHansardQuestion = hansardQuestion
		hd.splitterState.currentHansardPages = []HansardPage{newPage}

	} else {
		//  If is the start; we can attach special rules for extracting out
		// info like what session it is, ToC, and other stuff; for now; just ignore it
		if hd.splitterState.lastMarkedPage == 0 {
			//spew.Dump(hd.splitterState)
			fmt.Println("Probably front page; do semthign with this ..")
			// DEBUG:
			//spew.Dump(linesExcerpt)
			return nil
		}

		// just append the pages; with know metadata info ..
		newPage.isPossibleStartofQuestion = false
		newPage.possibleQuestionNum = hd.splitterState.lastMarkedQuestionNum
		newPage.plainTextContent = "" // WHY is this here? Remove?
		// Track current page; attach to existing one already ..
		hd.splitterState.currentHansardPages = append(hd.splitterState.currentHansardPages, newPage)
		// Up the page number to the current
		hd.splitterState.currentHansardQuestion.PageNumEnd = pageNum
	}

	// Mark tis invariant?
	hd.splitterState.lastMarkedPage = pageNum

	return nil
}

// String to dump out the structure we have derived; ready to output to pdfcpu to split Command!
func (hd *HansardDocument) String() {
	spew.Dump(hd)
}

// Finalize clean up all state and put it back into the structure
// TODO: Refactor this to be one clear structure
func (hd *HansardDocument) Finalize() {
	// Finalize from previous run; this will clean up any remaining
	hansardQuestion := hd.splitterState.currentHansardQuestion
	hansardQuestion.pages = hd.splitterState.currentHansardPages
	hd.HansardQuestions = append(hd.HansardQuestions, *hansardQuestion)

	// Clear out splitter state
	hd.splitterState = splitterState{}
}

// Debug function to dump out final state; it should be cleared after all the run ..
func (hd *HansardDocument) ShowState() {
	spew.Dump(hd.splitterState)
}

func (hd *HansardDocument) ShowQuestions() {
	spew.Dump(hd.HansardQuestions)
}

// SplitPDFByQuestions output to actual PDF based on derived data
func (hd *HansardDocument) SplitPDFByQuestions() error {
	// Guard checks; whgat if got nothing; check length ..
	for _, singleQuestion := range hd.HansardQuestions {
		q.Q("QUESTION: ", singleQuestion.QuestionNum, " START: ", singleQuestion.PageNumStart, " END: ", singleQuestion.PageNumEnd)
	}
	return nil
}

func Split(t string, c string) []string {
	return []string{"bob"}
}

func (hd *HansardDocument) PersistForSplit(absoluteRawDataPath string) error {

	// TODO: Restructure so that it is driven by the type? Or do we have full control?
	rawDataFolderSetup(absoluteRawDataPath)
	//spew.Dump(hd)
	b, err := yaml.Marshal(hd.HansardQuestions)
	if err != nil {
		panic(err)
	}

	werr := ioutil.WriteFile(absoluteRawDataPath+"/split.yml", b, 0744)
	if werr != nil {
		panic(werr)
	}

	return nil
}

// Need to restructure this utility func; I am sure there is a better implementation ;P
func rawDataFolderSetup(absoluteRawDataPath string) (proceedScraping bool) {
	fi, lerr := os.Stat(absoluteRawDataPath)
	if lerr != nil {
		if os.IsNotExist(lerr) {
			// Create the needed folder as per needed .. all along the chain
			mkerr := os.MkdirAll(absoluteRawDataPath, 0700)
			if mkerr != nil {
				panic(mkerr)
			}
			return true
		}
		panic(lerr)
	} else {
		// spew.Dump(fi)
		if !fi.IsDir() {
			panic(fmt.Errorf("NOT DIR!! %s", absoluteRawDataPath))
		}
		fmt.Println("Directory ", absoluteRawDataPath, " EXISTS!!")
		// When the folder already exist for the day, no need to proceed
		fmt.Println("Skipping... ")
	}

	return false
}
