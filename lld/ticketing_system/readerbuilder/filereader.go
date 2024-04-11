package readerbuilder

import (
	"errors"
	"os"
	"path/filepath"
)

// GetFileReader returns the input stream from the provided file path
func GetFileReader(filePath string) *os.File {
	path, err := filepath.Abs(filePath)
	if err != nil {
		panic(err)
	}
	inFile, err := os.Open(path)
	if err != nil {
		panic(errors.New(err.Error() + `: ` + path))
	}
	return inFile
}
