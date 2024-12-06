# yamlreader

**yamlreader** is a simple Go package designed to read and parse YAML files into Go structures. It provides a generic function to load YAML configuration files into any specified struct type.

## Features
 - Supports parsing any valid YAML file.
 - Generic function using Go 1.18+ generics for flexibility.
 - Minimal and efficient error handling.

---

## Installation

To use **yamlreader**, you need to have [GO](https://go.dev/doc/install) installed on your system.

1. Add the package to your project using `go get`:
```
go get -u github.com/itv-go/yamlreader
```

2. Import the package in your code:

```go
import github.com/itv-go/yamlreader
```

---

## Usage

### Basic Example

Suppose you have a YAML configuration file, `config.yaml`:

```yaml
app_name: MyApp
port: 8080
debug: true
```

And you want to load this file into a Go struct:
```go
package main

import (
    "fmt"
    "github.com/itv-go/yamlreader"
)

type Config struct {
    AppName string `yaml:"app_name"`
    Port    int    `yaml:"port"`
    Debug   bool   `yaml:"debug"`
}

func main() {
    var config Config

    // Load the YAML file
    if _, err := yamlreader.Read("config.yaml", &config); err != nil {
        fmt.Printf("Error loading config: %v\n", err)
        return
    }

    // Use the loaded configuration
    fmt.Printf("Loaded Config: %+v\n", config)
}
```

### Output
```
Loaded Config: {AppName:MyApp Port:8080 Debug:true}
```

---

## API Reference

### `func Read[T any](filePath string, config *T) (*T, error)`

Reads and parses a YAML file into the provided struct.

#### Parameters:
 - **filePath**: The path to the YAML file.
 - **config**: A pointer to the struct where the parsed data will be loaded.

#### Returns:
 - **\*T**: A pointer to the populated struct.
 - **error**: An error if the operation fails.

### Example:

```go
var myStruct MyStruct

result, err := yamlreader.Read("path/to/file.yaml", &myStruct)
if err != nil {
    log.Fatalf("Failed to read YAML: %v", err)
}
```

---

## Error Handling

This package uses `log.Fatalf` to terminate the program if there are issues opening or decoding the YAML file. You can modify the behavior by wrapping the Read function with custom error-handling logic if needed.

---

## Contributing

Contributions are welcome! Feel free to fork this repository and submit a pull request.

