package es_model

import (
	"testing"
	"time"

	"encoding/json"
	"github.com/zdglf/gofilesearch/base_struct"
)

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

func TestIsExistDocument(t *testing.T) {
	if exist, err := IsDocumentExist("12345678"); err != nil {
		t.Error(err.Error())
	} else {
		if exist {
			t.Error("error should not exist")
		}
	}
}

func TestSearchDocument(t *testing.T) {

	if resultArray, pageInfo, err := SearchDocument("链接", 0); err != nil {
		t.Error(err.Error())
	} else {
		if data, err := json.Marshal(resultArray); err != nil {
			t.Error(err.Error())
		} else {
			println(string(data))
		}

		if data, err := json.Marshal(pageInfo); err != nil {
			t.Error(err.Error())
		} else {
			println(string(data))
		}

	}
}
