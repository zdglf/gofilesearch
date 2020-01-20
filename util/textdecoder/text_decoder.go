package textdecoder

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/gogs/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func getText(data []byte, decoder *encoding.Decoder) (ret string, err error) {
	rInUTF8 := transform.NewReader(bytes.NewReader(data), decoder)
	var decBytes []byte
	if decBytes, err = ioutil.ReadAll(rInUTF8); err != nil {
		return
	} else {
		ret = string(decBytes)
		return
	}
}

func GetString(data []byte) (ret string, err error) {
	decoder := getTextDecoder(data[0:3])
	//UTF8 有Bom 格式返回decoder = nil
	if decoder == nil {
		ret = string(data[3:])
	} else {
		ret, err = getText(data, decoder)
	}
	return
}

func getTextDecoder(header []byte) (decoder *encoding.Decoder) {

	if len(header) < 3 {

		decoder = unicode.UTF8.NewDecoder()
		return
	}

	if header[0] == byte(0xEF) && header[1] == byte(0xBB) && header[2] == byte(0xBF) {

		decoder = nil
	} else if header[0] == byte(0xFE) && header[1] == byte(0xFF) {

		decoder = unicode.UTF16(unicode.BigEndian, unicode.UseBOM).NewDecoder()
	} else if header[0] == byte(0xFF) && header[1] == byte(0xFE) {

		decoder = unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
	} else {

		detector := chardet.NewTextDetector()
		var err error
		var r *chardet.Result
		if r, err = detector.DetectBest(header); err != nil {
			decoder = unicode.UTF8.NewDecoder()
			return
		}
		log.Println(r.Charset)
		if r.Charset == "UTF-8" {
			decoder = unicode.UTF8.NewDecoder()
		} else if r.Charset == "GB18030" {
			decoder = simplifiedchinese.GBK.NewDecoder()
		} else {
			decoder = unicode.UTF8.NewDecoder()
		}
		return

	}
	return

}
