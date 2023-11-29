package main

import (
	"fmt"
	"time"

	"github.com/deefstes/AdventOfCode2022/day22"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	board := day22.NewBoard(lines, true)
	password := board.Password()

	fmt.Printf("Password: %d\n", password)
}
