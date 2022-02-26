package models

// capital case so it can be exported
// `json:field` so it can be read as json
type book struct{
	ID			string 	`json:"id"`
	Title		string 	`json:"title"`
	Author		string 	`json:"author"`
	Quantity	int		`json:"quantity"`
}