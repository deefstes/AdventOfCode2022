package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day24"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()
	starttime := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(starttime))
	}()

	valley := day24.NewValley(lines)
	steps := valley.Solve(false)

	fmt.Printf("Expedition reached the exit in %d steps\n", steps)
}
