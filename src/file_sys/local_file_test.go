package file_sys

import "testing"

func TestLocalFile_ListFile(t *testing.T) {
    localFile := &LocalFile{"/Users/zhangmike/Documents/gopath/src/github.com/zdglf/gofilesearch"}
    if localFileList, err := localFile.ListFile();err!= nil{
        t.Error(err.Error())
    }else{
        for _,f := range(localFileList) {
            println(f.GetAbFilePath())
        }

    }
}

func TestLocalFile_IsDir(t *testing.T) {
    localFile := &LocalFile{"/Users/zhangmike/Documents/gopath/src/github.com/zdglf/gofilesearch"}
    if isDir, err := localFile.IsDir();err!= nil{
        t.Error(err.Error())
    }else{
        if(!isDir){
            t.Error("file should be Dir")
        }
    }
}
