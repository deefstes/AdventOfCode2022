package day15

import (
	"fmt"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type SBPair struct {
	sensor helpers.Coords
	beacon helpers.Coords
}

func (sb *SBPair) Dist() int {
	return sb.sensor.ManhattanDist(sb.beacon)
}

type Cave struct {
	limitL  int
	limitR  int
	limitT  int
	limitB  int
	sensors map[string]SBPair
	beacons map[string]SBPair
}

func (c *Cave) String() string {
	var s strings.Builder
	for y := c.limitT; y <= c.limitB; y++ {
		for x := c.limitL; x <= c.limitR; x++ {
			coords := helpers.NewCoords(x, y)
			if _, ok := c.sensors[coords.String()]; ok {
				fmt.Fprintf(&s, "S")
			} else if _, ok := c.beacons[coords.String()]; ok {
				fmt.Fprintf(&s, "B")
				// } else if _, ok := c.nonBeacons[coords.String()]; ok {
				// 	fmt.Fprintf(&s, "#")
			} else {
				fmt.Fprintf(&s, ".")
			}
		}
		fmt.Fprintln(&s)
	}

	return strings.TrimRight(s.String(), "\n")
}

func NewCave(input []string) Cave {
	cave := Cave{}
	sensors := make(map[string]SBPair)
	beacons := make(map[string]SBPair)

	for _, line := range input {
		comps := strings.Split(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						line,
						"Sensor at x=", "",
					),
					" closest beacon is at x=", "",
				),
				" y=", "",
			),
			":",
		)
		s := helpers.NewCoordsFromString(comps[0], ",")
		b := helpers.NewCoordsFromString(comps[1], ",")
		sensors[s.String()] = SBPair{
			sensor: s,
			beacon: b,
		}
		beacons[b.String()] = SBPair{
			sensor: s,
			beacon: b,
		}

		sbDist := s.ManhattanDist(b)
		cave.limitL = helpers.Min(cave.limitL, s.X-sbDist)
		cave.limitR = helpers.Max(cave.limitR, s.X+sbDist)
		cave.limitT = helpers.Min(cave.limitT, s.Y-sbDist)
		cave.limitB = helpers.Max(cave.limitB, s.Y+sbDist)
	}

	cave.sensors = sensors
	cave.beacons = beacons

	return cave
}

func (c *Cave) CountNonBeaconsOnRow(row int) int {
	cnt := 0
	for x := c.limitL; x <= c.limitR; x++ {
		coords := helpers.NewCoords(x, row)
		if _, isBeacon := c.beacons[coords.String()]; isBeacon {
			continue
		}

		// Check distance from each sensor
		for _, s := range c.sensors {
			if s.sensor.ManhattanDist(coords) <= s.sensor.ManhattanDist(s.beacon) {
				cnt = cnt + 1
				break
			}
		}
	}

	return cnt
}

func (c *Cave) FindBeacon(lower, upper int) *helpers.Coords {
	for y := lower; y <= upper; y++ {
		for x := lower; x <= upper; x++ {
			current := helpers.NewCoords(x, y)
			inRange := false
			for _, s := range c.sensors {
				d := current.ManhattanDist(s.sensor)
				if d <= s.sensor.ManhattanDist(s.beacon) {
					inRange = true
					x += s.sensor.ManhattanDist(s.beacon) - d
					break
				}
			}

			if !inRange {
				return &current
			}
		}
	}

	return nil
}
