package service

import (
	"testing"

	"github.com/zdglf/gofilesearch/db_model"
)

func TestLoadFileService_LoadFile(t *testing.T) {
	fls := &LoadFileService{}
	fileSpider := &db_model.FileSpider{
		Id:          "123",
		Type:        "file",
		Folder:      "/Users/zhangmike/Downloads",
		SizeLimit:   1024 * 1024,
		Regular:     "(pdf|txt|md|docx)$",
		ProcessSize: 20,
	}
	fls.LoadFile(fileSpider)

}
