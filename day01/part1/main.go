package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
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

	runningCnt := int64(0)
	maxCalories := int64(0)

	for _, line := range lines {
		if line == "" {
			if runningCnt > maxCalories {
				maxCalories = runningCnt
			}
			runningCnt = 0
			continue
		}
		i, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("reading input: %v", err)
		}
		runningCnt = runningCnt + i
	}

	// Check final elf's total
	if runningCnt > maxCalories {
		maxCalories = runningCnt
	}

	// Output max calories
	fmt.Printf("Maximum calories: %d\n", maxCalories)
}
