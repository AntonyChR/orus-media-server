package repositoryImplementations

import (
	"os"
	"testing"
	"time"
)

func TestExtractSearchParams(t *testing.T) {

	cases := []struct {
		fileName string
		expected []string
	}{
		{
			fileName: "The Matrix (1999).mp4",
			expected: []string{"The Matrix", "1999"},
		},
		{
			fileName: "The Matrix.mp4",
			expected: []string{"The Matrix", ""},
		},
		{
			fileName: "The Matrix (1999) 1080p.mp4",
			expected: []string{"The Matrix", "1999"},
		},
		{
			fileName: "The Matrix 1080p.mp4",
			expected: []string{"The Matrix", ""},
		},
		{
			fileName: "The Matrix (1999) 1080p 5.1  .mp4",
			expected: []string{"The Matrix", "1999"},
		},
		{
			fileName: "The Matrix 1080p  5.1.mp4",
			expected: []string{"The Matrix", ""},
		},
		{
			fileName: "Movie (2021).1080p.bluray.mkv",
			expected: []string{"Movie", "2021"},
		},
	}

	for i, c := range cases {
		got := extractSearchParams(c.fileName)
		if got[0] != c.expected[0] {
			t.Errorf("%v) extractSearchParams(%q) == %q, want %q", i+1, c.fileName, got[0], c.expected[0])
		}
	}

}

func TestSearchTitleInfo(t *testing.T) {

	apiKey := os.Getenv("OMDB_API_KEY")

	if apiKey == "" {
		t.Skip("OMDB_API_KEY is not set")
	}

	omdbProvider := NewOmdbProvider("http://www.omdbapi.com", &apiKey)

	fileNames := []string{
		"The Matrix (1999).mp4",
		"The Matrix.mp4",
		"The Matrix (1999) 1080p.mp4",
		"The Matrix 1080p.mp4",
	}

	for _, fileName := range fileNames {
		info, err := omdbProvider.Search(fileName)
		if err != nil {
			t.Errorf("Search(%q) failed with error: %v", fileName, err)
		}
		if info.Title == "" {
			t.Errorf("Search(%q) failed, title not found", fileName)
		}
		time.Sleep(1 * time.Second)
	}

}
