package debate

import (
	"fmt"

	"github.com/sanity-io/litter"
)

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

type ParliamentNum int

const (
	DEBATE_PAR01 ParliamentNum = iota
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

type ParliamentSession int

const (
	DEBATE_PENGGAL1 ParliamentSession = iota
	DEBATE_PENGGAL2
	DEBATE_PENGGAL3
	DEBATE_PENGGAL4
	DEBATE_PENGGAL5
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
	ParliamentNum     ParliamentNum
	ParliamentSession ParliamentSession
	SessionDay        SessionDay
}

type DebateSession struct {
	SessionName   string
	SessionDate   string
	SessionMeta   SessionMeta
	OrderedEvents []Events
	Attendees     []DebateAttendees
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
	topics                []DebateTopic
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

func NewDebateTOC(sourcePath string) (*DebateTOC, error) {
	pdfDoc := PDFDocument{
		sourcePath: sourcePath,
	}

	// Example of options ..
	options := &ExtractPDFOptions{NumPages: MaxPDFSample}
	//options := &ExtractPDFOptions{NumPages: 2}

	exerr := pdfDoc.extractPDF(options)
	if exerr != nil {
		return nil, fmt.Errorf("extractPDF FAIL: %w", exerr)
	}

	litter.Dump(pdfDoc.Pages)

	foundTOC := false
	// Look out for TOCs!
	if pdfDoc.NumPages > 1 {
		foundTOC = true
	}

	if !foundTOC {
		return nil, fmt.Errorf("NewDebateTOC FAIL: %w",
			ErrorNoTOCFound{err: fmt.Sprintf("PDF at %s has NO TOC!", sourcePath)})
	}
	debateTOC := DebateTOC{}

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

func hasSessionDateHeader(pageContent []string) bool {
	// Sample a low number of value from the  top ..
	// Look for a regexp DR <date> <pageNum>
	// If DebateTOC has a detected Session Date; use that to square things ..
	return false
}
