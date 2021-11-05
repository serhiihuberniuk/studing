package models

type ExcelFile struct {
	Name        string
	ID          string
	NeedConvert bool
	Sheets      []Sheet
}

type Sheet struct {
	Name string
	Rows []Row
}

type Row struct {
	Cells []Cell
}

type Cell struct {
	Value string
}
