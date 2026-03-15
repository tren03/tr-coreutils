package urlshortner

import "context"

type UrlShortenerReqeust struct {
	Url    string
	UserId *string
	ctx context.Context
}

type UrlShortenerResponse struct {
	Code     string
	ShortUrl string
}

type UrlShortenerRedirectRequest struct {
	Code string
	ctx context.Context
}

type UrlShortenerRedirectResponse struct {
	Url string
}
