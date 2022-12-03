package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
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

	var elfTotals []int64
	runningCnt := int64(0)

	for _, line := range lines {
		if line == "" {
			elfTotals = append(elfTotals, runningCnt)
			runningCnt = 0
			continue
		}
		i, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("reading input: %v", err)
		}
		runningCnt = runningCnt + i
	}

	elfTotals = append(elfTotals, runningCnt)

	// Sort elfTotals slice
	sort.Slice(elfTotals, func(i, j int) bool {
		return elfTotals[i] > elfTotals[j]
	})

	// Add top 3 elves' calories together
	total := elfTotals[0] + elfTotals[1] + elfTotals[2]
	fmt.Printf("Calories carried by top 3 elves: %d\n", total)
}
