package db_model

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"os"
)

const (
	sqlDriver   = "mysql"
	envUserName = "MYSQL_USER"     //Mysql 用户名
	envPassword = "MYSQL_PASSWORD" //Mysql 密码
	envHost     = "MYSQL_HOST"     //Mysql 服务器地址
	envPort     = "MYSQL_PORT"     //Mysql 端口号
	envDbName   = "MYSQL_DBNAME"   //Mysql 数据库名
)

func init() {
	initEngine(true)
}

func initEngine(needSync bool) (engine *xorm.Engine, err error) {
	if err = checkEnv(envUserName, envPassword, envHost, envPort, envDbName); err != nil {
		log.Println(err.Error())
		return
	}
	if engine, err = xorm.NewEngine(sqlDriver, os.ExpandEnv("$"+envUserName+":$"+envPassword+"@tcp($"+envHost+":$"+envPort+")/$"+envDbName+"?charset=utf8")); err != nil {
		return
	}
	engine.ShowSQL(true)
	if needSync {
		engine.Sync(new(FileSpider))
	}
	return

}

func checkEnv(v ...string) (err error) {
	for _, env := range v {
		if _, found := os.LookupEnv(env); !found {
			err = errors.New("env not set;" + env)
			return
		}
	}
	return

}
