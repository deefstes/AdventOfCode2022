package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day14"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	lines := helpers.ReadInputFile()

	rockWall := day14.NewRockWall(lines, true)

	var units int
	for i := 0; ; i++ {
		if !rockWall.AddSand(helpers.NewCoords(500, 0)) {
			units = i
			break
		}
	}

	fmt.Printf("%s\nOverflow after %d units of sand\n", "rockWall.String()", units)
}
