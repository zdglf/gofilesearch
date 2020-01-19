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

func ConfigEsSetting() (err error) {
	var client *elasticsearch7.Client
	if client, err = InitEsClient(); err != nil {
		return
	}
	config := esConfig{
		esConfigProperties{
			esConfigInfo{
				Type:     "text",
				Analyzer: "ik_max_word",
			},
			esConfigInfo{
				Type:     "text",
				Analyzer: "ik_max_word",
			},
			esConfigInfo{
				Type:     "text",
				Analyzer: "ik_max_word",
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
