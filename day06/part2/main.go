package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/day06"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	i := day06.FindNonRepeating([]byte(lines[0]), 14)
	fmt.Printf("First nonrepeating group of %d characters at position %d\n", 14, i)
}
