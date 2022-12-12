package day12

import (
	"fmt"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type Terrain struct {
	heightMap []rune
	width     int
	height    int
	Start     helpers.Coords
	End       helpers.Coords
}

func (t *Terrain) String() string {
	var retval strings.Builder
	for i, c := range t.heightMap {
		if i%t.width == 0 {
			fmt.Fprintf(&retval, "\n")
		}
		fmt.Fprintf(&retval, "%s", string(c))
	}

	return retval.String()
}

func (t *Terrain) ValidMove(c1, c2 helpers.Coords) bool {
	val1 := t.heightMap[helpers.OnedFromCoords(c1, t.width)]
	val2 := t.heightMap[helpers.OnedFromCoords(c2, t.width)]

	switch val1 {
	case 'S':
		if val2 == 'a' {
			return true
		}
	case 'z':
		if val2 == 'E' {
			return true
		}
	}

	return val2 <= val1+1
}

func (t *Terrain) Neighbours(c helpers.Coords) []helpers.Coords {
	var neighbours []helpers.Coords
	var nx, ny int

	// North
	nx, ny = c.X, c.Y-1
	if ny >= 0 {
		if t.ValidMove(c, helpers.Coords{X: nx, Y: ny}) {
			neighbours = append(neighbours, helpers.Coords{X: nx, Y: ny})
		}
	}

	// South
	nx, ny = c.X, c.Y+1
	if ny < t.height {
		if t.ValidMove(c, helpers.Coords{X: nx, Y: ny}) {
			neighbours = append(neighbours, helpers.Coords{X: nx, Y: ny})
		}
	}

	// East
	nx, ny = c.X+1, c.Y
	if nx < t.width {
		if t.ValidMove(c, helpers.Coords{X: nx, Y: ny}) {
			neighbours = append(neighbours, helpers.Coords{X: nx, Y: ny})
		}
	}

	// West
	nx, ny = c.X-1, c.Y
	if nx >= 0 {
		if t.ValidMove(c, helpers.Coords{X: nx, Y: ny}) {
			neighbours = append(neighbours, helpers.Coords{X: nx, Y: ny})
		}
	}

	return neighbours
}

func (g *Terrain) Solve() string {
	var leading helpers.Queue[helpers.Coords]
	leading.Add(g.Start)

	previous := make(map[helpers.Coords]*helpers.Coords)
	previous[g.Start] = nil

	for leading.Len() != 0 {
		current := leading.Get()
		if current == g.End {
			break
		}

		for _, n := range g.Neighbours(current) {
			if _, ok := previous[n]; !ok {
				leading.Add(n)
				previous[n] = &current
			}
		}
	}

	ret := []helpers.Coords{g.End}
	for n := previous[g.End]; n != nil; n = previous[*n] {
		ret = append(ret, *n)
	}

	var path strings.Builder
	for i := len(ret) - 1; i >= 0; i-- {
		fmt.Fprintf(&path, "%c", g.heightMap[helpers.OnedFromCoords(ret[i], g.width)])
	}

	return path.String()
}

func NewTerrain(lines []string) Terrain {
	var terrain Terrain
	terrain.height = len(lines)
	terrain.width = len(lines[0])

	for y, line := range lines {
		for x, c := range line {
			terrain.heightMap = append(terrain.heightMap, c)
			if c == 'S' {
				terrain.Start = helpers.Coords{X: x, Y: y}
			}
			if c == 'E' {
				terrain.End = helpers.Coords{X: x, Y: y}
			}
		}
	}

	return terrain
}
