package service

import (
	"sessionmanagement/api/model/request"
	"sessionmanagement/api/model/response"
	"sessionmanagement/api/repository"
)

type ArticleService interface {
	AddArticle(request.Article) error
	GetMyArticles(int64) ([]response.Article, error)
	GetArticleById(int64) (response.Article, error)
	UpdateArticle(request.Article) error
	DeleteArticle(int64) error
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(a repository.ArticleRepository) ArticleService {
	return articleService{
		articleRepository: a,
	}
}

func (a articleService) AddArticle(article request.Article) error {
	return a.articleRepository.AddArticle(article)
}

func (a articleService) GetMyArticles(author int64) ([]response.Article, error) {
	return a.articleRepository.GetMyArticles(author)
}

func (a articleService) GetArticleById(id int64) (response.Article, error) {
	return a.articleRepository.GetArticleById(id)
}

func (a articleService) UpdateArticle(article request.Article) error {
	return a.articleRepository.UpdateArticle(article)
}

func (a articleService) DeleteArticle(id int64) error {
	return a.articleRepository.DeleteArticle(id)
}