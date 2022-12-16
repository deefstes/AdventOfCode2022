package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day15"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	lines := helpers.ReadInputFile()

	cave := day15.NewCave(lines)
	//fmt.Printf("%s\n", cave.String())
	const row = 2000000
	nonB := cave.CountNonBeaconsOnRow(row)

	fmt.Printf("Positions in row %d where beacon can't be present: %d\n", row, nonB)
}
