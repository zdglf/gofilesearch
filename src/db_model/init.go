package db_model

import (
    "github.com/go-xorm/xorm"
    "os"
     _ "github.com/go-sql-driver/mysql"
)

func init()  {
    initEngine(true)
}

func initEngine(needSync bool)(engine *xorm.Engine, err error){
    if engine,err = xorm.NewEngine("mysql", os.ExpandEnv("$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_HOST:$MYSQL_PORT)/$MYSQL_DBNAME?charset=utf8"));err!=nil{
        return
    }
    engine.ShowSQL(true)
    if(needSync) {
        engine.Sync(new(FileSpider))
    }
    return



}