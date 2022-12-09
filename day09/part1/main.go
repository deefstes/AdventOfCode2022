package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/day09"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	r := day09.NewRope(2)
	for _, line := range lines {
		r.Move(line)
	}

	fmt.Printf("Number of spaces visited by the tail: %d\n", r.NumTailVisits())
}
