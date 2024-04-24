package repositoryImplementations

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	models "github.com/AntonyChR/orus-media-server/internal/domain/models"
)

// NewOmdbProvider creates a new instance of OmdbApiTitleInfoProv with the provided base URL API and API key.
func NewOmdbProvider(baseUrlApi, apiKey string) *OmdbApiTitleInfoProv {
	return &OmdbApiTitleInfoProv{
		ApiUrl: baseUrlApi + "/?apikey=" + apiKey,
	}
}

// OmdbApiTitleInfoProv represents a provider for obtaining movie information using the OMDb API.
// The OMDb API is a RESTful web service to obtain movie information.
// More information can be found at: https://www.omdbapi.com/
type OmdbApiTitleInfoProv struct {
	ApiUrl string
}

// Search searches for movie information based on the provided file name.
// It returns a models.TitleInfo struct containing the movie information, or an error if the search fails.
func (m *OmdbApiTitleInfoProv) Search(fileName string) (models.TitleInfo, error) {
	var info models.TitleInfo

	params := extractSearchParams(fileName)

	if params[0] == "" {
		return info, errors.New("invalid filename format")
	}
	year := ""
	title := "&t=" + strings.ReplaceAll(params[0], " ", "%20")

	if params[1] != "" {
		year = "&y=" + params[1]
	}

	url := m.ApiUrl + title + year

	log.Println("Search: ", title, ", year: ", year)
	resp, err := http.Get(url)

	if err != nil {
		return info, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Data not found: ", url)
		return info, errors.New("not found")
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return info, err
	}

	err = json.Unmarshal(bodyBytes, &info)

	if err != nil {
		return info, err
	}

	return info, nil
}

// extractSearchParams extracts the search parameters from the given file name.
// It removes certain strings from the file name and extracts the title and year
// information from it. The search parameters are returned as a slice of strings,
// where the first element is the title and the second element is the year.
//
// Example:
//
//	fileName := "Movie (2021).1080p.bluray.mkv"
//	searchParams := extractSearchParams(fileName)
//	// searchParams will be ["Movie", "2021"]
func extractSearchParams(fileName string) []string {
	toRemove := []string{"1080p", "720p", "5.1", "2.0", "bluray", "webrip", "web-dl", "brrip", "dvdrip", "dvdscr", "hdrip", "hdtv", "wmv"}

	clearedName := fileName
	for _, r := range toRemove {
		clearedName = strings.ReplaceAll(clearedName, r, "")
	}

	nameWithoutExt := strings.SplitN(clearedName, ".", 2)[0]

	// Extract year and title from the file name
	var year, title string
	if strings.Contains(nameWithoutExt, "(") && strings.Contains(nameWithoutExt, ")") {
		splitName := strings.SplitN(nameWithoutExt, "(", 2)
		title = strings.TrimSpace(splitName[0])
		year = strings.SplitN(splitName[1], ")", 2)[0]
	} else {
		title = strings.TrimSpace(nameWithoutExt)
	}

	return []string{title, year}
}
