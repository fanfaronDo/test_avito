package handler

import "errors"

var (
	ErrUnsupportedRequest = errors.New("unsupported request")
	ErrUserIdInvalidType  = errors.New("user id is of invalid type")
)
