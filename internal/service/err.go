package service

import (
	"errors"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrTenderNotFound   = errors.New("tender not found")
	ErrServiceTypeError = errors.New("service type error")
	ErrStatusError      = errors.New("status error")
)
