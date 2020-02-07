package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zdglf/gofilesearch/routers/api"
)

const (
	searchEngineApiPath = "/search"
	searchApiDocPath    = "/doc"

	tempaltePath = "../templates/*"
	assetsPath   = "/assets"
	assetsDir    = "assets"

	buildPath = "/web"
	buildDir  = "web/filesearch/dist"

	adminApiPath = "/admin"

	taskApiPath = "/task"

	createApiPath = "/create"
	execApiPath   = "/exec"
	listApiPath   = "/list"
	deleteApiPath = "/delete"
	modifyApiPath = "/modify"
	infoApiPath   = "/info"
)

func InitRouters() (routers *gin.Engine) {
	routers = gin.New()
	routers.Use(gin.Logger())
	routers.Use(gin.Recovery())

	//routers.LoadHTMLGlob(tempaltePath)
	routers.Static(assetsPath, assetsDir)
	routers.Static(buildPath, buildDir)

	searchEngineApi := routers.Group(searchEngineApiPath)
	{
		searchEngineApi.POST(searchApiDocPath, api.SearchDoc)

	}

	adminTaskApi := routers.Group(fmt.Sprintf("%s%s", adminApiPath, taskApiPath))
	{
		adminTaskApi.POST(createApiPath, api.AdminTaskCreate)
		adminTaskApi.POST(listApiPath, api.AdminTaskList)
		adminTaskApi.POST(execApiPath, api.AdminTaskExec)
		adminTaskApi.POST(deleteApiPath, api.AdminTaskDelete)
		adminTaskApi.POST(modifyApiPath, api.AdminTaskModify)
	}
	return

}
