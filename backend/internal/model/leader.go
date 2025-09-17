package model

type Leader struct {
	Name                   string
	Email                  string
	SchoolEmail            string
	Phone                  string
	Qualifications         []Qualification
	AssociatedAvailability map[string][]TimeBlock
	SessionsAvailability   map[string][]TimeBlock // day -> free time blocks
	Preferences            Preferences
}

type Qualification struct {
	Course string
	Grade  string
}

type TimeBlock struct {
	Start string
	End   string
}

type Preferences struct {
	WillingToDouble   bool
	AttendingTraining bool
	AttendingMeetings bool
	Notes             string
}
