export interface FileInfo {
    ID:        number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt: null;
    Name:      string;
    Path:      string;
    IsDir:     boolean;
    TitleId:   number;
    Episode:   number;
    Season:    number;
}