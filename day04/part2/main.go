package main

import (
	"fmt"
	"strings"

	"github.com/deefstes/AdventOfCode2022/day04"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	runningCount := 0
	for _, line := range lines {
		ranges := strings.Split(line, ",")
		r1 := day04.MakeRange(ranges[0])
		r2 := day04.MakeRange(ranges[1])

		if r1.Overlap(r2) {
			fmt.Printf("%s,%s\n", r1.String(), r2.String())
			runningCount = runningCount + 1
		}
	}

	fmt.Printf("Total overlaps: %d\n", runningCount)
}
