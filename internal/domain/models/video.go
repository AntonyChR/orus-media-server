package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Name    string `json:"Name"`
	Path    string `json:"Path"`
	TitleId uint   `json:"TitleId"`
	Episode uint   `json:"Episode"`
	Season  uint   `json:"Season"`
}
