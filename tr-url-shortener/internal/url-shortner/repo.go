package urlshortner

import "context"

type IUrlRepo interface {
	SaveShortUrl(context.Context, string, string) error
}

type IUserRepo interface {
	AuthCheck(context.Context, string) bool
}

type UrlRepo struct{
}

func (r *UrlRepo) SaveShortUrl(_ context.Context, _ string, _ string) error {
	return nil
}
