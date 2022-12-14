package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	var elfGroups []ElfGroup
	for g := 0; g < len(lines)/3; g++ {
		elfGroups = append(elfGroups, MakeElfGroup(lines[g*3:g*3+3]))
	}

	runningCount := 0
	for group, eGroup := range elfGroups {
		dupes := eGroup.Duplicates()
		fmt.Printf("Group %d: ", group)
		for _, item := range dupes {
			fmt.Printf("%c - %d ", item, Priority(item))
			runningCount = runningCount + Priority(item)
		}
		fmt.Println()
	}

	fmt.Printf("Total priority: %d\n", runningCount)
}

type ElfGroup struct {
	rucksacks []string
}

func MakeElfGroup(rucksacks []string) ElfGroup {
	return ElfGroup{
		rucksacks: rucksacks,
	}
}

func (e *ElfGroup) Duplicates() []byte {
	dups := []byte(e.rucksacks[0])
	for i := 1; i < len(e.rucksacks); i++ {
		dups = FindDuplicates(dups, []byte(e.rucksacks[i]))
	}

	return dups
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

func FindDuplicates(slice1 []byte, slice2 []byte) []byte {
	dups := make(map[byte]int, 0)
	var retval []byte

	for _, item := range slice1 {
		if Contains(slice2, item) {
			dups[item] = dups[item] + 1
		}
	}

	for k, _ := range dups {
		retval = append(retval, k)
	}

	return retval
}
