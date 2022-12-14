package day14

import (
	"fmt"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type RockWall struct {
	limitL    int
	limitR    int
	limitB    int
	floor     bool
	obstacles map[helpers.Coords]bool
	sand      map[helpers.Coords]bool
}

func (r *RockWall) String() string {
	var s strings.Builder

	for y := 0; y <= r.limitB; y++ {
		for x := r.limitL; x <= r.limitR; x++ {
			if r.obstacles[helpers.NewCoords(x, y)] {
				//fmt.Fprint(&s, "█")
				fmt.Fprint(&s, "#,")
			} else if r.sand[helpers.NewCoords(x, y)] {
				//fmt.Fprint(&s, "○")
				fmt.Fprint(&s, "o,")
			} else {
				//fmt.Fprint(&s, " ")
				fmt.Fprint(&s, ".,")
			}
		}
		fmt.Fprint(&s, "\n")
	}

	return strings.TrimRight(s.String(), "\n")
}

func (r *RockWall) addBarrier(c1, c2 helpers.Coords) {
	xinc, yinc := 0, 0
	if c1.X != c2.X {
		xinc = 1
	}
	if c1.Y != c2.Y {
		yinc = 1
	}

	for x, y := helpers.Min(c1.X, c2.X), helpers.Min(c1.Y, c2.Y); x <= helpers.Max(c1.X, c2.X) && y <= helpers.Max(c1.Y, c2.Y); x, y = x+xinc, y+yinc {
		r.obstacles[helpers.NewCoords(x, y)] = true
		r.limitL = helpers.Min(r.limitL, x)
		r.limitR = helpers.Max(r.limitR, x)
		r.limitB = helpers.Max(r.limitB, y)
	}
}

func (r *RockWall) AddSand(e helpers.Coords) bool {
	// Check if sand has reached entry point
	if r.sand[e] {
		return false
	}

	for {
		if !r.floor {
			// Check if sand is out of bounds
			if e.Y > r.limitB {
				return false
			}
		}

		// Check if sand can move down
		if !r.obstacles[e.Down()] && !r.sand[e.Down()] && e.Down().Y < r.limitB+2 {
			e = e.Down()
			continue
		}

		// Check if sand can move down-left
		if !r.obstacles[e.Down().Left()] && !r.sand[e.Down().Left()] && e.Down().Y < r.limitB+2 {
			e = e.Down().Left()
			continue
		}

		// Check if sand can move down-right
		if !r.obstacles[e.Down().Right()] && !r.sand[e.Down().Right()] && e.Down().Y < r.limitB+2 {
			e = e.Down().Right()
			continue
		}

		// Sand settles
		r.sand[e] = true
		r.limitL = helpers.Min(r.limitL, e.X)
		r.limitR = helpers.Max(r.limitR, e.X)
		return true
	}
}

func NewRockWall(input []string, withFloor bool) RockWall {
	retval := RockWall{
		obstacles: make(map[helpers.Coords]bool),
		sand:      make(map[helpers.Coords]bool),
		limitL:    500,
		floor:     withFloor,
	}

	for _, v := range input {
		var coords []helpers.Coords
		for _, c := range strings.Split(v, " -> ") {
			coords = append(coords, helpers.NewCoordsFromString(c, ","))
			if len(coords) > 1 {
				retval.addBarrier(coords[len(coords)-2], coords[len(coords)-1])
			}
		}
	}

	return retval
}
