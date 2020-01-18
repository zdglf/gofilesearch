package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zdglf/gofilesearch/routers/api"
)

const (
	searchEngineApiPath = "/search"
	searchDocPath       = "/doc"
	tempaltePath        = "../templates/*"
	assetsPath          = "/assets"
	assetsDir           = "../assets"

	adminPath = "/admin"
)

func InitRouters() (routers *gin.Engine) {
	routers = gin.New()
	routers.Use(gin.Logger())
	routers.Use(gin.Recovery())

	routers.LoadHTMLGlob(tempaltePath)
	routers.Static(assetsDir, assetsPath)

	searchEngineApi := routers.Group(searchEngineApiPath)
	{
		searchEngineApi.POST(searchDocPath, api.SearchDoc)

	}

	// adminApi := routers.Group(adminPath, gin.BasicAuth(gin.Accounts{}))
	// {
	// 	adminApi.POST()
	// }
	return

}
