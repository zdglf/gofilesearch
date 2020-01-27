package service

import (
	"strings"
	"time"

	"github.com/zdglf/gofilesearch/base_struct"
	"github.com/zdglf/gofilesearch/db_model"
	"github.com/zdglf/gofilesearch/es_model"
	"github.com/zdglf/gofilesearch/file_sys"
	"github.com/zdglf/gofilesearch/util/flog"
)

const (
	typeFile = "file"
	typeSvn  = "svn"
	typeFtp  = "ftp"
)

type LoadFileService struct {
}

// 加载Gfile 并发送到ElasticsSearch
func insertToESFromGfile(gfile file_sys.GFile) (err error) {
	var fileName string
	var hashId string
	var content string
	if hashId, content, err = gfile.GetFileContent(); err != nil {
		flog.Println(err.Error())
		return
	}

	if fileName, err = gfile.GetFileName(); err != nil {
		flog.Println(err.Error())
		return
	}
	var exist bool
	if exist, err = es_model.IsDocumentExist(hashId); err != nil {
		flog.Println(err.Error())
		return
	}
	//文件存在
	if exist {
		return
	}

	flog.Println("post file:", gfile.GetAbFilePath())
	return es_model.InsertDocument(es_model.FileSearch{
		Id:         hashId,
		FileName:   fileName,
		Content:    content,
		ClickCount: 0,
		CreateAt:   base_struct.JsonTime(time.Now()),
		Url:        gfile.GetAbFilePath(),
	})
}

func (this *LoadFileService) LoadFile(fileModel *db_model.FileSpider) {
	var gfile file_sys.GFile = nil
	switch strings.ToLower(fileModel.Type) {
	case typeFile:
		gfile = &file_sys.LocalFile{FilePath: fileModel.Folder}

	}
	if gfile != nil {
		var err error
		var verify bool
		if verify, err = gfile.Verify(fileModel.Username, fileModel.Password); err != nil {
			flog.Println("verify error", err.Error())
			return
		}
		if !verify {
			flog.Println("wrong user password")
			return
		}

		flog.Println("start search ", gfile.GetAbFilePath())
		file_sys.WalkGFile(gfile,
			fileModel.ProcessSize,
			fileModel.Regular,
			fileModel.SizeLimit,
			insertToESFromGfile)
		flog.Println("end search ", gfile.GetAbFilePath())
	} else {
		flog.Println("no such file type, nothing start")
	}

}
