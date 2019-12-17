package file_sys

import (
	"fmt"
	"regexp"
)

// General File 通用文件接口
type GFile interface {
	// 获取文件绝对路径
	GetAbFilePath() string
	// 获取文件大小
	GetFileSize() (fileSize int, err error)
	//获取文件名
	GetFileName() (fileName string, err error)
	// 验证账号密码
	Verify(username, password string) (isVerify bool, err error)
	// 返回文件hash Hex 值，和文件内容
	GetFileContent() (hashValue string, content string, err error)
	// 是否文件目录
	IsDir() (isDir bool, err error)
	// 列举文件
	ListFile() (fileList []GFile, err error)
}

// 是否文件大小超过限制
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

// 是否匹配正则表达
func isMatchString(re string, itemChild GFile) bool {
	if re != "" {
		var fileName string
		var err error
		if fileName,  err = itemChild.GetFileName();err!=nil{
			return false
		}
		if isOk, err := regexp.MatchString(re, fileName); err != nil {
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

// 遍历文件目录下的所有文件， 并调用 fileProcess 处理文件
func WalkGFile(file GFile, chanSize int, re string, sizeLimit int, fileProcess func(f GFile) error) {
	dirList := make([]GFile, 0, 100)
	if isDir, err := file.IsDir(); err != nil {
		if isDir {
			dirList = append(dirList, file)
		}
	} else {
		fmt.Println(err.Error())
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

				if isDir, err := itemChild.IsDir(); err != nil {
					fmt.Println(err.Error())
				} else {

					if isDir {
						//是目录，添加到待处理列表中
						dirList = append(dirList, itemChild)
					} else {
						//判断文件是否超过大小
						if isOverSizeLimit(sizeLimit, itemChild) {
							continue
						}
						//判断文件是否匹配正则表达式
						if !isMatchString(re, itemChild) {
							continue
						}

						wait_loop:
						for {
							//文件处理协程大于或等于，等待输出
							if len(buffList) >= cap(buffList) {
								<-buffList
								<-chanList
							} else {
								//继续创建协程，处理文件
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
	}
	//等待处理完剩余协程
	for i := 0; i < len(buffList); i++ {
		<-buffList
		<-chanList
	}

}

// 解析文件内容
func ParseFileContent(data []byte) string {
	return string(data)
}
