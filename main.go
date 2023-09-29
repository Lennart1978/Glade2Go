package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Check if enough arguments were passed
	if len(os.Args) != 3 {
		fmt.Println("Usage: programname source.glade target.go")
		return
	}

	// Read source and target filenames from the arguments
	sourceFile := os.Args[1]
	targetFile := os.Args[2]

	// Read file
	data, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("Error reading the source file:", err)
		return
	}

	// Process data: replace " with ' and remove spaces between <>
	processedData := strings.ReplaceAll(string(data), `"`, `'`)
	processedData = regexp.MustCompile(`>\s+<`).ReplaceAllString(processedData, "><")

	// Open the target file for writing
	target, err := os.Create(targetFile)
	if err != nil {
		fmt.Println("Error creating the target file:", err)
		return
	}
	defer target.Close()

	// Write the beginning of the target file with the string constant
	constant := "package ui\n\nconst ui = \"" + processedData

	// Remove the last newline character
	constant = strings.TrimSuffix(constant, "\n")

	// Add function GetUI to return the string
	constant += "\"\n\nfunc GetUI() string {\n\treturn ui\n}\n"

	_, err = target.WriteString(constant)
	if err != nil {
		fmt.Println("Error writing to the target file:", err)
		return
	}

	fmt.Println("Processing completed. The data has been saved in", targetFile, ".")
}
