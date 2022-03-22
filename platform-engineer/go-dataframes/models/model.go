package models

// Frame of data in the CSV
type DataFrame struct {
	GroupID      int    `csv:"group_id"` // .csv column headers
	IndividualID int    `csv:"individual_id"`
	Attr1        string `csv:"attr_1"`
	Attr2        string `csv:"attr_2"`
}

// Analysis on a group
type AnalyzedGroup struct {
	GroupID int `csv:"group_id"`
}


