package helpers

type Grid struct {
	Width   int
	Height  int
	gridMap []any
}

func NewGrid(input []string, conv func(c rune) any) Grid {
	grid := Grid{
		Height: len(input),
		Width:  len(input[0]),
	}

	for _, str := range input {
		for _, c := range str {
			grid.gridMap = append(grid.gridMap, conv(c))
		}
	}

	return grid
}

func (g *Grid) GetCell(coords Coords) any {
	return g.gridMap[OnedFromCoords(coords, g.Width)]
}

func (g *Grid) GetCoords() []Coords {
	var coords []Coords

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			coords = append(coords, NewCoords(x, y))
		}
	}

	return coords
}

func (g *Grid) Neighbours(c Coords) []Coords {
	var neighbours []Coords

	if c.Up(1).Y >= 0 {
		neighbours = append(neighbours, c.Up(1))
	}
	if c.Right(1).X < g.Width {
		neighbours = append(neighbours, c.Right(1))
	}
	if c.Down(1).Y < g.Height {
		neighbours = append(neighbours, c.Down(1))
	}
	if c.Left(1).X >= 0 {
		neighbours = append(neighbours, c.Left(1))
	}

	return neighbours
}
