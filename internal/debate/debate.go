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
	SessionName string
	SessionDate string

	OrderedEvents []Events
	Attendees     []DebateAttendees
}
