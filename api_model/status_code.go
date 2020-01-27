package api_model

const (
	CODE_MSG_SUCCESS = "成功"
	CODE_SUCCESS     = 0 //成功
	CODE_PARAM_ERROR = 1 //参数错误
	CODE_ES_ERROR    = 2 //ElasticSearch 错误
	CODE_DB_ERROR    = 3 //数据错误
)

type StatusCode struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
