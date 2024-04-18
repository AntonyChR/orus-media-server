package models

import "gorm.io/gorm"

type FileInfo struct {
	gorm.Model
	Name    string `json:"Name"`
	Path    string `json:"Path"`
	IsDir   bool   `json:"IsDir"`
	TitleId uint   `json:"TitleId"`
	Episode uint   `json:"Episode"`
	Season  uint   `json:"Season"`
}
