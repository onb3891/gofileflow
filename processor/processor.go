package processor

// ToggleCaseProcessor implements Processor interface to toggle character cases
type ToggleCaseProcessor struct{}

// NewToggleCaseProcessor creates a new toggle case processor
func NewToggleCaseProcessor() *ToggleCaseProcessor {
	return &ToggleCaseProcessor{}
}

// Process toggles the case of each character in the input
func (p *ToggleCaseProcessor) Process(data []byte) ([]byte, error) {
	result := make([]byte, len(data))
	for i, b := range data {
		if b >= 'a' && b <= 'z' {
			result[i] = b - 32 // Convert to uppercase
		} else if b >= 'A' && b <= 'Z' {
			result[i] = b + 32 // Convert to lowercase
		} else {
			result[i] = b // Keep non-alphabetic characters unchanged
		}
	}
	return result, nil
}
