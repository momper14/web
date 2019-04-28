package utils

import (
	"net/http"
	"strings"
)

// FileSystemStatic htp.FileSystem for static access
type FileSystemStatic struct {
	FileSystem http.FileSystem
}

// Open appen .html if no file-extension
func (fss FileSystemStatic) Open(name string) (http.File, error) {
	if !strings.Contains(name, ".") && len(name) > 1 {
		name += ".html"
	}

	return fss.FileSystem.Open(name)
}
