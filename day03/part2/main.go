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

	groups := make([][]string, 0)
	for g := 0; g < len(lines)/3; g++ {
		group := make([]string, 0)
		for i := 0; i < 3; i++ {
			group = append(group, lines[g*3+i])
		}
		groups = append(groups, group)
	}

	runningCount := 0
	for group, eGroup := range groups {
		dupes := FindDuplicates([]byte(eGroup[0]), []byte(eGroup[1]), []byte(eGroup[2]))
		fmt.Printf("Group %d: ", group)
		for k, v := range dupes {
			fmt.Printf("%c(%d) - %d ", k, v, Priority(k))
			runningCount = runningCount + Priority(k)
		}
		fmt.Println()
	}

	fmt.Printf("Total priority: %d\n", runningCount)
}

func Contains(compartment []byte, item byte) bool {
	for _, checkItem := range compartment {
		if checkItem == item {
			return true
		}
	}

	return false
}

func Priority(item byte) int {
	if item > 96 {
		return int(item) - 96
	}
	return int(item) - 64 + 26
}

func FindDuplicates(rucksack1 []byte, rucksack2 []byte, rucksack3 []byte) map[byte]int {
	dups := make(map[byte]int, 0)

	for _, item := range rucksack1 {
		if Contains(rucksack2, item) && Contains(rucksack3, item) {
			dups[item] = dups[item] + 1
		}
	}

	return dups
}
