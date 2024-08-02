package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yudai2929/asciiart"
)

const (
	filePath = "assets/example.jpg"
)

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	asciiArt, err := asciiart.Generate(file, asciiart.WithWidth(200))
	if err != nil {
		log.Fatalf("Failed to generate ASCII art: %v", err)
	}

	fmt.Println(asciiArt)
}
