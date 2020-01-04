package es_model

import "testing"
import "time"

import "github.com/zdglf/gofilesearch/src/base_struct"

func TestInitEsClient(t *testing.T) {
	if _, err := InitEsClient(); err != nil {
		t.Error(err.Error())
	}
}
func TestInsertDocument(t *testing.T) {
	f := FileSearch{
		Id:         "id",
		Content:    "content",
		Url:        "url",
		ClickCount: 0,
		FileName:   "fileName",
		CreateAt:   base_struct.JsonTime(time.Now()),
	}

	if err := InsertDocument(f); err != nil {
		t.Error(err.Error())
	}
}

func TestConfigEsSetting(t *testing.T) {
	if err := ConfigEsSetting(); err != nil {
		t.Error(err.Error())
	}
}
