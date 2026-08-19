package main

import (
	"fmt"
	"os"
)

// markdown_converter - Convert markdown formats
func markdown_converter(path string) {
	fmt.Println("========================================")
	fmt.Println("  Markdown-Converter")
	fmt.Println("  Convert markdown formats")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	markdown_converter(path)
}
