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

	for i, v := range crt.X {
		fmt.Printf("After %d cycles: x=%d\n", i, v)
	}

	fmt.Printf("Magic number: %d\n", crt.X[19]*20+crt.X[59]*60+crt.X[99]*100+crt.X[139]*140+crt.X[179]*180+crt.X[219]*220)
}
