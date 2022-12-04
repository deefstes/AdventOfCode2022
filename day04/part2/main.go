package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	runningCount := 0
	for _, line := range lines {
		ranges := strings.Split(line, ",")
		r1 := MakeRange(ranges[0])
		r2 := MakeRange(ranges[1])

		if r1.Overlap(r2) {
			fmt.Printf("%d-%d,%d-%d\n", r1.start, r1.end, r2.start, r2.end)
			runningCount = runningCount + 1
		}
	}

	fmt.Printf("Total overlaps: %d\n", runningCount)
}

type Range struct {
	start int
	end   int
}

func MakeRange(input string) Range {
	vals := strings.Split(input, "-")
	start, err := strconv.ParseInt(vals[0], 10, 64)
	if err != nil {
		log.Fatalf("reading input: %v", err)
	}
	end, err := strconv.ParseInt(vals[1], 10, 64)
	if err != nil {
		log.Fatalf("reading input: %v", err)
	}

	return Range{
		start: int(start),
		end:   int(end),
	}
}

func (r1 *Range) Overlap(r2 Range) bool {
	return (r1.start >= r2.start && r1.start <= r2.end) ||
		(r1.end >= r2.start && r1.end <= r2.end) ||
		(r2.start >= r1.start && r2.start <= r1.end) ||
		(r2.end >= r1.start && r2.end <= r1.end)
}
