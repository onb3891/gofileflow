package output

import (
	"os"
)

// Config holds configuration for the writer
type Config struct {
	Filename string
}

// FileWriter implements Writer interface for file-based output
type FileWriter struct {
	file *os.File
}

// NewFileWriter creates a new file-based writer
func NewFileWriter(config Config) (*FileWriter, error) {
	file, err := os.Create(config.Filename)
	if err != nil {
		return nil, err
	}
	return &FileWriter{file: file}, nil
}

// Write writes data to the file
func (w *FileWriter) Write(data []byte) error {
	_, err := w.file.Write(data)
	return err
}

// Close closes the file
func (w *FileWriter) Close() error {
	return w.file.Close()
}
