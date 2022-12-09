package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/day08"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	forest := day08.MakeForest(lines)
	fmt.Printf("%s\n", forest.String())
	fmt.Printf("Number of visible trees: %d", forest.CountVisible())
}
