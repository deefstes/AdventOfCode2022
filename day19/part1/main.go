package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/deefstes/AdventOfCode2022/day19"
	"github.com/deefstes/AdventOfCode2022/helpers"
)

func main() {
	lines := helpers.ReadInputFile()
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed time: %s", time.Since(start))
	}()

	// regex := regexp.MustCompile(` Each (\w*) robot costs (\d*) (\w*)( and \d* \w*)?.`)

	for _, line := range lines {
		var blueprint day19.Blueprint
		blueprint.Id, _ = strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		formulae := strings.Split(strings.Split(line, ":")[1], ".")
		for _, formula := range formulae {
			c1 := strings.Split(formula, " robot costs ")
			c1[0] = strings.TrimSuffix(c1[0], " Each ")
			c2 := strings.Split(c1[1], " and ")
			// comps := regex.FindStringSubmatch(formula)
			_ = c1
			_ = c2
		}
	}

	// scan := func() error {
	// 	_, err := fmt.Fscanf(f, "Blueprint %d: "+
	// 		"Each ore robot costs %d ore. "+
	// 		"Each clay robot costs %d ore. "+
	// 		"Each obsidian robot costs %d ore and %d clay. "+
	// 		"Each geode robot costs %d ore and %d obsidian.\n",
	// 		&blueprint.Id,
	// 		&blueprint.Robots[0].Cost[day19.OreIndex],
	// 		&blueprint.Robots[1].Cost[day19.OreIndex],
	// 		&blueprint.Robots[2].Cost[day19.OreIndex], &blueprint.Robots[2].Cost[day19.ClayIndex],
	// 		&blueprint.Robots[3].Cost[day19.OreIndex], &blueprint.Robots[3].Cost[day19.ObsidianIndex],
	// 	)
	// 	return err
	// }
	// sum := 0
	// for err = scan(); err == nil; err = scan() {
	// 	for i := 0; i < day19.BotTypes; i++ {
	// 		blueprint.Robots[i].Output[i] = 1
	// 	}

	// 	s := day19.NewState(&blueprint)
	// 	key := day19.NewKey()
	// 	geodes := s.MaxGeodes(key)
	// 	quality := blueprint.Id * geodes
	// 	sum += quality

	// 	fmt.Printf("geodes: %d quality: %d sum: %d\n", geodes, quality, sum)

	// 	for *key != (day19.Key{}) {
	// 		m := s.M[*key]
	// 		//fmt.Println(key, m)
	// 		key = &m.NextKey
	// 	}
	// }
	// if !errors.Is(err, io.ErrUnexpectedEOF) {
	// 	log.Fatal(err)
	// }
}
