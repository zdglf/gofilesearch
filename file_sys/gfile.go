package file_sys

import (
	"bytes"
	"log"
	"path"
	"regexp"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/zdglf/gofilesearch/util/pool"
	"github.com/zdglf/gofilesearch/util/textdecoder"
)

const (
	typePdfFile      = ".pdf"
	typeDocFile      = ".doc"
	typeDocxFile     = ".docx"
	typeTextFile     = ".txt"
	typeMarkDownFile = ".md"
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

	if sizeLimit > 0 {
		if size, err := itemChild.GetFileSize(); err != nil {
			//获取文件大小异常跳过
			log.Println(err.Error())
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
		if fileName, err = itemChild.GetFileName(); err != nil {
			log.Println(err.Error())
			return false
		}
		if isOk, err := regexp.MatchString(re, fileName); err != nil {
			//正则异常跳过
			log.Println(err.Error())
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
		log.Println(err.Error())
	} else {

		if isDir {
			dirList = append(dirList, file)
		}
	}
	if chanSize <= 0 {
		chanSize = 1
	}

	threadPool := pool.NewPool(chanSize)

	for len(dirList) != 0 {
		item := dirList[0]
		dirList = dirList[1:]
		if itemList, err := item.ListFile(); err != nil {
			log.Println(err.Error())
		} else {
			//遍历子目录的元素
			for _, itemChild := range itemList {

				if isDir, err := itemChild.IsDir(); err != nil {
					log.Println(err.Error())
				} else {

					if isDir {
						//是目录，添加到待处理列表中
						dirList = append(dirList, itemChild)
					} else {
						//判断文件是否匹配正则表达式
						if !isMatchString(re, itemChild) {
							continue
						}
						//判断文件是否超过大小
						if isOverSizeLimit(sizeLimit, itemChild) {
							continue
						}
						//继续创建协程，处理文件
						threadPool.Add(1)
						go func(f GFile, fp func(f GFile) error) {
							defer threadPool.Done()
							if err := fp(f); err != nil {
								log.Println(err.Error())
							}
						}(itemChild, fileProcess)

					}
				}
			}
		}
	}
	//等待处理完剩余协程
	threadPool.Wait()
}

// 解析文件内容
func parseFileContent(data []byte, filePath string) (ret string) {

	reader := bytes.NewReader(data)
	dataSize := len(data)
	fileSuffix := strings.ToLower(path.Ext(filePath))
	switch fileSuffix {
	case typePdfFile:
		ret = parsePdfContent(reader, dataSize)
	case typeTextFile, typeMarkDownFile:
		ret, _ = textdecoder.GetString(data)
	case typeDocxFile:
		ret = parseDocxContent(reader, dataSize)
	}

	return
}

func parseDocxContent(reader *bytes.Reader, dataSize int) (ret string) {
	var err error
	if ret, err = readDocxFile(reader, dataSize); err != nil {
		log.Println(err.Error())
		return
	}

	return

}

func parsePdfContent(reader *bytes.Reader, dataSize int) (ret string) {
	var r *pdf.Reader
	var err error
	if r, err = pdf.NewReader(reader, int64(dataSize)); err != nil {
		log.Println(err.Error())
		return
	}
	var buf bytes.Buffer
	if b, err := r.GetPlainText(); err != nil {
		log.Println(err.Error())
		return
	} else {
		buf.ReadFrom(b)
		ret = buf.String()
		return
	}

}
