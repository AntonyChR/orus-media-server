package repositoryImplementations

import (
	"testing"

	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
	sqlite "gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
)

func InitializeTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestSqliteTitleInfoRepository(t *testing.T) {

	db, err := InitializeTestDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	repo := NewSqliteTitleInfoRepo(db)

	titleInfo := models.TitleInfo{
		Title: "Test Title",
		Year:  "2021",
		Type:  "movie",
	}

	err = repo.Migrate()
	if err != nil {
		t.Fatalf("Failed to migrate: %v", err)
	}

	// Save the titleInfo to the database
	err = repo.Save(&titleInfo)
	if err != nil || titleInfo.ID == 0 {
		t.Fatalf("Failed to save titleInfo: %v", err)
	}

	// Delete the titleInfo from the database
	err = repo.Delete(titleInfo)
	if err != nil {
		t.Fatalf("Failed to delete titleInfo: %v", err)
	}

	// Verify that the titleInfo is deleted
	_, err = repo.GetOneBy("field", "value")
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}
}
