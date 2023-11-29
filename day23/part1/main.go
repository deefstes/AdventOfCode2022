package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day23"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	elfMob := day23.NewElfMob(lines)
	fmt.Printf("== Initial State ==\n%s\n", elfMob.String())
	emptySquares, rounds := elfMob.Solve(10, false)
	fmt.Printf("Number of empty squares after %d rounds: %d\n", rounds, emptySquares)
}
