package handler

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"testing"
)

type FormFileMock struct {
	filename  string
	fieldname string
}

func (m *FormFileMock) Generate(t *testing.T) (*bytes.Buffer, string) {
	if m.fieldname == "" {
		m.fieldname = FORM_FIELD_FILE_NAME
	}

	if m.filename == "" {
		m.filename = "temp_file"
	}

	body := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(body)

	fileWriter, _ := bodyWriter.CreateFormFile(m.fieldname, m.filename)

	file, _ := os.OpenFile(t.TempDir()+m.filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	io.Copy(fileWriter, file)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	return body, contentType
}
