package file_info

import "errors"

var ErrNoFileNameSpecified = errors.New("no file name specified")
var ErrNoSuchFile = errors.New("no such file")
