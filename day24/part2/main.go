package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day24"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	valley := day24.NewValley(lines)
	steps := valley.Solve(true)

	fmt.Printf("Mission complete in %d steps\n", steps)
}
