package textdecoder

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetString(t *testing.T) {
	var err error
	var data []byte
	if data, err = readAllFile("data/test_gbk.txt"); err != nil {
		t.Error(err.Error())
	}
	var ret string
	if ret, err = GetString(data); err != nil {
		t.Error(err.Error())
	}
	if ret != "测试GBK" {
		t.Error(ret, "不等于 '测试GBK'")
	}

	if data, err = readAllFile("data/test_utf8.txt"); err != nil {
		t.Error(err.Error())
	}
	if ret, err = GetString(data); err != nil {
		t.Error(err.Error())
	}

	if ret != "测试UTF8" {
		t.Error(ret, "不等于 '测试UTF8'")
	}

	if data, err = readAllFile("data/test_utf8_nobom.txt"); err != nil {
		t.Error(err.Error())
	}
	if ret, err = GetString(data); err != nil {
		t.Error(err.Error())
	}

	if ret != "测试UTF8" {
		t.Error(ret, "不等于 '测试UTF8'")
	}

	if data, err = readAllFile("data/test_utf16le.txt"); err != nil {
		t.Error(err.Error())
	}
	if ret, err = GetString(data); err != nil {
		t.Error(err.Error())
	}

	if ret != "测试UTF16LE" {
		t.Error(ret, "不等于 '测试UTF16LE'")
	}

	if data, err = readAllFile("data/test_utf16be.txt"); err != nil {
		t.Error(err.Error())
	}
	if ret, err = GetString(data); err != nil {
		t.Error(err.Error())
	}

	if ret != "测试UTF16BE" {
		t.Error(ret, "不等于 '测试UTF16BE'")
	}

}

func readAllFile(fileName string) (data []byte, err error) {
	var f *os.File
	if f, err = os.Open(fileName); err != nil {
		return
	}
	defer f.Close()
	if data, err = ioutil.ReadAll(f); err != nil {
		return
	}
	return
}
