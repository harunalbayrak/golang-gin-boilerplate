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

type ProjectController struct{}

type CreateProjectInput struct {
	Name string `json:"name" binding:"required"`
}

func (p ProjectController) CreateProject(c *gin.Context) {
	var input CreateProjectInput
	var appGin = app.Gin{C: c}

	if err := c.ShouldBindJSON(&input); err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_CREATING_PROJECT_WRONG_INPUT, nil)
		return
	}

	project := models.Project{Name: input.Name}
	models.AddProject(&project)

	appGin.Response(http.StatusOK, e.SUCCESS, project)
}

func (p ProjectController) GetProjects(c *gin.Context) {
	var appGin = app.Gin{C: c}
	config := config.GetConfig()

	projects, err := models.GetProjects(0, config.GetInt("pageSize"), make(map[string]interface{}))
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_GETTING_PROJECTS, nil)
		return
	}

	appGin.Response(http.StatusOK, e.SUCCESS, projects)
}

func (p ProjectController) GetProject(c *gin.Context) {
	var appGin = app.Gin{C: c}

	projectID, _ := c.Params.Get("project_id")
	projectIDInt, err := strconv.Atoi(projectID)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_GETTING_PROJECT_NOT_INTEGER_PROJECT_ID, nil)
		return
	}

	project, err := models.GetProject(projectIDInt)
	if err != nil {
		appGin.Response(http.StatusInternalServerError, e.ERROR_GETTING_PROJECT, nil)
		return
	}

	appGin.Response(http.StatusOK, e.SUCCESS, project)
}
