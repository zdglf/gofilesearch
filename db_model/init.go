package db_model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

const (
	sqlDriver = "mysql"
)

func init() {
	initEngine(true)
}

func initEngine(needSync bool) (engine *xorm.Engine, err error) {
	if engine, err = xorm.NewEngine(sqlDriver, os.ExpandEnv("$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_HOST:$MYSQL_PORT)/$MYSQL_DBNAME?charset=utf8")); err != nil {
		return
	}
	engine.ShowSQL(true)
	if needSync {
		engine.Sync(new(FileSpider))
	}
	return

}
