package es_model

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/zdglf/gofilesearch/base_struct"
	"github.com/zdglf/gofilesearch/util/flog"
)

type FileSearch struct {
	Id         string               `json:"id"`
	Content    string               `json:"content"`
	Url        string               `json:"url"`
	ClickCount int                  `json:"click_count"`
	FileName   string               `json:"file_name"`
	CreateAt   base_struct.JsonTime `json:"create_at"`
}

func InsertDocument(fileModel FileSearch) (err error) {
	var client *elasticsearch7.Client
	if client, err = InitEsClient(); err != nil {
		return
	}
	var jsonData []byte
	if jsonData, err = json.Marshal(fileModel); err != nil {
		return
	}
	req := esapi.IndexRequest{
		Index:        esIndex,
		DocumentType: esType,
		DocumentID:   fileModel.Id,
		Body:         bytes.NewReader(jsonData),
		Refresh:      "false",
	}
	var res *esapi.Response
	if res, err = req.Do(context.Background(), client); err != nil {
		return
	}
	if res.IsError() {
		err = errors.New(res.String())
		return
	}
	flog.Println(res.String())
	return

}

func SearchDocument(keyword string, pageIndex int) (result string, err error) {
	var client *elasticsearch7.Client
	if client, err = InitEsClient(); err != nil {
		return
	}
	pageSize := esPageCount
	pageFrom := esPageCount * pageIndex

	req := esapi.SearchRequest{
		Index:        []string{esIndex},
		DocumentType: []string{esType},
		Query:        keyword,
		From:         &pageFrom,
		Size:         &pageSize,
		Pretty:       true,
	}
	var res *esapi.Response
	if res, err = req.Do(context.Background(), client); err != nil {
		return
	}
	if res.IsError() {
		err = errors.New(res.String())
		return
	}
	result = res.String()
	flog.Println(res.String())
	return
}
