package domain

import "github.com/AntonyChR/orus-media-server/internal/domain/models"

type TitleInfoProvider interface {
	Search(fileName string) (models.TitleInfo, error)
}
