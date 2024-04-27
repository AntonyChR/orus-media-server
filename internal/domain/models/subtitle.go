package models

import "gorm.io/gorm"

type Subtitle struct {
	gorm.Model
	Path      string `json:"Path"`
	Name      string `json:"Name"`
	IsDefault bool   `json:"IsDefault"`
	VideoId   uint   `json:"VideoId"`
}
