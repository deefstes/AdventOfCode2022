package main

import (
	"fmt"
	"sort"

	"github.com/deefstes/AdventOfCode2022/day13"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()

	packets := day13.NewRawPackets(lines)
	packets = append(packets, day13.NewPacket("[[2]]"))
	packets = append(packets, day13.NewPacket("[[6]]"))
	sort.Slice(packets, func(i, j int) bool { return day13.Compare(packets[i], packets[j]) == 1 })

	key := 1
	for i, v := range packets {
		if v.String() == "[[2]]" || v.String() == "[[6]]" {
			key = key * (i + 1)
		}
	}

	fmt.Printf("Decoder key: %d\n", key)
}
