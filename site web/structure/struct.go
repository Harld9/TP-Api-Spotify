package structure

type AlbumData struct {
	Name        string
	TotalTracks int
	ReleaseDate string
	ImageUrl    string
}

type ApiData struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

