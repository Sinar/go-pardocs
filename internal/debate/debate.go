package debate

import (
	"fmt"
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

type DebateSession struct {
	SessionName   string
	SessionDate   string
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
