package input

import (
	"os"
)

// Config holds configuration for the reader
type Config struct {
	Filename string
}

// FileReader implements Reader interface for file-based input
type FileReader struct {
	file *os.File
}

// NewFileReader creates a new file-based reader
func NewFileReader(config Config) (*FileReader, error) {
	file, err := os.Open(config.Filename)
	if err != nil {
		return nil, err
	}
	return &FileReader{file: file}, nil
}

// Read reads data from the file
func (r *FileReader) Read() ([]byte, error) {
	buffer := make([]byte, 1024)
	n, err := r.file.Read(buffer)
	if err != nil {
		return nil, err
	}
	return buffer[:n], nil
}

// Close closes the file
func (r *FileReader) Close() error {
	return r.file.Close()
}
