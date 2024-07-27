package model

type Report struct {
	ID        string `json:"id"`
	AuthorID  string `json:"author_id"`
	Count 	 int    `json:"count"`
	Title     string `json:"title"`
	Style     string `json:"style"`
	Language  string `json:"language"`
}
