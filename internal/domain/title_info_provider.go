package domain

import "github.com/AntonyChR/orus-media-server/internal/domain/models"

type TitleInfoProvider interface {
	// Search for a title info in the provider
	Search(fileName string) (models.TitleInfo, error)
}
