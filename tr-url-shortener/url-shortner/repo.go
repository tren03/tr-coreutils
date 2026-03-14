package urlshortner

type Repo interface {
	CreateShortUrl(string, string, string) string
}
