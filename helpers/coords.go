package helpers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coords struct {
	X int
	Y int
}

func (c *Coords) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func (c Coords) Down() Coords {
	return NewCoords(c.X, c.Y+1)
}

func (c Coords) Left() Coords {
	return NewCoords(c.X-1, c.Y)
}

func (c Coords) Up() Coords {
	return NewCoords(c.X, c.Y-1)
}

func (c Coords) Right() Coords {
	return NewCoords(c.X+1, c.Y)
}

func (c1 *Coords) SimpleDist(c2 Coords) int {
	dx := Abs(c1.X - c2.X)
	dy := Abs(c1.Y - c2.Y)
	d := math.Sqrt(float64(dx*dx + dy*dy))

	return int(d)
}

func NewCoords(x, y int) Coords {
	return Coords{X: x, Y: y}
}

func NewCoordsFromString(input string, delim string) Coords {
	comps := strings.Split(input, delim)
	x, _ := strconv.Atoi(comps[0])
	y, _ := strconv.Atoi(comps[1])
	return Coords{X: x, Y: y}
}

func OnedFromXY(x, y, width int) int {
	return y*width + x
}

func OnedFromCoords(c Coords, width int) int {
	return c.Y*width + c.X
}
