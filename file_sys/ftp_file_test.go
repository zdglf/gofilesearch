package file_sys

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestFtpFile_GetAbFilePath(t *testing.T) {
	ftpFile := NewFtpFile("ftp://192.168.0.102:2121/")
	println(ftpFile.GetAbFilePath())
}

func TestFtpFile_ListFile(t *testing.T) {
	ftpFile := NewFtpFile("ftp://192.168.0.102:2121/")
	var verify bool
	var err error
	if verify, err = ftpFile.Verify("guest", "guest"); err != nil {
		t.Error(err.Error())
	}
	if !verify {
		t.Error("not verify")
		return
	}
	var list []GFile
	if list, err = ftpFile.ListFile(); err != nil {
		t.Error(err.Error())
	}
	if err != nil {
		t.Error(err)
	}
	for _, l := range list {
		println(l.GetAbFilePath())
	}
}

func TestFtpFile_GetFileContent(t *testing.T) {
	ftpFile := NewFtpFile("ftp://192.168.0.102:2121/Download/a.pdf")
	var verify bool
	var err error
	if verify, err = ftpFile.Verify("guest", "guest"); err != nil {
		t.Error(err.Error())
	}
	if !verify {
		t.Error("not verify")
		return
	}
	var hashValue string
	var content string
	if hashValue, content, err = ftpFile.GetFileContent(); err != nil {
		t.Error(err.Error())
	}
	println(content)
	println(hashValue)

}

func TestFtpFile_GetFileName(t *testing.T) {
	ftpFile := NewFtpFile("ftp://192.168.0.102:2121/Download/a.pdf")
	var verify bool
	var err error
	if verify, err = ftpFile.Verify("guest", "guest"); err != nil {
		t.Error(err.Error())
	}
	if !verify {
		t.Error("not verify")
		return
	}
	var fileName string

	if fileName, err = ftpFile.GetFileName(); err != nil {
		t.Error(err.Error())
	}
	println(fileName)
}

func TestFtpFile_GetFileSize(t *testing.T) {
	ftpFile := NewFtpFile("ftp://192.168.0.102:2121/Download/a.pdf")
	var verify bool
	var err error
	if verify, err = ftpFile.Verify("guest", "guest"); err != nil {
		t.Error(err.Error())
	}
	if !verify {
		t.Error("not verify")
		return
	}
	var size int

	if size, err = ftpFile.GetFileSize(); err != nil {
		t.Error(err.Error())
	}
	println(size)
}

func TestNewFtpFile(t *testing.T) {
	ftpFile := NewFtpFile("ftp://192.168.0.102:2121")

	println(ftpFile.GetAbFilePath())

	ftpFile = NewFtpFile("ftp://192.168.0.102/filePath")

	println(ftpFile.GetAbFilePath())
}

func TestFtpFile_IsDir(t *testing.T) {
	ftpFile := NewFtpFile("ftp://ftp.yxftp.com/影像电子书")
	var verify bool
	var err error
	if verify, err = ftpFile.Verify("guest", "guest"); err != nil {
		t.Error(err.Error())
	}
	if !verify {
		t.Error("not verify")
		return
	}
	var isDir bool

	if isDir, err = ftpFile.IsDir(); err != nil {
		t.Error(err.Error())
	}
	println(isDir)
}

func TestFtpFile_WalkGFile(t *testing.T) {
	dir := "ftp://192.168.0.102:2121/Download"
	localFile := NewFtpFile(dir)
	var count = 0
	countAddress := &count
	var lock = &sync.Mutex{}
	if isVerify, err := localFile.Verify("guest", "guest"); err != nil {
		t.Error(err.Error())
		return
	} else {
		if !isVerify {
			t.Error("password or name error")
		}
	}

	fileProcess := func(f GFile) error {
		lock.Lock()
		(*countAddress)++
		lock.Unlock()
		fmt.Println(f.GetAbFilePath())
		if hashValue, content, err := f.GetFileContent(); err != nil {
			t.Error(err.Error())
		} else {
			println(hashValue)
			println(content)
		}
		return nil
	}
	WalkGFile(localFile, 10, "(.*)", 0, fileProcess)
	if count != 4 {
		t.Error("data folder file number must be 4 but not " + strconv.Itoa(count))
	}
}
