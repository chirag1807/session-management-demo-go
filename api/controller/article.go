package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"sessionmanagement/api/model/request"
	"sessionmanagement/api/model/response"
	"sessionmanagement/api/service"
	"sessionmanagement/constants"
	"sessionmanagement/error"
	"sessionmanagement/utils"
	"strconv"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
)

type ArticleController interface {
	AddArticle(w http.ResponseWriter, r *http.Request)
	GetMyArticles(w http.ResponseWriter, r *http.Request)
	GetArticleById(w http.ResponseWriter, r *http.Request)
	UpdateArticle(w http.ResponseWriter, r *http.Request)
	DeleteArticle(w http.ResponseWriter, r *http.Request)
}

type articleController struct {
	articleService service.ArticleService
	sessionManager *scs.SessionManager
}

func NewArticleController(a service.ArticleService, sessionManager *scs.SessionManager) ArticleController {
	return articleController{
		articleService: a,
		sessionManager: sessionManager,
	}
}

func (a articleController) AddArticle(w http.ResponseWriter, r *http.Request) {
	var article request.Article

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadBodyError)
		return
	}

	err = json.Unmarshal(body, &article)
	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	author := a.sessionManager.GetInt64(r.Context(), "userid")
	article.Author = author
	err = a.articleService.AddArticle(article)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ARTICLE_ADDED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
	return
}

func (a articleController) GetMyArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := a.articleService.GetMyArticles(a.sessionManager.GetInt64(r.Context(), "userid"))

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.ArticleResponse{
		Article: articles,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
	return
}

func (a articleController) GetArticleById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	article, err := a.articleService.GetArticleById(id)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	articles := make([]response.Article, 0)
	articles = append(articles, article)
	response := response.ArticleResponse{
		Article: articles,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
	return
}

func (a articleController) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var article request.Article
	err := json.NewDecoder(r.Body).Decode(&article)

	if err != nil {
		utils.ErrorGenerator(w, errorhandling.ReadDataError)
		return
	}

	err = a.articleService.UpdateArticle(article)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ARTICLE_UPDATED,
	}
	utils.ResponseGenerator(w, http.StatusOK, response)
	return
}

func (a articleController) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "ID"), 10, 64)
	err := a.articleService.DeleteArticle(id)

	if err != nil {
		utils.ErrorGenerator(w, err)
		return
	}

	response := response.SuccessResponse{
		Message: constants.ARTICLE_DELETED,
	}
	utils.ResponseGenerator(w, 200, response)
	return
}
