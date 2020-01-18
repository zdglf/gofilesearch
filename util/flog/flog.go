package flog

import (
	"log"
)

func Println(v ...interface{}) {
	log.Println(v...)
}

func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}
