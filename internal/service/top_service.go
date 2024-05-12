package service

import (
	"fmt"
	"time"

	"github.com/eduardospek/notabaiana-backend-golang/internal/domain/entity"
)

type TopRepository interface {
	Create(tops []entity.Top) error
	TopTruncateTable() error
	FindAll() ([]entity.Top, error)
}

type TopService struct {
	TopRepository TopRepository
	NewsService NewsService
}

func NewTopService(toprepo TopRepository, newsservice NewsService) *TopService {
	return &TopService{  TopRepository: toprepo, NewsService: newsservice }
}

func (t *TopService) TopCreate() {

	news, err := t.NewsService.FindAllViews()

	if err != nil {
		fmt.Println(err)
	}	

	var tops []entity.Top
	var newtop entity.Top
	var ntop entity.Top
	
	for _, top := range news {
		
		newtop = entity.Top{
			Title: top.Title,
			Link: top.Link,
			Image: top.Image,
			CreatedAt: top.CreatedAt,
			Views: top.Views,
		}

		ntop = *entity.NewTop(newtop)		

		tops = append(tops, ntop)
	}	

	err = t.TopRepository.TopTruncateTable()

	if err != nil {
		fmt.Println(err)
	}

	err = t.TopRepository.Create(tops)

	if err != nil {
		fmt.Println(err)
	}

	err = t.NewsService.ClearViews()

	if err != nil {
		fmt.Println(err)
	}
}

func (t *TopService) FindAll() ([]entity.Top, error) {
	tops, err := t.TopRepository.FindAll()

	if err != nil {
		return []entity.Top{}, err
	}

	return tops, nil
}

func (t *TopService) Start(minutes time.Duration) {
	
	go t.TopCreate()

	ticker := time.NewTicker(minutes * time.Minute)
    defer ticker.Stop()

    for range ticker.C {
		go t.TopCreate()
	}
}