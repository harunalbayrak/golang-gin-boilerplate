package server

import (
	"github.com/gin-gonic/gin"
	"github.com/harunalbayrak/golang-gin-boilerplate/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	// router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CustomLogger)

	health := new(controllers.HealthController)
	project := new(controllers.ProjectController)
	article := new(controllers.ArticleController)

	router.GET("/health", health.Status)

	// Projects
	router.POST("/projects/create", project.CreateProject)
	router.GET("/projects", project.GetProjects)
	router.GET("/projects/:project_id", project.GetProject)

	// // Articles
	router.POST("/projects/:project_id/articles/create", article.CreateArticle)
	router.GET("/projects/:project_id/articles", article.GetArticles)
	router.GET("/projects/:project_id/articles/:article_id", article.GetArticle)

	return router
}
