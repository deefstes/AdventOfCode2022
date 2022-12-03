package main

import (
	"fmt"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	runningScore := 0
	for _, line := range lines {
		symbols := strings.Split(line, " ")
		if len(symbols) < 2 {
			continue
		}
		runningScore = runningScore + MoveScore(symbols[1]) + OutCome(symbols[0], symbols[1])
	}

	fmt.Printf("Final score: %d\n", runningScore)
}

func OutCome(opponent, player string) int {
	switch opponent {
	case "A":
		if player == "Y" {
			return 6
		}
		if player == "X" {
			return 3
		}
	case "B":
		if player == "Z" {
			return 6
		}
		if player == "Y" {
			return 3
		}
	case "C":
		if player == "X" {
			return 6
		}
		if player == "Z" {
			return 3
		}
	}

	return 0
}

func MoveScore(move string) int {
	switch move {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}

	return 0
}
