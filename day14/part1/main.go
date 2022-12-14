package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/day14"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	rockWall := day14.NewRockWall(lines, false)

	var units int
	for i := 0; ; i++ {
		if !rockWall.AddSand(helpers.NewCoords(500, 0)) {
			units = i
			break
		}
	}

	fmt.Printf("%s\nOverflow after %d units of sand\n", rockWall.String(), units)
}
