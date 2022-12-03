package helpers

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInputFile() []string {
	var inputFileName string
	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.Parse()
	fmt.Printf("input file name: %s\n", inputFileName)

	// Read input file
	file, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalf("reading input file: %v", err)
	}
	lines := strings.Split(string(file), "\n")

	return lines
}
