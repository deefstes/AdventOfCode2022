package main

import (
	"fmt"
	"sort"

	"github.com/deefstes/AdventOfCode2022/day07"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	fs := day07.ParseDir(&lines)
	sizes := fs.AllDirSizes()
	usedSpace := fs.CalcDirSize()
	freeSpace := int64(70000000) - usedSpace
	required := int64(30000000) - freeSpace

	fmt.Println("Total disk space: 70000000")
	fmt.Printf("Total used space: %d\n", usedSpace)
	fmt.Printf("Total free space: %d\n", freeSpace)
	fmt.Printf("  Required space: %d\n", required)
	var candidates []day07.File
	for _, v := range sizes {
		//fmt.Printf("%s (%d)\n", v.Name, v.Size)
		if v.Size > required {
			candidates = append(candidates, v)
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Size > candidates[j].Size
	})

	fmt.Println("Candidate directories for deletion:")
	for _, d := range candidates {
		fmt.Printf("%s (%d)\n", d.Name, d.Size)
	}
}
