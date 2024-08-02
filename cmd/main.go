package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yudai2929/asciiart"
)

const (
	filePath = "assets/input.png"
)

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	asciiArt, err := asciiart.Generate(file, 100)

	if err != nil {
		log.Fatalf("Failed to generate ASCII art: %v", err)
	}
	fmt.Println(asciiArt)
}
