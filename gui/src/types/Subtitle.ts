export interface Subtitle {
    ID:        number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt: null;
    Path:      string;
    Name:      string;
    VideoId:   number;
    Lang:      string;
}
