package db_model

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/zdglf/gofilesearch/base_struct"
)

const (
	FieldFileSpiderLastRunningTime = "last_running_time"
)

type FileSpider struct {
	Id              string               `xorm:"varchar(32) pk" json:"id"`             //UUID
	Type            string               `xorm:"varchar(32) not null" json:"type" `    //文件类型,ftp,file,svn,nfs,等
	Folder          string               `xorm:"text not null" json:"folder"`          //爬虫目录
	Username        string               `xorm:"varchar(64)" json:"userName"`          //用户名
	Password        string               `xorm:"varchar(64)" json:"password"`          //密码
	CreateAt        base_struct.JsonTime `xorm:"created datetime" json:"createAt"`     //记录创建时间
	Enable          int                  `xorm:"int not null default 0" json:"enable"` //是否启用定时执行 <=0 不启用
	Regular         string               `xorm:"text not null" json:"regular"`         //文件名匹配正则表达式 不能为空
	Timing          int                  `xorm:"int" json:"timing"`                    //定时执行 <=0 不执行
	LastRunningTime base_struct.JsonTime `xorm:"datetime" json:"lastRunningTime"`      //上次执行时间 执行后需要更新日期
	SizeLimit       int                  `xorm:"int" json:"sizeLimit"`                 //<=0 表示不限制
	ProcessSize     int                  `xorm:"int" json:"processSize"`               //<=0 表示使用1个协程处理
}

func (this *FileSpider) Insert() (err error) {
	var engine *xorm.Engine
	if engine, err = initEngine(false); err != nil {
		return
	}
	_, err = engine.Insert(this)
	return
}

func (this *FileSpider) UpdateAll() (err error) {
	var engine *xorm.Engine
	if engine, err = initEngine(false); err != nil {
		return
	}
	_, err = engine.ID(this.Id).Update(this)
	return
}

func (this *FileSpider) Update(columns ...string) (err error) {
	var engine *xorm.Engine
	if engine, err = initEngine(false); err != nil {
		return
	}
	_, err = engine.ID(this.Id).Cols(columns...).Update(this)
	return
}

func (this *FileSpider) Delete() (err error) {
	var engine *xorm.Engine
	if engine, err = initEngine(false); err != nil {
		return
	}
	_, err = engine.ID(this.Id).Delete(this)
	return
}

func (this *FileSpider) Exist() (found bool, err error) {
	var engine *xorm.Engine
	if engine, err = initEngine(false); err != nil {
		return
	}
	found, err = engine.ID(this.Id).Exist(this)
	return
}

func QueryFileSpiderList(pageIndex int) (dataList []*FileSpider, err error) {
	var engine *xorm.Engine
	if engine, err = initEngine(false); err != nil {
		return
	}
	dataList = make([]*FileSpider, 0)
	err = engine.Limit(envDBPageCount, pageIndex*envDBPageCount).Find(&dataList)
	return
}

func GetFileSpiderTotal() (total int, err error) {
	var engine *xorm.Engine
	if engine, err = initEngine(false); err != nil {
		return
	}
	var totalCount int64 = 0
	totalCount, err = engine.Count(&FileSpider{})
	total = int(totalCount)
	return
}

func GetFileSpiderById(id string) (fp *FileSpider, err error) {
	var engine *xorm.Engine
	if engine, err = initEngine(false); err != nil {
		return
	}
	var has bool
	if has, err = engine.ID(id).Get(fp); err != nil {
		return
	}
	if !has {
		err = errors.New("The Id Not Found")
	}
	return
}
