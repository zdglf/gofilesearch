package main
import (
    "github.com/gin-gonic/gin"
)


const(
    SEARCH_ENGINE_API_PATH = "/search"
    TEMPLATE_PATH = "templates/*"
    ASSETS_PATH = "/assets"
    ASSETS_DIR = "./assets"

    ADMIN_PATH = "/admin"
)

func main() {
    r := gin.Default()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.LoadHTMLGlob(TEMPLATE_PATH)
    r.Static(ASSETS_PATH, ASSETS_DIR)
    search_engine_api := r.Group(SEARCH_ENGINE_API_PATH)
    {
        //search_engine_api.POST()
    }

    admin_api := r.Group(ADMIN_PATH,gin.BasicAuth(gin.Accounts{

    }))
    {
        //admin_api.POST()
    }
    r.Run()
}
