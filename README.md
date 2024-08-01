# asciiart

The `asciiart` package is a Go library for generating ASCII art from images.

## What is ASCII art?

ASCII art is a form of artistic expression where images are created using only the characters from the ASCII standard, which includes letters, numbers, and symbols. Originating in the early days of computing, it transforms plain text into visual designs, ranging from simple line art to complex compositions. ASCII art remains popular for its unique aesthetic and is used in various digital platforms and software environments to add creativity and visual interest.

## Features

- Supports JPEG and PNG image formats.
- Generates ASCII art from base64-encoded image data.
- Supports custom encoding types (standard and URL encoding).

## Installation

```bash
go get -u github.com/yudai2929/asciiart
```

## Usage

### Generate ASCII Art from an Image File

```go
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/yudai2929/asciiart"
)

func main() {
	file, err := os.Open("path/to/your/image.jpg")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	asciiArt, err := asciiart.Generate(file)
	if err != nil {
		log.Fatalf("Failed to generate ASCII art: %v", err)
	}
	fmt.Println(asciiArt)
}
```

### Generate ASCII Art from a Base64 Encoded Image

```go
asciiArt, err := asciiart.GenerateFromBase64(encodedString, asciiart.StdEncoding)
if err != nil {
	log.Fatalf("Failed to generate ASCII art: %v", err)
}
fmt.Println(asciiArt)
```

## Contributions

Contributions are welcome! Please feel free to submit pull requests, report bugs, or suggest new features.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/yudai2929/asciiart/blob/main/LICENSE) file for details.
