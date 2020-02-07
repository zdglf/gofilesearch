package main

import (
	_ "net/http/pprof"

	"github.com/zdglf/gofilesearch/routers"
)

func main() {
	r := routers.InitRouters()
	r.Run(":8090")
}
