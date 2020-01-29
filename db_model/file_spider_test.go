package db_model

import (
	"github.com/zdglf/gofilesearch/api_model"
	"github.com/zdglf/gofilesearch/util/gen_data"
	"testing"
)

func TestInsertFileSpider(t *testing.T) {
	var err error
	var id string
	if id, err = gen_data.GenUUID(); err != nil {
		t.Error(err)
	}
	f := &FileSpider{
		Id:          id,
		Type:        "file",
		Folder:      "/Users/zhangmike/Downloads",
		SizeLimit:   1024 * 1024,
		Regular:     "(pdf|txt|md|docx)$",
		ProcessSize: 20,
	}
	var found bool
	if found, err = f.Exist(); err != nil {
		t.Error(err.Error())
		return
	}
	if found {
		t.Error("mysql contains ", f.Id)
		return
	}
	if err = f.Insert(); err != nil {
		t.Error(err.Error())
		return
	}

	if found, err = f.Exist(); err != nil {
		t.Error(err.Error())
		return
	}
	if !found {
		t.Error("insert error")
		return
	}

	f.Enable = 1
	if err = f.Update(); err != nil {
		t.Error("update error")
	}
	var list []*FileSpider
	var pageInfo *api_model.Page
	if list, pageInfo, err = QueryFileSpiderList(0); err != nil {
		t.Error(err.Error())
		return
	}
	if len(list) != 1 {
		t.Error("data size error")
		return
	}
	if pageInfo.Count != 1 {
		t.Error("data size error")
		return
	}

	if list[0].Enable != 1 {
		t.Error("data update error")
		return
	}

	if err = f.Delete(); err != nil {
		t.Error(err.Error())
		return
	}

	if list, pageInfo, err = QueryFileSpiderList(0); err != nil {
		t.Error(err.Error())
		return
	}
	if len(list) != 0 {
		t.Error("data size error")
		return
	}

	if pageInfo.Count != 0 {
		t.Error("data size error")
		return
	}

}
