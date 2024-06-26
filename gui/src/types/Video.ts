export interface Video {
    ID:        number;
    CreatedAt: Date;
    UpdatedAt: Date;
    Name:      string;
    Path:      string;
    IsDir:     boolean;
    TitleId:   number;
    Episode:   number;
    Season:    number;
    Ext:       string;
}