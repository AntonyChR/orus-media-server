package models

import "gorm.io/gorm"

type Subtitle struct {
	gorm.Model
	Path      string `json:"Path"`
	Lang      string `json:"Lang"`
	IsDefault bool   `json:"IsDefault"`
	VideoId   uint   `json:"VideoId"`
}
