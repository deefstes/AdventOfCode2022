package day24

import (
	"math"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

const (
	maxSteps = 1024
)

type PosAtStep struct {
	helpers.Coords
	step int
}

func NewPosAtStep(c helpers.Coords, t int) PosAtStep {
	return PosAtStep{Coords: c, step: t}
}

type Blizzard uint8

func (b Blizzard) String() string {
	switch b {
	case 1:
		return ">"
	case 2:
		return "v"
	case 4:
		return "<"
	case 8:
		return "^"
	case 16:
		return "#"
	}
	return " "
}

func NewBlizzard(b byte) Blizzard {
	switch b {
	case '>':
		return 1
	case 'v':
		return 2
	case '<':
		return 4
	case '^':
		return 8
	case '#':
		return 16
	}
	panic(b)
}

type BlizzardBits uint8

func (b BlizzardBits) blizzards() (blizzards []Blizzard) {
	for bit := uint8(1); bit <= 16; bit <<= 1 {
		if uint8(b)&bit != 0 {
			blizzards = append(blizzards, Blizzard(bit))
		}
	}
	return
}

func (b *BlizzardBits) set(blizzard Blizzard) {
	*b |= BlizzardBits(blizzard)
}

func (b BlizzardBits) empty() bool {
	return uint8(b) == 0
}

type Valley struct {
	blizzards     [][][]BlizzardBits
	entrance      helpers.Coords
	exit          helpers.Coords
	width, height int
}

func NewValley(lines []string) Valley {
	var valley Valley
	valley.blizzards = make([][][]BlizzardBits, maxSteps+1)
	valley.blizzards[0] = make([][]BlizzardBits, len(lines))

	for row, line := range lines {
		valley.blizzards[0][row] = make([]BlizzardBits, len(line))
		for col := range line {
			if line[col] == '.' {
				if row == 0 {
					valley.entrance = helpers.NewCoords(0, col)
				}
				if row == len(lines)-1 {
					valley.exit = helpers.NewCoords(row, col)
				}
				continue
			}
			valley.blizzards[0][row][col] = BlizzardBits(NewBlizzard(line[col]))
		}
	}

	valley.height = len(valley.blizzards[0])
	valley.width = len(valley.blizzards[0][0])

	// Calculate future states of blizzards
	for i := 0; i < maxSteps; i++ {
		valley.CalcBlizzardState(i)
	}

	return valley
}

func (v *Valley) Solve(threeWay bool) int {
	steps1 := v.CalcRoute(NewPosAtStep(v.entrance, 0), v.exit)
	if !threeWay {
		return steps1
	}

	steps2 := v.CalcRoute(NewPosAtStep(v.exit, steps1), v.entrance)
	steps3 := v.CalcRoute(NewPosAtStep(v.entrance, steps2), v.exit)
	return steps3
}

func (v *Valley) CalcRoute(from PosAtStep, to helpers.Coords) int {
	seen := map[PosAtStep]bool{
		from: true,
	}
	q := []PosAtStep{from}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		possibleMoves := curr.Neighbours(false)
		possibleMoves = append(possibleMoves, curr.Coords)
		for _, n := range possibleMoves {
			next := NewPosAtStep(n, curr.step+1)
			if next.X < 0 || next.X >= len(v.blizzards[0]) {
				continue
			}
			if next.X == to.X && next.Y == to.Y {
				return next.step
			}
			if v.blizzards[next.step][next.X][next.Y].empty() && !seen[next] {
				seen[next] = true
				q = append(q, next)
			}
		}
	}
	return math.MaxInt
}

func (v *Valley) CalcBlizzardState(t int) {
	v.blizzards[t+1] = make([][]BlizzardBits, len(v.blizzards[t]))
	for i := range v.blizzards[t] {
		v.blizzards[t+1][i] = make([]BlizzardBits, len(v.blizzards[t][i]))
	}
	for y := range v.blizzards[t] {
		for x := range v.blizzards[t][y] {
			for _, b := range v.blizzards[t][y][x].blizzards() {
				newPos := v.MoveBlizzard(b, helpers.NewCoords(x, y))
				v.blizzards[t+1][newPos.Y][newPos.X].set(b)
			}
		}
	}
}

func (v *Valley) MoveBlizzard(b Blizzard, pos helpers.Coords) helpers.Coords {
	switch b {
	case 1:
		// >
		pos = pos.Right(1)
		if pos.X == v.width-1 {
			pos.X = 1
		}
	case 2:
		// v
		pos = pos.Down(1)
		if pos.Y == v.height-1 {
			pos.Y = 1
		}
	case 4:
		// <
		pos = pos.Left(1)
		if pos.X == 0 {
			pos.X = v.width - 2
		}
	case 8:
		// ^
		pos = pos.Up(1)
		if pos.Y == 0 {
			pos.Y = v.height - 2
		}
	case 16:
		// #
		return pos
	}

	return pos
}
