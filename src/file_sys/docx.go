package file_sys

import (
    "archive/zip"
    "bufio"
    "bytes"
    "encoding/xml"
    "errors"
    "io"
    "os"
    "strings"
)

type ReplaceDocx struct {
    zipReader *zip.ReadCloser
    content   string
}

func (r *ReplaceDocx) Editable() *Docx {
    return &Docx{
        files:   r.zipReader.File,
        content: r.content,
    }
}

func (r *ReplaceDocx) Close() error {
    return r.zipReader.Close()
}

type Docx struct {
    files   []*zip.File
    content string
}

func (d *Docx) Replace(oldString string, newString string, num int) (err error) {
    oldString, err = encode(oldString)
    if err != nil {
        return err
    }
    newString, err = encode(newString)
    if err != nil {
        return err
    }
    d.content = strings.Replace(d.content, oldString, newString, num)

    return nil
}

func (d *Docx) WriteToFile(path string) (err error) {
    var target *os.File
    target, err = os.Create(path)
    if err != nil {
        return
    }
    defer target.Close()
    err = d.Write(target)
    return
}

func (d *Docx) Write(ioWriter io.Writer) (err error) {
    w := zip.NewWriter(ioWriter)
    for _, file := range d.files {
        var writer io.Writer
        var readCloser io.ReadCloser

        writer, err = w.Create(file.Name)
        if err != nil {
            return err
        }
        readCloser, err = file.Open()
        if err != nil {
            return err
        }
        if file.Name == "word/document.xml" {
            writer.Write([]byte(d.content))
        } else {
            writer.Write(streamToByte(readCloser))
        }
    }
    w.Close()
    return
}

func readDocxFile(r io.ReaderAt, size int) (string, error) {

    reader, err := zip.NewReader(r, int64(size))
    if err != nil {
        return "", err
    }
    content, err := readText(reader.File)
    if err != nil {
        return "", err
    }

    return content, nil
}

func readText(files []*zip.File) (text string, err error) {
    var documentFile *zip.File
    documentFile, err = retrieveWordDoc(files)
    if err != nil {
        return text, err
    }
    var documentReader io.ReadCloser
    documentReader, err = documentFile.Open()
    if err != nil {
        return text, err
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
            if (err != io.EOF) {
                return
            } else {
                err = nil
                break
            }
        }
        switch token.(type){
        case xml.StartElement:

            se := token.(xml.StartElement)
            if(se.Name.Local==READ_TAG){
                isAppend = true
            }
        case xml.EndElement:
            ee := token.(xml.EndElement)
            if(ee.Name.Local==READ_TAG){
                isAppend = false
            }
        case xml.CharData:
            cd := token.(xml.CharData)
            if(isAppend){
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

func streamToByte(stream io.Reader) []byte {
    buf := new(bytes.Buffer)
    buf.ReadFrom(stream)
    return buf.Bytes()
}

func encode(s string) (string, error) {
    var b bytes.Buffer
    enc := xml.NewEncoder(bufio.NewWriter(&b))
    if err := enc.Encode(s); err != nil {
        return s, err
    }
    return strings.Replace(strings.Replace(b.String(), "<string>", "", 1), "</string>", "", 1), nil
}