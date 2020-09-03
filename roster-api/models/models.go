package models

// DayRoster ...
type DayRoster struct {
	UpperStaff    string `json:"upperStaff,omitempty"`
	UpperTime     string `json:"upperTime,omitempty"`
	LowerStaff    string `json:"lowerStaff,omitempty"`
	LowerTime     string `json:"lowerTime,omitempty"`
	CustomMessage string `json:"customMessage,omitempty"`
}

// WeekRoster ...
type WeekRoster struct {
	ID   int     `json:"Id"`
	Year int     `json:"Year"`
	Week int     `json:"Week"`
	Mon  []uint8 `json:"Mon"`
	Tue  []uint8 `json:"Tue"`
	Wed  []uint8 `json:"Wed"`
	Thu  []uint8 `json:"Thu"`
	Fri  []uint8 `json:"Fri"`
	Sat  []uint8 `json:"Sat"`
	Sun  []uint8 `json:"Sun"`
}
