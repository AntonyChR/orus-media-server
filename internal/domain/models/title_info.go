package models

import "gorm.io/gorm"

type TitleInfo struct {
	gorm.Model

	Title        string `json:"Title"`
	Year         string `json:"Year"`
	Rated        string `json:"Rated"`
	Released     string `json:"Released"`
	Runtime      string `json:"Runtime"`
	Genre        string `json:"Genre"`
	Director     string `json:"Director"`
	Plot         string `json:"Plot"`
	Poster       string `json:"Poster"`
	ImdbRating   string `json:"imdbRating"`
	ImdbID       string `json:"imdbID"`
	Type         string `json:"Type"`
	TotalSeasons string `json:"totalSeasons"`
	Folder       string `json:"Folder"`
}
