package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	duplicates := make(map[byte]int, 0)
	runningCount := 0
	for backpack, line := range lines {
		if line == "" {
			continue
		}
		items := []byte(line)
		size := len(items) / 2
		for _, item := range items[:size] {
			if Contains(items[size:], item) {
				duplicates[item] = duplicates[item] + 1
			}
		}
		fmt.Printf("Backpack %d: ", backpack)
		for k, v := range duplicates {
			fmt.Printf("%c(%d) - %d ", k, v, Priority(k))
			runningCount = runningCount + Priority(k)
		}
		fmt.Println()
		duplicates = make(map[byte]int, 0)
	}

	fmt.Printf("Total priority: %d\n", runningCount)
}

func Contains(compartment []byte, item byte) bool {
	for _, checkItem := range compartment {
		if checkItem == item {
			return true
		}
	}

	return false
}

func Priority(item byte) int {
	if item > 96 {
		return int(item) - 96
	}
	return int(item) - 64 + 26
}
