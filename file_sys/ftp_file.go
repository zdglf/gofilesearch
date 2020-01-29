package file_sys

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/zdglf/gofilesearch/util/hash"
)

var _ GFile = (*FtpFile)(nil)

const (
	ftpTimeOutSecond = 30
)

type FtpFile struct {
	FtpServer string
	FilePath  string

	userName string
	password string

	ftpClient     *ftp.ServerConn
	fileSizeMutex *sync.Mutex
}

func NewFtpFile(filePath string) (gfile *FtpFile) {
	var err error
	var u *url.URL
	if u, err = url.Parse(filePath); err != nil {
		return nil
	}
	port := u.Port()
	if port == "" {
		port = "21"
	}
	hostName := u.Hostname()
	serverName := fmt.Sprintf("%s:%s", hostName, port)
	leftPath := u.EscapedPath()
	gfile = &FtpFile{FtpServer: serverName, FilePath: leftPath}
	return
}

// 获取文件绝对路径
func (this *FtpFile) GetAbFilePath() string {
	return fmt.Sprintf("ftp://%s%s", this.FtpServer, this.FilePath)
}

func (this *FtpFile) getParentFolderAndFileName() (folder, fileName string) {
	lastIndex := strings.LastIndex(this.FilePath, "/")
	folder = "/"
	if lastIndex != -1 {
		folder = this.FilePath[0:lastIndex]
	}
	fileName = this.FilePath[lastIndex+1:]
	return
}

// 获取文件大小
func (this *FtpFile) GetFileSize() (fileSize int, err error) {
	if err = this.initFtpClient(); err != nil {
		return
	}
	if this.fileSizeMutex == nil {
		this.fileSizeMutex = &sync.Mutex{}
	}
	this.fileSizeMutex.Lock()
	defer this.fileSizeMutex.Unlock()
	folder, name := this.getParentFolderAndFileName()

	if err = this.ftpClient.ChangeDir(folder); err != nil {
		return
	}
	var ftpFileSize int64 = 0
	if ftpFileSize, err = this.ftpClient.FileSize(name); err != nil {
		return
	}
	if err = this.ftpClient.ChangeDir("/"); err != nil {
		return
	}
	fileSize = int(ftpFileSize)
	return
}

func (this *FtpFile) initFtpClient() (err error) {
	if this.ftpClient != nil {
		return
	}
	if this.ftpClient, err = ftp.Dial(this.FtpServer, ftp.DialWithTimeout(ftpTimeOutSecond*time.Second)); err != nil {
		return
	}
	if this.userName == "" || this.password == "" {
		return
	}
	err = this.ftpClient.Login(this.userName, this.password)
	return
}

//获取文件名
func (this *FtpFile) GetFileName() (fileName string, err error) {
	_, fileName = this.getParentFolderAndFileName()
	return
}

// 验证账号密码
func (this *FtpFile) Verify(username, password string) (isVerify bool, err error) {
	if err = this.initFtpClient(); err != nil {
		return
	}
	this.userName = username
	this.password = password
	if username == "" || password == "" {
		isVerify = true
		return
	}
	if err = this.ftpClient.Login(username, password); err != nil {
		return
	}
	isVerify = true
	return
}

// 返回文件hash Hex 值，和文件内容
func (this *FtpFile) GetFileContent() (hashValue string, content string, err error) {
	if err = this.initFtpClient(); err != nil {
		return
	}
	var ftpResponse *ftp.Response
	if ftpResponse, err = this.ftpClient.Retr(this.FilePath); err != nil {
		return
	}
	var data []byte
	if data, err = ioutil.ReadAll(ftpResponse); err != nil {
		return
	}
	hashValue = hash.CalculateSha256AndHex(data)
	content = parseFileContent(data, this.GetAbFilePath())
	return
}

// 是否文件目录
func (this *FtpFile) IsDir() (isDir bool, err error) {
	if err = this.initFtpClient(); err != nil {
		return
	}

	folder, name := this.getParentFolderAndFileName()
	var entries []*ftp.Entry
	if entries, err = this.ftpClient.List(folder); err != nil {
		return
	}
	for _, e := range entries {
		if e.Name == name {
			if e.Type == ftp.EntryTypeFile {
				isDir = false
				return
			} else if e.Type == ftp.EntryTypeFolder {
				isDir = true
				return
			}
		}
	}
	err = errors.New("file not found")
	return
}

// 列举文件
func (this *FtpFile) ListFile() (fileList []GFile, err error) {
	if err = this.initFtpClient(); err != nil {
		return
	}
	var entries []*ftp.Entry
	if entries, err = this.ftpClient.List(this.FilePath); err != nil {
		return
	}
	for _, e := range entries {
		newPath := this.FilePath + "/" + e.Name
		if this.FilePath == "/" || this.FilePath == "" {
			newPath = "/" + e.Name
		}
		fileList = append(fileList, &FtpFile{
			FilePath:  newPath,
			FtpServer: this.FtpServer,
			password:  this.password,
			userName:  this.userName,
		})
	}
	println("end list folder", this.FilePath)

	return
}
