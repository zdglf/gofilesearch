package file_sys

import (
	"testing"
)

const (
	ftpServer   = "192.168.0.102:2121"
	ftpFilePath = ""
)

func TestFtpFile_GetAbFilePath(t *testing.T) {
	ftpFile := &FtpFile{FtpServer: ftpServer, FilePath: ftpFilePath}
	println(ftpFile.GetAbFilePath())
}

func TestFtpFile_ListFile(t *testing.T) {
	ftpFile := &FtpFile{FtpServer: ftpServer, FilePath: ftpFilePath}
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
	ftpFile := &FtpFile{FtpServer: ftpServer, FilePath: "/Download/a.pdf"}
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
	ftpFile := &FtpFile{FtpServer: ftpServer, FilePath: "/Download/a.pdf"}
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
	ftpFile := &FtpFile{FtpServer: ftpServer, FilePath: "/Download/a.pdf"}
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

func TestFtpFile_IsDir(t *testing.T) {
	ftpFile := &FtpFile{FtpServer: ftpServer, FilePath: "/Download/a.pdf"}
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
