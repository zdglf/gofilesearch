package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zdglf/gofilesearch/service"
)

func AdminTaskCreate(c *gin.Context) {
	task := &service.AdminTaskService{Context: c}
	task.ResponseApiTaskCreate()
}

func AdminTaskExec(c *gin.Context) {
	task := &service.AdminTaskService{Context: c}
	task.ResponseApiTaskExec()

}

func AdminTaskList(c *gin.Context) {
	task := &service.AdminTaskService{Context: c}
	task.ResponseApiTaskList()
}

func AdminTaskDelete(c *gin.Context) {
	task := &service.AdminTaskService{Context: c}
	task.ResponseApiTaskDelete()
}

func AdminTaskModify(c *gin.Context) {
	task := &service.AdminTaskService{Context: c}
	task.ResponseApiTaskModify()

}
