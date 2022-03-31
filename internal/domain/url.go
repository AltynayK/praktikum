package domain

type Url struct {
	LongURL  string `json:"LongUrl"`
	ShortURL string `json:"short_url"`
}

type PostResponse struct {
	ID int `json:"id"`
}

type GetResponse struct {
	Url string `json:"url"`
}

type ErrResponse struct {
	Msg string `json:"message"`
}
