package urlshortner

import "fmt"

type IUrlShortner interface {
	Create(UrlShortenerReqeust) (*UrlShortenerResponse, error)
	Redirect(UrlShortenerRedirectRequest) (*UrlShortenerRedirectResponse, error)
}

type UrlShortner struct {
	repo IUrlRepo
}

func (u *UrlShortner) Create(req UrlShortenerReqeust) (*UrlShortenerResponse, error) {
	if req.UserId == nil {
		return nil, fmt.Errorf("user does not exist %s", req.UserId)
	}
	// perform some hash and get a code
	code := "xyz" // we can check for collision later

	if err := u.repo.SaveShortUrl(req.ctx, *req.UserId, code); err != nil {
		return nil, err
	}

	return &UrlShortenerResponse{
		Code:     code,
		ShortUrl: code,
	}, nil
}
