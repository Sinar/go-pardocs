package debate

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
	MaxDPFSample = 3
)

func NewDebateTOC(sourcePath string) (*DebateTOC, error) {
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
