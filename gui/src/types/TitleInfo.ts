export interface TitleInfo {
    ID:           number;
    CreatedAt:    Date;
    UpdatedAt:    Date;
    DeletedAt:    null;
    Title:        string;
    Year:         string;
    Rated:        string;
    Released:     string;
    Runtime:      string;
    Genre:        string;
    Director:     string;
    Writer:       string;
    Actors:       string;
    Plot:         string;
    Language:     string;
    Country:      string;
    Awards:       string;
    Poster:       string;
    imdbRating:   string;
    imdbID:       string;
    Type:         TitleType;
    totalSeasons: string;
    Folder:       string;
}

type TitleType = |"series" | "movie"