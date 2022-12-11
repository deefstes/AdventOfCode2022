package main

import (
	"fmt"

	"github.com/deefstes/AdventOfCode2022/day11"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	troop := day11.NewTroop()
	for i := 0; i < len(lines); i = i + 7 {
		troop.AddMonkey(day11.NewMonkey(lines[i : i+6]))
	}

	fmt.Printf("Starting troop:\n%s\n\n", troop.String())

	for i := 0; i < 10000; i++ {
		troop.PlayRound(false, true)
		fmt.Printf("After round %d:\n%s\n\n", i+1, troop.String())
	}

	monkeys := troop.SortInspections()
	for _, m := range monkeys {
		fmt.Printf("%s inspected %d times\n", m.Name, m.Inspections)
	}
	fmt.Printf("Top two activity coefficient: %d\n", monkeys[0].Inspections*monkeys[1].Inspections)
}
