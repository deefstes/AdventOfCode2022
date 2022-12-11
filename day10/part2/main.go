package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/day10"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	crt := day10.NewCRT()
	for _, line := range lines {
		crt.Input(line)
	}
	fmt.Println(crt.Draw("#", "."))
}
