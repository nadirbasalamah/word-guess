package models

type WordBank struct {
	Data []Datum `json:"data"`
}

type Datum struct {
	Letters []string `json:"letters"`
	Answers []string `json:"answers"`
}
