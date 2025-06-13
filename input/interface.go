package input

// Reader defines the interface for reading input data
type Reader interface {
	Read() ([]byte, error)
	Close() error
}
