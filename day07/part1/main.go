package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/day07"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	fs := day07.ParseDir(&lines)
	sizes := fs.AllDirSizes()
	var sum int64
	for _, v := range sizes {
		fmt.Printf("%s (%d)\n", v.Name, v.Size)
		if v.Size < 100000 {
			sum = sum + v.Size
		}
	}
	fmt.Printf("Total of directories with sub 100K size: %d\n", sum)
}
