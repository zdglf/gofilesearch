package file_sys

import (
	"fmt"
	"regexp"
)

type GFile interface {
	GetAbFilePath() string //获取文件绝对路径

	GetFileSize() (int, error) //获取文件大小

	GetFileName() string //货物文件名

	Verify(username, password string) (bool, error) //验证账号密码

	GetFileContent(hashAlgo string) (string, string, error) //返回文件hash Hex 值，和文件内容

	IsDir() (bool, error) //是否文件目录

	ListFile() ([]GFile, error) //列举文件
}

//是否文件大小超过限制
func isOverSizeLimit(sizeLimit int, itemChild GFile) bool {

	if sizeLimit >= 0 {
		if size, err := itemChild.GetFileSize(); err != nil {
			//获取文件大小异常跳过
			fmt.Println(err.Error())
			return true
		} else {
			//大于限制就跳过文件
			return size > sizeLimit
		}
	} else {
		return false
	}
}

func isMatchString(re string, itemChild GFile) bool {
	if re != "" {
		if isOk, err := regexp.MatchString(re, itemChild.GetFileName()); err != nil {
			//正则异常跳过
			fmt.Println(err.Error())
			return false
		} else {
			//正则匹配失败跳过
			return isOk
		}
	} else {
		return false
	}

}

func WalkGFile(file GFile, chanSize int, re string, sizeLimit int, fileProcess func(f GFile) error) {
	dirList := make([]GFile, 0, 100)
	if file.IsDir() {
		dirList = append(dirList, file)
	}
	if chanSize <= 0 {
		chanSize = 1
	}
	chanList := make(chan int, chanSize)
	buffList := make(chan int, chanSize)
	defer close(chanList)
	defer close(buffList)

	for len(dirList) != 0 {
		item := dirList[0]
		dirList = dirList[1:]
		if itemList, err := item.ListFile(); err != nil {
			fmt.Println(err.Error())
		} else {
			//遍历子目录的元素
			for _, itemChild := range itemList {

				fmt.Println(itemChild.GetFileName())

				if itemChild.IsDir() {
					dirList = append(dirList, itemChild)
				} else {
					if isOverSizeLimit(sizeLimit, itemChild) {
						continue
					}

					if !isMatchString(re, itemChild) {
						continue
					}

				wait_loop:
					for {
						//文件处理协程大于
						if len(buffList) >= cap(buffList) {
							<-buffList
							<-chanList
						} else {
							buffList <- 0
							go func(c chan int, f GFile, fp func(f GFile) error) {
								if err := fp(f); err != nil {
									fmt.Println(err.Error())
								}
								<-c
							}(chanList, itemChild, fileProcess)
							break wait_loop
						}
					}

				}
			}
		}
	}
	for i := 0; i < len(buffList); i++ {
		<-buffList
		<-chanList
	}

}
