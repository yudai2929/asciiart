package main

import (
	"fmt"
	"log"

	"github.com/yudai2929/asciiart"
)

const (
	filePath      = "assets/input.png"
	encodedString = "data:image/png;base64,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"
)

func main() {
	asciiArt, err := asciiart.GenerateFromBase64(encodedString, asciiart.StdEncoding)
	// asciiArt, err := asciiart.Generate(filePath)
	if err != nil {
		log.Fatalf("Failed to generate ASCII art: %v", err)
	}
	fmt.Println(asciiArt)
}
