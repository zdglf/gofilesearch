package es_model

import (
	"github.com/zdglf/gofilesearch/base_struct"
)

type FileSearch struct {
	Id         string               `json:"id"`
	Content    string               `json:"content"`
	Url        string               `json:"url"`
	ClickCount int                  `json:"click_count"`
	FileName   string               `json:"file_name"`
	CreateAt   base_struct.JsonTime `json:"create_at"`
}
