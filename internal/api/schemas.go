package api

type CreateShortURLSchema struct {
	URL string `json:"url"`
}

type CreateShortURLResponseSchema struct {
	ShortCode string `json:"short_code"`
	LongURL   string `json:"long_url"`
}
