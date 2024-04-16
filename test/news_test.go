package test

import (
	"testing"

	"github.com/eduardospek/bn-api/internal/domain/entity"
	database "github.com/eduardospek/bn-api/internal/infra/database/memorydb"
	"github.com/eduardospek/bn-api/internal/service"
	"github.com/eduardospek/bn-api/internal/utils"
)

type TestCase struct {
	Esperado string
	Recebido string
}

func Resultado(t *testing.T, esperado string, recebido string) {
    t.Helper()
    if esperado != recebido {
        t.Errorf("Esperado: %s | Recebido: %s", esperado, recebido)
    }
}

func TestNewsEntity(t *testing.T) {
	t.Parallel()
	
	news := entity.News{		
		Title: "Titulo",
		Text: "Texto",
		Link: "http://www.eduardospek.com.br",
		Image: "https://www.bahianoticias.com.br/fotos/holofote_noticias/73825/IMAGEM_NOTICIA_original.jpg",
	}

	n := entity.NewNews(news)

	testcases := []TestCase{
		{
			Esperado: news.Title,
			Recebido: n.Title,
		},
		{
			Esperado: news.Text,
			Recebido: n.Text,
		},
		{
			Esperado: news.Link,
			Recebido: n.Link,
		},
		{
			Esperado: news.Image,
			Recebido: n.Image,
		},
	}

	for _, teste := range testcases {
		Resultado(t, teste.Esperado, teste.Recebido)
	}

}

func TestNewsService(t *testing.T) {
	t.Parallel()

	news_repo := database.NewNewsMemoryRepository()
	imagedownloader := utils.NewImgDownloader()
	news_service := service.NewNewsService(news_repo, imagedownloader)

	t.Run("Deve criar uma nova noticia no banco", func (t *testing.T)  {
		news := entity.News{		
			Title: "Titulo",
			Text: "Texto",
			Link: "http://www.eduardospek.com.br",
			Image: "https://www.bahianoticias.com.br/fotos/holofote_noticias/73825/IMAGEM_NOTICIA_original.jpg",
		}
	
		_, err := news_service.CreateNews(news)
	
		if err != nil {
			t.Error(err)
		}		
	})

	t.Run("Deve listar as noticias do banco", func (t *testing.T)  {

		lista := news_service.FindAllNews(1, 4)
		
		newsList := lista.(struct {
			Lista []entity.News `json:"news"`
		});
	
		if string(newsList.Lista[0].Title) == "" {
			t.Error("Nenhuma notícia no banco")
		}		
	})
	

}