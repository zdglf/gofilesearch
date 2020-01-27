package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zdglf/gofilesearch/api_model"
	"github.com/zdglf/gofilesearch/base_struct"
	"github.com/zdglf/gofilesearch/db_model"
	"github.com/zdglf/gofilesearch/util/flog"
)

type AdminTaskService struct {
	Context *gin.Context
}

func (this *AdminTaskService) ResponseApiTaskCreate() {

}

func (this *AdminTaskService) ResponseApiTaskModify() {

}

func (this *AdminTaskService) ResponseApiTaskDelete() {
	var statusCode = &api_model.StatusCode{}
	var taskIdReq = &api_model.TaskIdRequest{}
	var err error

	if err = this.Context.ShouldBindJSON(taskIdReq); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_PARAM_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	var fileSpider *db_model.FileSpider
	if fileSpider, err = db_model.GetFileSpiderById(taskIdReq.Id); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_DB_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	if err = fileSpider.Delete(); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_DB_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	statusCode.Code = api_model.CODE_SUCCESS
	statusCode.Msg = api_model.CODE_MSG_SUCCESS
	this.Context.JSON(http.StatusOK, statusCode)
	return

}

func (this *AdminTaskService) ResponseApiTaskExec() {
	var statusCode = &api_model.StatusCode{}
	var taskIdReq = &api_model.TaskIdRequest{}
	var err error

	if err = this.Context.ShouldBindJSON(taskIdReq); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_PARAM_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	var fileSpider *db_model.FileSpider
	if fileSpider, err = db_model.GetFileSpiderById(taskIdReq.Id); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_DB_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	go this.runLoadFileAndUpdate(fileSpider)
	statusCode.Code = api_model.CODE_SUCCESS
	statusCode.Msg = api_model.CODE_MSG_SUCCESS
	this.Context.JSON(http.StatusOK, statusCode)
	return
}

func (this *AdminTaskService) runLoadFileAndUpdate(fs *db_model.FileSpider) {
	loadFileService := &LoadFileService{}
	loadFileService.LoadFile(fs)
	fs.LastRunningTime = base_struct.JsonTime(time.Now())
	if err := fs.Update(db_model.FieldFileSpiderLastRunningTime); err != nil {
		flog.Fatalln(err.Error())
	}
}

func (this *AdminTaskService) ResponseApiTaskList() {

}
