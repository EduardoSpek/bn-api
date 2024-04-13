package memorydb

import "github.com/eduardospek/bn-api/internal/domain/entity"

type NewsMemoryRepository struct {
	Newsdb map[string]entity.News
}

func NewNewsMemoryRepository() *NewsMemoryRepository {
	return &NewsMemoryRepository{ Newsdb: make(map[string]entity.News) }
}

func (r *NewsMemoryRepository) Create(news entity.News) (entity.News, error) {
	r.Newsdb[news.ID] = news
	return news, nil
}

func (r *NewsMemoryRepository) FindAll() ([]entity.News, error) {
	var news []entity.News
	for _, n := range r.Newsdb {
		news = append(news, n)
	}
	return news, nil
}