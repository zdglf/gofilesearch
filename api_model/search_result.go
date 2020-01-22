package api_model

type SearchRequest struct {
	Keyword   string `json:"keyword" binding:"required"` //搜索关键字
	PageIndex int    `json:"pageIndex"`                  //检索开始页面
}

type SearchResult struct {
	Id       string   `json:"id"`       //文档ID SHA256
	Url      string   `json:"url"`      //文档地址
	Name     string   `json:"name"`     //文件名
	Desc     []string `json:"desc"`     //搜索到的关键字段
	CreateAt string   `json:"createAt"` //文件导入ES时间
}

type SearchResultResponse struct {
	*Page
	*StatusCode
	Data []*SearchResult `json:"data"`
}
