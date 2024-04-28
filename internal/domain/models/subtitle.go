package models

import "gorm.io/gorm"

type Subtitle struct {
	gorm.Model
	Path    string `json:"Path"`
	Name    string `json:"Name"`
	Lang    string `json:"Lang"`
	VideoId uint   `json:"VideoId"`
}
