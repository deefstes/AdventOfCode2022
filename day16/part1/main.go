package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day16"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	lines := helpers.ReadInputFile()
	cave := day16.NewCave(lines)
	pressure := cave.Solve(false)
	fmt.Printf("Pressure released: %d\n", pressure)
}
