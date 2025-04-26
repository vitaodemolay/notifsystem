package internalerrors

import "errors"

var ErrInternal error = errors.New("internal server error")
var ErrNotFound error = errors.New("not found")
var ErrBadRequest error = errors.New("bad request")
