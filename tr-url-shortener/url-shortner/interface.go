package urlshortner

type IUrlShortner interface {
	Create(UrlShortenerReqeust) (UrlShortenerResponse, error)
	Redirect(UrlShortenerRedirectRequest) (UrlShortenerRedirectResponse, error)
}
