package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zdglf/gofilesearch/api_model"
	"github.com/zdglf/gofilesearch/es_model"
	"github.com/zdglf/gofilesearch/util/flog"
	"net/http"
)

type SearchFileService struct {
	Context *gin.Context
}

func (this *SearchFileService) ResponseApiFileSearch() {
	var searchReq = api_model.SearchRequest{}
	var err error
	searchResp := &api_model.SearchResultResponse{}
	statusCode := &api_model.StatusCode{}
	searchResp.StatusCode = statusCode
	if err = this.Context.ShouldBindJSON(&searchReq); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_PARAM_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, searchResp)
		return
	}
	if searchResp.Data, searchResp.Page, err = es_model.SearchDocument(searchReq.Keyword, searchReq.PageIndex); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_ES_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, searchResp)
		return
	}
	statusCode.Code = api_model.CODE_SUCCESS
	statusCode.Msg = api_model.CODE_MSG_SUCCESS
	this.Context.JSON(http.StatusOK, searchResp)
	return

}
