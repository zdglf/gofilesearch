package db_model

import "github.com/zdglf/gofilesearch/src/base_struct"

type FileSpider struct {
	Id              string               `xorm:"varchar(32) pk"`         //UUID
	Type            string               `xorm:"varchar(32) not null"`   //文件类型,ftp,file,svn,nfs,等
	Folder          string               `xorm:"text not null"`          //爬虫目录
	Username        string               `xorm:"varchar(64)"`            //用户名
	Password        string               `xorm:"varchar(64)"`            //密码
	CreateAt        base_struct.JsonTime `xorm:"created datetime"`       //记录创建时间
	Enable          int                  `xorm:"int not null default 0"` //是否启用 <=0 不启用
	Regular         string               `xorm:"text not null"`          //文件名匹配正则表达式 不能为空
	Timing          int                  `xorm:"int"`                    // <=0 不执行
	LastRunningTime base_struct.JsonTime `xorm:"datetime"`               //上次执行时间 执行后需要更新日期
	sizeLimit       int                  `xorm:"int"`                    //<=0 表示不限制
	processSize     int                  `xorm:"int"`                    //<=0 表示使用1个协程处理
}
