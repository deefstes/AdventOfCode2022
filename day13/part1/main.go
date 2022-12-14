package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/day13"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	signal := day13.NewSignal(lines)
	indices, pairs := signal.RightOrderedPairs()

	fmt.Printf("Pairs in right order:\n")
	for _, pair := range pairs {
		fmt.Println(pair.String())
	}
	var sum int
	for _, index := range indices {
		sum = sum + index
	}

	fmt.Printf("Sum of correct indices: %d\n", sum)
}
