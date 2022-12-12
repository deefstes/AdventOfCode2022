package main

import (
	"fmt"
	"strings"

	"github.com/deefstes/AdventOfCode2022/day12"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	terrain := day12.NewTerrain(lines)
	fmt.Printf("Original map:\n%s\n\n", terrain.String())

	path := terrain.Solve()
	path = fmt.Sprintf("S%s", strings.TrimLeft(path, "Sa"))
	fmt.Println(path)
	fmt.Printf("Shortest path: %d\n", len(path)-1)
}
