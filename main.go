package main

import (
	"github.com/zdglf/gofilesearch/routers"
)

func main() {
	r := routers.InitRouters()
	r.Run()
}
