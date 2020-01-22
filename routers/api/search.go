package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zdglf/gofilesearch/service"
)

func SearchDoc(c *gin.Context) {
	searchFileService := &service.SearchFileService{Context: c}
	searchFileService.ResponseApiFileSearch()
}
