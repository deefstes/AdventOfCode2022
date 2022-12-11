package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type CRT struct {
	X []int
}

func NewCRT() CRT {
	return CRT{
		X: []int{1},
	}
}

func (crt *CRT) Input(input string) {
	comps := strings.Split(input, " ")

	switch comps[0] {
	case "noop":
		crt.X = append(crt.X, crt.X[len(crt.X)-1])
	case "addx":
		i, _ := strconv.Atoi(comps[1])
		crt.X = append(crt.X, crt.X[len(crt.X)-1])
		crt.X = append(crt.X, crt.X[len(crt.X)-1]+i)
	}
}

func (crt *CRT) Draw(on, off string) string {
	var output string

	for i, sprite := range crt.X {
		if i%40 == 0 {
			output = fmt.Sprintf("%s\n", output)
		}

		x := i % 40
		if x >= sprite-1 && x <= sprite+1 {
			output = fmt.Sprintf("%s%s", output, on)
		} else {
			output = fmt.Sprintf("%s%s", output, off)
		}
	}

	return output
}
