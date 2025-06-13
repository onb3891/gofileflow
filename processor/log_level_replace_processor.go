package processor

import "strings"

// LogLevelReplaceProcessor implements Processor interface to replace log level codes with full words
type LogLevelReplaceProcessor struct{}

// NewLogLevelReplaceProcessor creates a new log level replace processor
func NewLogLevelReplaceProcessor() *LogLevelReplaceProcessor {
	return &LogLevelReplaceProcessor{}
}

// Process replaces ' I ' with ' INFO ', ' D ' with ' DEBUG ', ' E ' with ' ERROR ', and ' W ' with ' WARN '
func (p *LogLevelReplaceProcessor) Process(data []byte) ([]byte, error) {
	replacer := strings.NewReplacer(
		" I ", " INFO ",
		" D ", " DEBUG ",
		" E ", " ERROR ",
		" W ", " WARN ",
	)
	return []byte(replacer.Replace(string(data))), nil
}
