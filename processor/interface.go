package processor

// Processor defines the interface for processing data
type Processor interface {
	Process([]byte) ([]byte, error)
}
