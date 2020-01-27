package api_model

type TaskIdRequest struct {
	Id string `json:"id" binding:"required"`
}

type TaskCreateRequest struct {
	Type        string `json:"type" binding:"required"`    //文件类型,ftp,file,svn,nfs,等
	Folder      string `json:"folder" binding:"required"`  //爬虫目录
	Username    string `json:"userName"`                   //用户名
	Password    string `json:"password"`                   //密码
	Enable      int    `json:"enable"`                     //是否启用定时执行 <=0 不启用
	Regular     string `json:"regular" binding:"required"` //文件名匹配正则表达式 不能为空
	Timing      int    `json:"timing"`                     //定时执行 <=0 不执行
	SizeLimit   int    `json:"sizeLimit"`                  //<=0 表示不限制
	ProcessSize int    `json:"processSize"`                //<=0 表示使用1个协程处理
}

type TaskModifyRequest struct {
	Id          string `json:"id" binding:"required"` //UUID
	Type        string `json:"type" `                 //文件类型,ftp,file,svn,nfs,等
	Folder      string `json:"folder"`                //爬虫目录
	Username    string `json:"userName"`              //用户名
	Password    string `json:"password"`              //密码
	Enable      int    `json:"enable"`                //是否启用定时执行 <=0 不启用
	Regular     string `json:"regular"`               //文件名匹配正则表达式 不能为空
	Timing      int    `json:"timing"`                //定时执行 <=0 不执行
	SizeLimit   int    `json:"sizeLimit"`             //<=0 表示不限制
	ProcessSize int    `json:"processSize"`           //<=0 表示使用1个协程处理
}

type TaskListRequest struct {
	PageIndex int `json:"pageIndex"`
}
