package main

import (
	"fmt"
	"io"
	"log"

	"github.com/pborman/getopt/v2"

	"rtmp-go/input"
	"rtmp-go/output"
	"rtmp-go/processor"
)

// Context holds all the dependencies
type Context struct {
	reader    input.Reader
	writer    output.Writer
	processor processor.Processor
}

// NewContext creates a new context with all dependencies
func NewContext(
	reader input.Reader,
	writer output.Writer,
	processor processor.Processor,
) *Context {
	return &Context{
		reader:    reader,
		writer:    writer,
		processor: processor,
	}
}

// ProcessFile processes the entire file
func (c *Context) ProcessFile() error {
	defer c.reader.Close()
	defer c.writer.Close()

	for {
		// Read data
		data, err := c.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading: %v", err)
		}

		// Process data
		processed, err := c.processor.Process(data)
		if err != nil {
			return fmt.Errorf("error processing: %v", err)
		}

		// Write processed data
		if err := c.writer.Write(processed); err != nil {
			return fmt.Errorf("error writing: %v", err)
		}
	}

	return nil
}

func main() {
	inputFile := getopt.StringLong("input", 'i', "input.txt", "Input file name")
	outputFile := getopt.StringLong("output", 'o', "output.txt", "Output file name")
	processorType := getopt.StringLong("processor", 'p', "case", "Processor type: 'case' or 'log'")
	getopt.Parse()

	reader, err := input.NewFileReader(input.Config{
		Filename: *inputFile,
	})
	if err != nil {
		log.Fatalf("Error creating reader: %v", err)
	}

	writer, err := output.NewFileWriter(output.Config{
		Filename: *outputFile,
	})
	if err != nil {
		log.Fatalf("Error creating writer: %v", err)
	}

	var proc processor.Processor
	switch *processorType {
	case "case":
		proc = processor.NewCaseToggleProcessor()
	case "log":
		proc = processor.NewLogLevelReplaceProcessor()
	default:
		log.Fatalf("Unknown processor type: %s. Use 'case' or 'log'", *processorType)
	}

	ctx := NewContext(reader, writer, proc)

	if err := ctx.ProcessFile(); err != nil {
		log.Fatalf("Error processing file: %v", err)
	}

	fmt.Println("File processing completed successfully")
}
