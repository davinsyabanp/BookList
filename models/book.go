package models

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationDate string `json:"publication_date"`
	Publisher       string `json:"publisher"`
	NumberOfPages   int    `json:"number_of_pages"`
	CategoryID      int    `json:"category_id"`
}
