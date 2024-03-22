package storage

import "errors"

var (
	ErrUrlNotFound = errors.New("url not found")
	ErrUrlExist    = errors.New("url exists")
)
