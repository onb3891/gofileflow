package output

// Writer defines the interface for writing output data
type Writer interface {
	Write([]byte) error
	Close() error
}
