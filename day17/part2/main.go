package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day17"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	chamber := day17.NewChamber(lines[0])
	const iterations = 1000000000000
	chamber.Simulate(iterations)
	fmt.Printf("After %d blocks:\n%s\n", iterations, chamber.String(true))
	fmt.Printf("The tower is %d units tall after %d rocks\n", chamber.Height(), iterations)
}
