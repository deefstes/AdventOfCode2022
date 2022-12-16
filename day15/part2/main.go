package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day15"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	lines := helpers.ReadInputFile()

	cave := day15.NewCave(lines)
	//fmt.Printf("%s\n", cave.String())
	beacon := cave.FindBeacon(0, 4000000)
	if beacon != nil {
		fmt.Printf("Beacon found at %s\n", beacon.String())
		fmt.Printf("Tuning frequency: %d\n", beacon.X*4000000+beacon.Y)
	}
}
