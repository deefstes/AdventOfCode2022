package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

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
