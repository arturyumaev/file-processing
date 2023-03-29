package file_info

import "errors"

type HttpResponseErr struct {
	Error string `json:"error"`
}

var ErrNoFileNameSpecified = errors.New("no file name specified")
var ErrNoSuchFile = errors.New("no such file")
