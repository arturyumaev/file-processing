package file_info

import (
	"errors"
	"fmt"
)

type HttpResponseErr struct {
	Error string `json:"error"`
}

var ErrNoFileNameSpecified = errors.New("no file name specified")
var ErrNoSuchFile = errors.New("no such file")
var ErrRequestTimeoutReached = errors.New("request timeout reached")
var ErrMethodNotAllowed = errors.New("method not allowed")
var ErrEmptyParameterName = errors.New("empty parameter: name")
var ErrRetrievingFile = errors.New("error while retrieving file")
var ErrFileNameLengthTooLong = fmt.Errorf("max file name length is: %d", MAX_FILE_NAME_LENGTH)
