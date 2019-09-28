package debate

import (
	"fmt"
	"regexp"
	"strings"
)

type DebateSession struct {
	SessionName   string
	SessionDate   string
	SessionMeta   SessionMeta
	DebateTopics  []DebateTopic
	OrderedEvents []Events
	Attendees     []DebateAttendees
}

type DebateAttendees struct {
	FullName     string
	ShortName    string
	Constituency string
	Party        string
}

type Organization struct {
	OrgName    string
	OrgMembers []DebateAttendees
}

type Events struct {
	EventTime     string
	EventTitle    string
	EventSubTitle string
	PageStart     int
	PageEnd       int
	RelatedOrgs   []Organization
}

type ParliamentSession int

const (
	DEBATE_PAR01 ParliamentSession = iota
	DEBATE_PAR02
	DEBATE_PAR03
	DEBATE_PAR04
	DEBATE_PAR05
	DEBATE_PAR06
	DEBATE_PAR07
	DEBATE_PAR08
	DEBATE_PAR09
	DEBATE_PAR10
	DEBATE_PAR11
	DEBATE_PAR12
	DEBATE_PAR13
	DEBATE_PAR14
	DEBATE_PAR15
	DEBATE_PAR16
)

type ParliamentTerm int

const (
	DEBATE_PENGGAL1 ParliamentTerm = iota
	DEBATE_PENGGAL2
	DEBATE_PENGGAL3
	DEBATE_PENGGAL4
	DEBATE_PENGGAL5
)

type MeetingNum int

const (
	DEBATE_MESYUARAT1 MeetingNum = iota
	DEBATE_MESYUARAT2
	DEBATE_MESYUARAT3
	DEBATE_MESYUARAT4
	DEBATE_MESYUARAT5
)

type SessionDay int

const (
	DEBATE_MONDAY SessionDay = iota
	DEBATE_TUESDAY
	DEBATE_WEDNESDAY
	DEBATE_THURSDAY
	DEBATE_FRIDAY
	DEBATE_SATURDAY
	DEBATE_SUNDAY
)

type SessionMeta struct {
	SessionNum        int
	ParliamentSession ParliamentSession
	ParliamentTerm    ParliamentTerm
	MeetingNum        MeetingNum
	SessionDay        SessionDay
}

type DebateTopic struct {
	TopicName    string
	PageNumStart int
	PageNumEnd   int
}

type DebateTOC struct {
	//pdfPages              *[]PDFPage
	detectedSessionName   string
	detectedSessionDate   string
	normalizedSessionDate string
	detectedTopics        []string
}

const (
	MaxPDFSample = 3
)

type ErrorNoTOCFound struct {
	err string // Error description ..
}

func (e ErrorNoTOCFound) Error() string {
	return e.err
}

func NewDebateSession() {
	// Combine SessionMeta + DebateTopics to form the initial DebateSession??

}

func NewSessionMeta() {
	// from detected TOC; get the Session Metadata..

}

func NewDebateTopic() {
	// from detected TOC; get the Topics ..
	// Take a first cut naive for below ..
	//TopicName    string
	//PageNumStart int
	//PageNumEnd   int

}

func NewPDFDocForTOC(sourcePath string) (*PDFDocument, error) {
	pdfDoc := PDFDocument{
		SourcePath: sourcePath,
	}

	// Example of options ..
	options := &ExtractPDFOptions{NumPages: MaxPDFSample}
	//options := &ExtractPDFOptions{NumPages: 2}

	exerr := pdfDoc.extractPDF(options)
	if exerr != nil {
		return nil, fmt.Errorf("extractPDF FAIL: %w", exerr)
	}
	return &pdfDoc, nil
}

func NewDebateTOC(sourcePath string) (*DebateTOC, error) {
	pdfDoc, err := NewPDFDocForTOC(sourcePath)
	if err != nil {
		return nil, err
	}
	return NewDebateTOCPDFContent(pdfDoc)
}

func NewDebateTOCPDFContent(pdfDoc *PDFDocument) (*DebateTOC, error) {
	foundTOC := false
	// Look out for TOCs!
	//if pdfDoc.NumPages > 1 {
	//	foundTOC = true
	//}

	// Init to zero value ..
	debateTOC := DebateTOC{}
	// Regexp to detect date via 2019 year and against DR pattern?
	// Regexp to detect Parlimen Ke* pattern
	// Regexp to detect DR pattern in header e,g DR.1.7.2019 1 ; also tells  you halaman page?
	// Regexp to detect halaman for use later ..
	var detectedTopics []string
	for _, p := range pdfDoc.Pages {
		for _, r := range p.PDFTxtSameLines {
			// If detect DR header escape immediately!
			if hasSessionDateHeader(r) {
				break
			}

			pmatch, _ := regexp.MatchString(`PARLIMEN KE*`, r)
			if pmatch {
				// Debug
				//fmt.Println("PAGE: ", p.PageNo, "ROW: ", i, " PAR: ", r)
				debateTOC.detectedSessionName = strings.TrimSpace(r)
				//spew.Dump(debateTOC)
				//litter.Sdump(debateTOC)
			}
			dmatch, _ := regexp.MatchString(`\d+\s+\w+\s+2019`, r)
			if dmatch {
				// DEBUG
				//fmt.Println("PAGE: ", p.PageNo, "ROW: ", i, " DATE: ", r)
				debateTOC.detectedSessionDate = r
			}
			hmatch, _ := regexp.MatchString(`Halaman`, r)
			if hmatch {
				// DEBUG
				//fmt.Println("PAGE: ", p.PageNo, "ROW: ", i, " HAL: ", r)
				detectedTopics = append(detectedTopics, r)
			}
		}
	}
	// If minimum detect session date; consider found TOC?
	debateTOC.detectedTopics = detectedTopics
	//spew.Dump(debateTOC)
	// Private fields hidden by default!!
	// DEBUG
	//sq := litter.Options{
	//	HidePrivateFields: false,
	//}
	//sq.Dump(debateTOC)

	if len(debateTOC.detectedTopics) > 0 {
		foundTOC = true
	}
	if !foundTOC {
		return nil, fmt.Errorf("NewDebateTOC FAIL: %w",
			ErrorNoTOCFound{err: fmt.Sprintf("PDF at %s has NO TOC!", pdfDoc.SourcePath)})
	}

	// Extract lines only; sampling with the size
	// 	as per in MaxPDFSample; look out for the Header page ..
	return &debateTOC, nil
}

func extractTOC(samplePages []PDFPage, debateTOC *DebateTOC) error {
	// Take the sampling of pages
	return nil
}

func detectTOCTopicPage() (topic string, pageNum string, err error) {
	return "", "", nil
}

func detectSessionDate() (string, error) {
	return "", nil
}

func normalizeSessionDate() error {
	return nil
}

func hasSessionDateHeader(rowContent string) bool {
	// Sample a low number of value from the  top ..
	// Look for a regexp DR <date> <pageNum>
	// If DebateTOC has a detected Session Date; use that to square things ..
	return false
}
