package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harunalbayrak/golang-gin-boilerplate/config"
	"github.com/harunalbayrak/golang-gin-boilerplate/models"
	"github.com/harunalbayrak/golang-gin-boilerplate/pkg/app"
	"github.com/harunalbayrak/golang-gin-boilerplate/pkg/e"
)

type ArticleController struct{}

type CreateArticleInput struct {
	Name      string `json:"name" binding:"required"`
	Protocol  string `json:"protocol" binding:"required"`
	ProjectID int    `json:"project_id" binding:"required"`
}

func (n ArticleController) CreateArticle(c *gin.Context) {
	var input CreateArticleInput
	var err error
	var appGin = app.Gin{C: c}

	projectID, _ := c.Params.Get("project_id")
	input.ProjectID, err = strconv.Atoi(projectID)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_CREATING_ARTICLE_NOT_INTEGER_PROJECT_ID, nil)
		return
	}
	existsProject, _ := models.ExistProjectByID(input.ProjectID)
	if !existsProject {
		appGin.Response(http.StatusInternalServerError, e.ERROR_CREATING_ARTICLE_NOT_EXISTS_PROJECT, nil)
		return
	}

	if err = c.ShouldBindJSON(&input); err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_CREATING_ARTICLE_WRONG_INPUT, nil)
		return
	}

	article := models.Article{Name: input.Name, Protocol: input.Protocol, ProjectID: input.ProjectID}
	models.AddArticle(&article)

	appGin.Response(http.StatusOK, e.ERROR_CREATING_ARTICLE_WRONG_INPUT, nil)
}

func (n ArticleController) GetArticles(c *gin.Context) {
	var appGin = app.Gin{C: c}
	config := config.GetConfig()

	projectId, _ := c.Params.Get("project_id")
	_, err := strconv.Atoi(projectId)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_GETTING_ARTICLES_NOT_INTEGER_PROJECT_ID, nil)
		return
	}

	articles, err := models.GetArticles(0, config.GetInt("pageSize"), projectId)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_GETTING_ARTICLES, nil)
		return
	}

	appGin.Response(http.StatusOK, e.SUCCESS, articles)
}

func (n ArticleController) GetArticle(c *gin.Context) {
	var appGin = app.Gin{C: c}

	projectID, _ := c.Params.Get("project_id")
	projectIDInt, err := strconv.Atoi(projectID)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_GETTING_ARTICLE_NOT_INTEGER_PROJECT_ID, nil)
		return
	}

	articleID, _ := c.Params.Get("article_id")
	articleIDInt, err := strconv.Atoi(articleID)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_GETTING_ARTICLE_NOT_INTEGER_ARTICLE_ID, nil)
		return
	}

	article, err := models.GetArticle(projectIDInt, articleIDInt)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_GETTING_ARTICLE, nil)
		return
	}

	appGin.Response(http.StatusOK, e.SUCCESS, article)
}
