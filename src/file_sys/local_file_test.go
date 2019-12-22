package file_sys

import (
	"os"
	"testing"
	"fmt"
	"path"
	"sync"
	"strconv"
)

func TestLocalFile_ListFile(t *testing.T) {
    dir, _ := os.Getwd()
	dir = path.Join(dir, "data")
	localFile := &LocalFile{dir}
	if localFileList, err := localFile.ListFile(); err != nil {
		t.Error(err.Error())
	} else {
		if len(localFileList) != 4{
			t.Error("data folder ListFile must be 4 but not " +strconv.Itoa(len(localFileList)))
		}

	}
}

func TestLocalFile_IsDir(t *testing.T) {
	dir, _ := os.Getwd()
	dir = path.Join(dir, "data")
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
	var count = 0
	countAddress := &count
	var lock = &sync.Mutex{}
	fileProcess := func(f GFile)error{
		lock.Lock()
		(*countAddress) ++
		lock.Unlock()
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
	if(count!=4){
		t.Error("data folder file number must be 4 but not "+ strconv.Itoa(count))
	}
}

func BenchmarkLocalFile_WalkGFile(b *testing.B) {

	for i:=0;i<b.N;i++{
		dir, _ := os.Getwd()
		dir = path.Join(dir, "data")
		localFile := &LocalFile{dir}
		fileProcess := func(f GFile)error{
			return nil
		}
		WalkGFile(localFile, 10, "(.*)", 0, fileProcess)
	}
}
