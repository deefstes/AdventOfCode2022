package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/day05"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	storage := day05.MakeStorage(lines)
	fmt.Println(storage.String())
	r := storage.DoWork(false)
	fmt.Println(storage.String())
	fmt.Printf("Final arrangement: %s\n", r)
}
