package file_sys

import (
	_ "crypto/sha256"
	"crypto"
	"encoding/hex"
	"io/ioutil"
    "path"
	"os"
)

var _ GFile = (*LocalFile)(nil)

type LocalFile struct {
	filePath string
}

func (lf *LocalFile) GetAbFilePath() string {
	return lf.filePath
}

func (lf *LocalFile) GetFileSize() (fileSize int, err error) {
	var f *os.File
	var fileInfo os.FileInfo

	if f, err = os.Open(lf.filePath); err != nil {
		return
	}
	defer f.Close()
	if fileInfo, err = f.Stat(); err != nil {
		return
	}else{
        fileSize = int(fileInfo.Size())
    }
	return
}

func (lf *LocalFile) GetFileName() (fileName string, err error) {
	var f *os.File
	var fileInfo os.FileInfo

	if f, err = os.Open(lf.filePath); err != nil {
		return
	}
	defer f.Close()
	if fileInfo, err = f.Stat(); err == nil {
        return
	}else{
        fileName = fileInfo.Name()
    }
	return
}

func (lf *LocalFile) Verify(username, password string) (isVerify bool, err error) {
	return true, nil
}

func (lf *LocalFile) GetFileContent() (hashValue string, content string, err error) {
	var f *os.File
	if f, err = os.Open(lf.filePath); err != nil {
		return
	}
	defer f.Close()
	var data []byte
	if data, err = ioutil.ReadAll(f); err != nil {
		return
	}
	var hash = crypto.SHA256.New()
	hash.Write(data)
	hashValue = hex.EncodeToString(hash.Sum(nil))
	content = parseFileContent(data, lf.GetAbFilePath())
	return

}

func (lf *LocalFile) IsDir() (isDir bool, err error) {

	var f *os.File
	var fileInfo os.FileInfo
	isDir = false
	if f, err = os.Open(lf.filePath); err != nil {
		return
	}
	defer f.Close()
	if fileInfo, err = f.Stat(); err != nil {
		return
	}else{
        isDir = fileInfo.IsDir()
    }
	return

}

func (lf *LocalFile) ListFile() (fileList []GFile, err error) {
	var f *os.File
	var childNameList []string
	if f, err = os.Open(lf.filePath); err != nil {
		return
	}
	defer f.Close()
	if childNameList, err = f.Readdirnames(0); err != nil {
        return
	}else{
        for _, childName := range childNameList {
            fileList = append(fileList, &LocalFile{path.Join(lf.filePath,childName)})
        }
    }
	return

}
