package es_model

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
)

const (
	// es 的数据库
	esIndex = "file_search"
	// es 的表
	esType = "doc"
)

func init() {

}

func InitEsClient() (client *elasticsearch7.Client, err error) {

	if client, err = elasticsearch7.NewDefaultClient(); err != nil {
		log.Println(err.Error())
		return
	}
	return
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
		Refresh:      "true",
	}
	var res *esapi.Response
	if res, err = req.Do(context.Background(), client); err != nil {
		return
	}
	if res.IsError() {
		err = errors.New(res.String())
		return
	}
	log.Println(res.String())
	return

}

func SearchDocument(keyword string) (result string, err error) {
	var client *elasticsearch7.Client
	if client, err = InitEsClient(); err != nil {
		return
	}
	req := esapi.SearchRequest{
		Index:        []string{esIndex},
		DocumentType: []string{esType},
	}
	var res *esapi.Response
	if res, err = req.Do(context.Background(), client); err != nil {
		return
	}
	if res.IsError() {
		err = errors.New(res.String())
		return
	}
	log.Println(res.String())
	return
}

func ConfigEsSetting() (err error) {
	var client *elasticsearch7.Client
	if client, err = InitEsClient(); err != nil {
		return
	}
	config := esConfig{
		esConfigProperties{
			esConfigContent{
				Type:           "text",
				Analyzer:       "ik_max_word",
				SearchAnalyzer: "ik_smart",
			},
			esConfigFileName{
				Type:           "text",
				Analyzer:       "ik_max_word",
				SearchAnalyzer: "ik_smart",
			},
		},
	}
	var jsonData []byte
	if jsonData, err = json.Marshal(config); err != nil {
		return
	}
	log.Println(string(jsonData))
	req := esapi.IndicesPutMappingRequest{
		Index:        []string{esIndex},
		DocumentType: esType,
		Body:         bytes.NewBuffer(jsonData),
	}
	var res *esapi.Response
	if res, err = req.Do(context.Background(), client); err != nil {
		return
	}
	if res.IsError() {
		err = errors.New(res.String())
		return
	}
	log.Println(res.String())
	return

}
