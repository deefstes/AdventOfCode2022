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

func (c Coords) Down(d int) Coords {
	return NewCoords(c.X, c.Y+d)
}

func (c Coords) Left(d int) Coords {
	return NewCoords(c.X-d, c.Y)
}

func (c Coords) Up(d int) Coords {
	return NewCoords(c.X, c.Y-d)
}

func (c Coords) Right(d int) Coords {
	return NewCoords(c.X+d, c.Y)
}

func (c Coords) Rotate(dir, width, height int) Coords {
	if dir == 0 {
		return c
	}
	if dir > 1 {
		c = c.Rotate(dir-1, width, height)
		dir = 1
	}
	if dir < -1 {
		c = c.Rotate(dir+1, width, height)
		dir = -1
	}

	// Transform
	c.X, c.Y = c.Y, c.X

	switch dir {
	case 1:
		// Flip x
		c.X = height - 1 - c.X
	case -1:
		// Flip Y
		c.Y = width - 1 - c.Y
	default:
		return NewCoords(0, 0)
	}

	return c
}

func (c Coords) Wrap(width, height int) Coords {
	c.X = (width + c.X) % width
	c.Y = (height + c.Y) % height
	return c
}

func (c1 *Coords) SimpleDist(c2 Coords) int {
	dx := Abs(c1.X - c2.X)
	dy := Abs(c1.Y - c2.Y)
	d := math.Sqrt(float64(dx*dx + dy*dy))

	return int(d)
}

func (c1 *Coords) ManhattanDist(c2 Coords) int {
	dx := Abs(c1.X - c2.X)
	dy := Abs(c1.Y - c2.Y)
	d := dx + dy

	return int(d)
}

func (c Coords) ManhattanNeighbourhood(dist int) []Coords {
	var neighbours []Coords
	neighbours = append(neighbours, c)
	for x := 1; x <= dist; x++ {
		for y := 0; y <= dist-x; y++ {
			neighbours = append(neighbours, c.Right(x).Down(y))
			neighbours = append(neighbours, c.Up(x).Right(y))
			neighbours = append(neighbours, c.Left(x).Up(y))
			neighbours = append(neighbours, c.Down(x).Left(y))
		}
	}

	return neighbours
}

func (c Coords) Neighbours(inclDiag bool) []Coords {
	var neighbours []Coords
	neighbours = append(neighbours, c.Up(1))
	neighbours = append(neighbours, c.Right(1))
	neighbours = append(neighbours, c.Down(1))
	neighbours = append(neighbours, c.Left(1))

	if inclDiag {
		neighbours = append(neighbours, c.Up(1).Right(1))
		neighbours = append(neighbours, c.Right(1).Down(1))
		neighbours = append(neighbours, c.Down(1).Left(1))
		neighbours = append(neighbours, c.Left(1).Up(1))
	}

	return neighbours
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
