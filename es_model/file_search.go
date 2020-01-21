package es_model

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/zdglf/gofilesearch/api_model"
	"github.com/zdglf/gofilesearch/base_struct"
	"github.com/zdglf/gofilesearch/util/flog"
)

const (
	esSearchId         = "id"          //对应json字段, 用于ES搜索解析
	esSearchContent    = "content"     //对应json字段， 用于ES搜索解析
	esSearchUrl        = "url"         //对应json字段，用于ES搜索解析
	esSearchClickCount = "click_count" //对应json字段，用于ES搜索解析
	esSearchFileName   = "file_name"   //对应json字段， 用于ES搜索解析
	esSearchCreateAt   = "create_at"   //对应json字段，用于ES搜索解析
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

func SearchDocument(keyword string, pageIndex int) (result *api_model.SearchResultResponse, err error) {
	var client *elasticsearch7.Client
	if client, err = InitEsClient(); err != nil {
		return
	}
	pageSize := esPageCount
	pageFrom := esPageCount * pageIndex

	var buf bytes.Buffer

	query := map[string]interface{}{
		"_source": []string{esSearchId, esSearchUrl, esSearchClickCount, esSearchFileName, esSearchCreateAt},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							esSearchUrl: keyword,
						},
					}, map[string]interface{}{
						"match": map[string]interface{}{
							esSearchContent: keyword,
						},
					},
				},
			},
		},
		"highlight": map[string]interface{}{
			// "pre_tags":  []string{"<strong>"},
			// "post_tags": []string{"</strong>"},
			"fields": map[string]interface{}{
				esSearchUrl:     map[string]interface{}{},
				esSearchContent: map[string]interface{}{},
			},
		},
	}

	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		return
	}

	var res *esapi.Response
	res, err = client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(esIndex),
		client.Search.WithBody(&buf),
		client.Search.WithSize(pageSize),
		client.Search.WithFrom(pageFrom),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)

	if res.IsError() {
		err = errors.New(res.String())
		return
	}

	var r map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return
	}
	result = &api_model.SearchResultResponse{}
	if esHits, ok := r["hits"].(map[string]interface{}); ok {
		//查找总数据
		if searchTotal, foundTotal := esHits["total"].(map[string]interface{}); foundTotal {
			var pageInfo = &api_model.Page{}
			pageInfo.Total = int(searchTotal["value"].(float64))
			pageInfo.Index = pageFrom
			pageInfo.Count = esPageCount
			result.Page = pageInfo
		}

		//查找搜索内容
		if searchArray, foundArray := esHits["hits"].([]interface{}); foundArray {
			var searchResultArray = make([]*api_model.SearchResult, 0)
			for _, value := range searchArray {
				if item, foundItem := value.(map[string]interface{}); foundItem {

					var searchResult = &api_model.SearchResult{}
					if docInfo, foundDoc := item["_source"].(map[string]interface{}); foundDoc {

						searchResult.Name = docInfo[esSearchFileName].(string)
						searchResult.Id = docInfo[esSearchId].(string)
						searchResult.Url = docInfo[esSearchUrl].(string)
						searchResult.CreateAt = docInfo[esSearchCreateAt].(string)
					}

					if highlight, foundHighLight := item["highlight"].(map[string]interface{}); foundHighLight {

						if highContent, foundHighLightContent := highlight[esSearchContent].([]interface{}); foundHighLightContent {

							for _, highContentItem := range highContent {
								searchResult.Desc = append(searchResult.Desc, highContentItem.(string))
							}
						}
						if highsearchUrl, foundHighLightURL := highlight[esSearchUrl].([]interface{}); foundHighLightURL {
							for _, highUrlItem := range highsearchUrl {
								searchResult.Desc = append(searchResult.Desc, highUrlItem.(string))
							}
						}
					}
					searchResultArray = append(searchResultArray, searchResult)
				}

			}
			result.Data = searchResultArray
		}

	}

	return
}
