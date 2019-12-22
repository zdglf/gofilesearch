package file_sys

import (
	"os"
	"testing"
	"fmt"
	"path"
	"log"
)

func TestLocalFile_ListFile(t *testing.T) {
    dir, _ := os.Getwd()
	localFile := &LocalFile{dir}
	if localFileList, err := localFile.ListFile(); err != nil {
		t.Error(err.Error())
	} else {
		for _, f := range localFileList {
			println(f.GetAbFilePath())
		}

	}
}

func TestLocalFile_IsDir(t *testing.T) {
	dir, _ := os.Getwd()
	localFile := &LocalFile{dir}
	if isDir, err := localFile.IsDir(); err != nil {
		t.Error(err.Error())
	} else {
		if !isDir {
			t.Error("file should be Dir")
		}
	}
}

func TestLocalFile_WalkGFile(t *testing.T) {
	dir, _ := os.Getwd()
	dir = path.Join(dir, "data")
	localFile := &LocalFile{dir}
	fileProcess := func(f GFile)error{
		fmt.Println(f.GetAbFilePath())
		if hashValue, content, err := f.GetFileContent();err!=nil{
			t.Error(err.Error())
		}else{
			println(hashValue)
			println(content)
		}
		return nil
	}
	WalkGFile(localFile, 10, "(.*)", 0, fileProcess)
}

func BenchmarkLocalFile_WalkGFile(b *testing.B) {
	for i:=0;i<b.N;i++{
		dir, _ := os.Getwd()
		localFile := &LocalFile{dir}
		fileProcess := func(f GFile)error{
			return nil
		}
		WalkGFile(localFile, 10, "(.*)", 0, fileProcess)
	}
}
