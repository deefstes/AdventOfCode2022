package day04

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

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

func (r1 *Range) String() string {
	return fmt.Sprintf("%d-%d", r1.start, r1.end)
}

func (r1 *Range) Inside(r2 Range) bool {
	return (r1.start >= r2.start && r1.end <= r2.end)
}

func (r1 *Range) Overlap(r2 Range) bool {
	return (r1.start >= r2.start && r1.start <= r2.end) ||
		(r1.end >= r2.start && r1.end <= r2.end) ||
		(r2.start >= r1.start && r2.start <= r1.end) ||
		(r2.end >= r1.start && r2.end <= r1.end)
}
