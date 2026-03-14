package urlshortner

type UrlShortenerReqeust struct {
	Url    string
	UserId *string
}

type UrlShortenerResponse struct {
	Code     string
	ShortUrl string
}

type UrlShortenerRedirectRequest struct {
	Code string
}

type UrlShortenerRedirectResponse struct {
	Url string
}
