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

func NewOmdbProvider(baseUrlApi, apiKey string) *OmdbApiTitleInfoProv {
	return &OmdbApiTitleInfoProv{
		ApiUrl: baseUrlApi + "/?apikey=" + apiKey,
	}
}

// The OMDb API is a RESTful web service to obtain movie information.
// https://www.omdbapi.com/
type OmdbApiTitleInfoProv struct {
	ApiUrl string
}

func (m *OmdbApiTitleInfoProv) Search(fileName string) (models.TitleInfo, error) {
	var info models.TitleInfo

	params := extractSearchParams(fileName)

	if len(params) == 0 {
		return info, errors.New("invalid filename format")
	}
	year := ""
	title := "&t=" + strings.ReplaceAll(params[0], " ", "%20")

	if len(params) >= 2 {
		year = "&y=" + params[1]
	}

	url := m.ApiUrl + title + year

	log.Println("Get: ", url)
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

// extract valid parameters (title, year) from filename.
// E.g. from file name: "godzilla-2014.mp4" returns ["godzilla","2014"]
func extractSearchParams(fileName string) []string {
	nameWithoutExt := strings.Split(fileName, ".")[0]

	return strings.Split(nameWithoutExt, "-")
}
