package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day25"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	var sum int
	for _, line := range lines {
		dval := day25.SnafuToDec(line)
		sum += dval
	}

	fmt.Printf("Sum in dec: %d\n", sum)
	fmt.Printf("Sum in SNAFU: %s\n", day25.DecToSnafu(sum))
}
