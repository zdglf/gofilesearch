package service

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zdglf/gofilesearch/api_model"
	"github.com/zdglf/gofilesearch/base_struct"
	"github.com/zdglf/gofilesearch/db_model"
	"github.com/zdglf/gofilesearch/util/flog"
	"github.com/zdglf/gofilesearch/util/gen_data"
)

type AdminTaskService struct {
	Context *gin.Context
}

func (this *AdminTaskService) ResponseApiTaskCreate() {
	var statusCode = &api_model.StatusCode{}
	var taskCreateReq = &api_model.TaskCreateRequest{}
	var err error

	if err = this.Context.ShouldBindJSON(taskCreateReq); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_PARAM_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	var fileSpider = &db_model.FileSpider{}
	if fileSpider.Id, err = gen_data.GenUUID(); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_GEN_UUID_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	fileSpider.Type = taskCreateReq.Type
	fileSpider.Folder = taskCreateReq.Folder
	fileSpider.Username = taskCreateReq.Username
	fileSpider.Password = taskCreateReq.Password
	fileSpider.SizeLimit = taskCreateReq.SizeLimit
	fileSpider.ProcessSize = taskCreateReq.ProcessSize
	fileSpider.Enable = taskCreateReq.Enable
	fileSpider.Regular = taskCreateReq.Regular
	fileSpider.Timing = taskCreateReq.Timing

	if _, err = regexp.Compile(taskCreateReq.Regular); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_REG_EXP_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}

	if err = fileSpider.Insert(); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_DB_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	if fileSpider, err = db_model.GetFileSpiderById(fileSpider.Id); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_DB_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	statusCode.Code = api_model.CODE_SUCCESS
	statusCode.Msg = api_model.CODE_MSG_SUCCESS
	response := &api_model.DataResponse{}
	response.StatusCode = statusCode
	response.Data = fileSpider
	this.Context.JSON(http.StatusOK, response)
	return

}

func (this *AdminTaskService) ResponseApiTaskModify() {
	var statusCode = &api_model.StatusCode{}
	var request = &api_model.TaskModifyRequest{}
	var err error

	if err = this.Context.ShouldBindJSON(request); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_PARAM_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	var fileSpider *db_model.FileSpider

	if fileSpider, err = db_model.GetFileSpiderById(request.Id); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_DB_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}

	fileSpider.Type = request.Type
	fileSpider.Folder = request.Folder
	fileSpider.Username = request.Username
	fileSpider.Password = request.Password
	fileSpider.SizeLimit = request.SizeLimit
	fileSpider.ProcessSize = request.ProcessSize
	fileSpider.Enable = request.Enable
	fileSpider.Regular = request.Regular
	fileSpider.Timing = request.Timing

	if _, err = regexp.Compile(request.Regular); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_REG_EXP_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}

	if err = fileSpider.Update(); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_DB_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}

	statusCode.Code = api_model.CODE_SUCCESS
	statusCode.Msg = api_model.CODE_MSG_SUCCESS
	response := &api_model.DataResponse{}
	response.StatusCode = statusCode
	response.Data = fileSpider
	this.Context.JSON(http.StatusOK, response)
	return
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
	var statusCode = &api_model.StatusCode{}
	var request = &api_model.TaskListRequest{}
	var err error

	if err = this.Context.ShouldBindJSON(request); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_PARAM_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	var fileSpiderList []*db_model.FileSpider
	var pageInfo *api_model.Page
	if fileSpiderList, pageInfo, err = db_model.QueryFileSpiderList(request.PageIndex); err != nil {
		flog.Println(err.Error())
		statusCode.Code = api_model.CODE_DB_ERROR
		statusCode.Msg = err.Error()
		this.Context.JSON(http.StatusOK, statusCode)
		return
	}
	statusCode.Code = api_model.CODE_SUCCESS
	statusCode.Msg = api_model.CODE_MSG_SUCCESS
	response := &api_model.WithPageDataResponse{}
	response.StatusCode = statusCode
	response.Page = pageInfo
	response.Data = fileSpiderList
	this.Context.JSON(http.StatusOK, response)
	return
}
