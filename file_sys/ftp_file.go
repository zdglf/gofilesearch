package file_sys

import (
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
	lastIndex := strings.LastIndex(this.FilePath, "/")
	folder := "/"
	if lastIndex != -1 {
		folder = this.FilePath[0:lastIndex]
	}
	name := this.FilePath[lastIndex+1:]

	if err = this.ftpClient.ChangeDir(folder); err != nil {
		return
	}
	var ftpFileSize int64 = 0
	if ftpFileSize, err = this.ftpClient.FileSize(name); err != nil {
		return
	}
	fileSize = int(ftpFileSize)
	return
}

func (this *FtpFile) initFtpClient() (err error) {
	if this.ftpClient != nil {
		return
	}
	this.ftpClient, err = ftp.Dial(this.FtpServer, ftp.DialWithTimeout(ftpTimeOutSecond*time.Second))
	return
}

//获取文件名
func (this *FtpFile) GetFileName() (fileName string, err error) {
	pathSplit := strings.Split(this.FilePath, "/")
	fileName = pathSplit[len(pathSplit)-1]
	return
}

// 验证账号密码
func (this *FtpFile) Verify(username, password string) (isVerify bool, err error) {
	if err = this.initFtpClient(); err != nil {
		return
	}
	if username == "" && password == "" {
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
	_, err = this.ftpClient.NameList(this.FilePath)
	//Maybe A file or Error
	if err != nil {
		// 读取文件
		if _, err = this.ftpClient.Retr(this.FilePath); err != nil {
			return
		} else {
			// is File
			isDir = false
			return
		}
	} else {
		isDir = true
		return
	}

}

// 列举文件
func (this *FtpFile) ListFile() (fileList []GFile, err error) {
	if err = this.initFtpClient(); err != nil {
		return
	}
	var entries []string
	if entries, err = this.ftpClient.NameList(this.FilePath); err != nil {
		return
	}
	for _, e := range entries {
		newPath := this.FilePath + "/" + e
		if this.FilePath == "/" || this.FilePath == "" {
			newPath = "/" + e
		}
		fileList = append(fileList, &FtpFile{
			FilePath:  newPath,
			FtpServer: this.FtpServer,
			ftpClient: this.ftpClient,
		})
	}

	return
}
