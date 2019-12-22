package file_sys

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"io"
	"strings"
)

func readDocxFile(r io.ReaderAt, size int) (conetnt string, err error) {
	var reader *zip.Reader
	if reader, err = zip.NewReader(r, int64(size)); err != nil {
		return
	}
	if conetnt, err = readText(reader.File); err != nil {
		return
	}

	return
}

func readText(files []*zip.File) (text string, err error) {
	var documentFile *zip.File
	if documentFile, err = retrieveWordDoc(files); err != nil {
		return
	}
	var documentReader io.ReadCloser
	if documentReader, err = documentFile.Open(); err != nil {
		return
	}
	text, err = wordDocToString(documentReader)
	return
}

func wordDocToString(reader io.Reader) (content string, err error) {
	xmlDecoder := xml.NewDecoder(reader)
	const READ_TAG = "t"
	isAppend := false
	for {
		var token xml.Token
		if token, err = xmlDecoder.Token(); err != nil {
			if err != io.EOF {
				return
			} else {
				err = nil
				break
			}
		}
		switch token.(type) {
		case xml.StartElement:

			se := token.(xml.StartElement)
			if se.Name.Local == READ_TAG {
				isAppend = true
			}
		case xml.EndElement:
			ee := token.(xml.EndElement)
			if ee.Name.Local == READ_TAG {
				isAppend = false
			}
		case xml.CharData:
			cd := token.(xml.CharData)
			if isAppend {
				content += strings.TrimSpace(string(cd))
			}
		}

	}

	return

}

func retrieveWordDoc(files []*zip.File) (file *zip.File, err error) {
	for _, f := range files {
		if f.Name == "word/document.xml" {
			file = f
		}
	}
	if file == nil {
		err = errors.New("document.xml file not found")
	}
	return
}
