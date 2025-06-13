package processor

// CaseToggleProcessor implements Processor interface to toggle character cases
type CaseToggleProcessor struct{}

// NewCaseToggleProcessor creates a new case toggle processor
func NewCaseToggleProcessor() *CaseToggleProcessor {
	return &CaseToggleProcessor{}
}

// Process toggles the case of each character in the input
func (p *CaseToggleProcessor) Process(data []byte) ([]byte, error) {
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
