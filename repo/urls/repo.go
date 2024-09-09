package urls

import (
	"context"
	"urlshortner/clients/cache"
	"urlshortner/clients/db"
	"urlshortner/models"

	_ "github.com/lib/pq"
)

type IUrlsRepo interface {
	ListUrls(ctx context.Context) ([]models.Url, error)
	AddUrl(ctx context.Context, url models.Url) error
	GetUrlWithShortUrl(ctx context.Context, url string) (string, error)
}

type UrlsRepo struct {
	db    db.DB
	cache cache.ICache
}

func NewUrlsRepo(db db.DB, cache cache.ICache) *UrlsRepo {
	repo := &UrlsRepo{db: db, cache: cache}
	return repo
}

func (repo *UrlsRepo) GetCache() cache.ICache { return repo.cache }
