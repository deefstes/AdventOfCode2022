package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var inputFileName string
	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.Parse()
	fmt.Printf("input file name: %s\n", inputFileName)

	// Read input file
	file, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalf("reading input file: %v", err)
	}
	lines := strings.Split(string(file), "\n")

	runningScore := 0
	for _, line := range lines {
		symbols := strings.Split(line, " ")
		if len(symbols) < 2 {
			continue
		}
		move := PickMove(symbols[0], symbols[1])
		runningScore = runningScore + MoveScore(move) + OutCome(symbols[0], move)
	}

	fmt.Printf("Final score: %d\n", runningScore)
}

func PickMove(opponent, desired string) string {
	switch opponent {
	case "A":
		switch desired {
		case "X":
			return "C"
		case "Y":
			return "A"
		case "Z":
			return "B"
		}
	case "B":
		switch desired {
		case "X":
			return "A"
		case "Y":
			return "B"
		case "Z":
			return "C"
		}
	case "C":
		switch desired {
		case "X":
			return "B"
		case "Y":
			return "C"
		case "Z":
			return "A"
		}
	}

	return "?"
}

func OutCome(opponent, player string) int {
	switch opponent {
	case "A":
		if player == "B" {
			return 6
		}
		if player == "A" {
			return 3
		}
	case "B":
		if player == "C" {
			return 6
		}
		if player == "B" {
			return 3
		}
	case "C":
		if player == "A" {
			return 6
		}
		if player == "C" {
			return 3
		}
	}

	return 0
}

func MoveScore(move string) int {
	switch move {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}

	return 0
}
