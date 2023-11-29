package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day20"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	cb := day20.NewCircBuf(lines)
	coords := cb.CalculateCoordinates(811589153, 10)

	fmt.Printf("Coordinates: %d\n", coords)
}
