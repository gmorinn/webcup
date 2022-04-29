package api

import (
	"log"
	fileapi "webcup/back/gen/fileapi"
)

// fileapi service example implementation.
// The example methods log the requests and return zero values.
type fileapisrvc struct {
	logger *log.Logger
}

// NewFileapi returns the fileapi service implementation.
func NewFileapi(logger *log.Logger) fileapi.Service {
	return &fileapisrvc{logger}
}
