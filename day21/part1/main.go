package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day21"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	troop := day21.NewTroop(lines)
	answer := troop.SolveBasic()
	fmt.Printf("Root monkey's answer: %.f\n", answer)
}
