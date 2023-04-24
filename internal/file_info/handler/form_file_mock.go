package handler

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

type FormFileMock struct {
	fieldname string
	filename  string
}

func (m *FormFileMock) Generate() (*bytes.Buffer, string) {
	if m.fieldname == "" {
		m.fieldname = FORM_FIELD_FILE_NAME
	}

	if m.filename == "" {
		m.filename = "temp_file"
	}

	filePath := "./"

	body := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(body)

	fileWriter, _ := bodyWriter.CreateFormFile(m.fieldname, m.filename)

	file, _ := os.OpenFile(filePath+m.filename, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	io.Copy(fileWriter, file)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	return body, contentType
}
