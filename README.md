# Modular File Processor in Go

This project demonstrates a clean, extensible, and idiomatic Go architecture for processing files using dependency injection and small interfaces. It is designed for easy expansion and testing, making it a great reference for larger Go projects.

---

## Architecture Overview

The application is composed of three main modules, each defined by a small interface:

- **Input Module**: Reads data from a source (e.g., file, network, etc.)
- **Output Module**: Writes data to a destination (e.g., file, network, etc.)
- **Processor Module**: Processes or transforms data (e.g., toggling case, replacing log levels, etc.)

The `main.go` file wires these modules together using dependency injection, allowing you to swap implementations without changing the core logic.

---

## Module Connections

```mermaid
graph TD
    A[Input Module<br/>(input.Reader)] -- Read() --> B[Processor Module<br/>(processor.Processor)]
    B -- Process() --> C[Output Module<br/>(output.Writer)]
    C -- Write() --> D[Destination File]
```

- **Input Module** implements the `Reader` interface: `Read() ([]byte, error)` and `Close() error`.
- **Processor Module** implements the `Processor` interface: `Process([]byte) ([]byte, error)`.
- **Output Module** implements the `Writer` interface: `Write([]byte) error` and `Close() error`.

The `Context` struct in `main.go` holds these interfaces and coordinates the flow:

```go
type Context struct {
    reader    input.Reader
    writer    output.Writer
    processor processor.Processor
}
```

---

## How to Expand Each Module

### Input Module
- **Current Implementation:** `FileReader` (reads from a file)
- **To Add:** Implement the `Reader` interface (e.g., `NetworkReader`, `StdinReader`)
- **No Interface Change Needed:** Just add a new struct and implement the methods.

### Output Module
- **Current Implementation:** `FileWriter` (writes to a file)
- **To Add:** Implement the `Writer` interface (e.g., `NetworkWriter`, `StdoutWriter`)
- **No Interface Change Needed:** Just add a new struct and implement the methods.

### Processor Module
- **Current Implementations:**
  - `CaseToggleProcessor`: toggles character case
  - `LogLevelReplaceProcessor`: replaces log level codes with full words
- **To Add:** Implement the `Processor` interface (e.g., `EncryptProcessor`, `CompressProcessor`)
- **No Interface Change Needed:** Just add a new struct and implement the method.

---

## Example: Adding a New Processor

1. Create a new file in `processor/`, e.g., `reverse_processor.go`.
2. Implement the `Processor` interface:

```go
package processor

type ReverseProcessor struct{}

func NewReverseProcessor() *ReverseProcessor { return &ReverseProcessor{} }

func (p *ReverseProcessor) Process(data []byte) ([]byte, error) {
    // reverse logic here
}
```

3. Update `main.go` to allow selection of the new processor via CLI.

---

## How to Run

```sh
go run main.go -i input.txt -o output.txt -p case
# or
# go run main.go --input=input.txt --output=output.txt --processor=log
```

- `-i`/`--input`: Input file
- `-o`/`--output`: Output file
- `-p`/`--processor`: Processor type (`case` or `log`)

---

## Why This Design?

- **Separation of Concerns:** Each module does one thing.
- **Testability:** Interfaces make it easy to mock modules for testing.
- **Extensibility:** Add new functionality by implementing interfaces, not by changing them.
- **Maintainability:** Clear structure and small interfaces make the codebase easy to understand and grow.

---

## License

MIT 