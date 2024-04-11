package readerbuilder

import "os"

// GetStdinReader returns the stdin input stream
func GetStdinReader() *os.File {
	return os.Stdin
}
