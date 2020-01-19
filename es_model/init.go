package es_model

import (
	"log"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
)

const (
	// es 的数据库
	esIndex = "file_search"
	// es 的表
	esType = "_doc"
	//默认页数
	esPageCount = 50

	esQueryKwContent = "content"
	esQueryKwUrl     = "url"
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
