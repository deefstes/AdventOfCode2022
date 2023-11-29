package day23

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type Elf struct {
	id           int
	position     helpers.Coords
	proposedMove *helpers.Coords
}

func NewElf(id, x, y int) Elf {
	return Elf{
		id:       id,
		position: helpers.NewCoords(x, y),
	}
}

type ElfMob struct {
	elves        []Elf
	movePriority []string
	occupied     map[helpers.Coords]bool
}

func (e *ElfMob) String() string {
	var s strings.Builder
	c1, c2 := e.CalcLimits()

	for y := c1.Y; y <= c2.Y; y++ {
		for x := c1.X; x <= c2.X; x++ {
			if e.occupied[helpers.NewCoords(x, y)] {
				fmt.Fprint(&s, "#")
			} else {
				fmt.Fprint(&s, ".")
			}
		}
		fmt.Fprintln(&s)
	}

	return s.String()
}

func (e *ElfMob) CreatePNG(name string, topLeft helpers.Coords, botRight helpers.Coords, imgSize helpers.Coords) {
	width := botRight.X - topLeft.X + 1
	height := botRight.Y - topLeft.Y + 1
	pixelSize := helpers.NewCoords(imgSize.X/width, imgSize.Y/height)

	upLeft := image.Point{0, 0}
	// lowRight := image.Point{imgSize.X, imgSize.Y}
	lowRight := image.Point{pixelSize.X * width, pixelSize.Y * height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := topLeft.X; x <= botRight.X; x++ {
		for y := topLeft.Y; y < botRight.Y; y++ {
			var col color.Color
			if e.occupied[helpers.NewCoords(x, y)] {
				col = color.White
				// img.Set(x-topLeft.X, y-topLeft.Y, color.White)
			} else {
				col = color.Black
				// img.Set(x-topLeft.X, y-topLeft.Y, color.Black)
			}

			for px := 0; px < pixelSize.X; px++ {
				for py := 0; py < pixelSize.Y; py++ {
					img.Set(pixelSize.X*(x-topLeft.X)+px, pixelSize.Y*(y-topLeft.Y)+py, col)
				}
			}
		}
	}

	// newImage := resize.Resize(600, 600, img, resize.Bilinear)

	// Encode as PNG.
	f, _ := os.Create(name)
	png.Encode(f, img)
}

func (e *ElfMob) CalcLimits() (helpers.Coords, helpers.Coords) {
	limitN := math.MaxInt
	limitW := math.MaxInt
	limitS := 0
	limitE := 0
	for _, elf := range e.elves {
		limitN = helpers.Min(limitN, elf.position.Y)
		limitW = helpers.Min(limitW, elf.position.X)
		limitS = helpers.Max(limitS, elf.position.Y)
		limitE = helpers.Max(limitE, elf.position.X)
	}

	return helpers.NewCoords(limitW, limitN), helpers.NewCoords(limitE, limitS)
}

func (e *ElfMob) CountEmptySquares() int {
	c1, c2 := e.CalcLimits()
	width := c2.X - c1.X + 1
	height := c2.Y - c1.Y + 1
	return width*height - len(e.elves)
}

func (e *ElfMob) Solve(rounds int, showProgress bool) (int, int) {
	if rounds <= 0 {
		rounds = math.MaxInt
	}

	var roundsCnt int
	for i := 0; i < rounds; i++ {
		roundsCnt++
		// First have all elves propose their moves
		proposedMoves := make(map[helpers.Coords]int)
		for i := 0; i < len(e.elves); i++ {
			pm := e.Propose(&(e.elves[i]))
			if pm != nil {
				proposedMoves[*pm] += 1
			}
		}

		// If no elf proposed a move, we're done
		if len(proposedMoves) == 0 {
			break
		}

		// Now have all elves perform their proposed moves, if valid
		for i := 0; i < len(e.elves); i++ {
			if e.elves[i].proposedMove != nil && proposedMoves[*(e.elves[i]).proposedMove] == 1 {
				e.occupied[e.elves[i].position] = false
				e.elves[i].position = *(e.elves[i]).proposedMove
				e.occupied[e.elves[i].position] = true
			}
		}

		e.RotateMovePriority()

		if showProgress {
			fmt.Printf("Round %d\n%s\n", i+1, e.String())
			e.CreatePNG(fmt.Sprintf("frame%02d.png", i), helpers.NewCoords(-13, -14), helpers.NewCoords(121, 120), helpers.NewCoords(600, 600))
		}
	}

	return e.CountEmptySquares(), roundsCnt
}

func (e *ElfMob) Propose(elf *Elf) *helpers.Coords {
	elf.proposedMove = nil
	var hasNeighbours bool
	for _, n := range elf.position.Neighbours(true) {
		if e.occupied[n] {
			hasNeighbours = true
		}
	}
	if !hasNeighbours {
		return nil
	}

	for _, dir := range e.movePriority {
		switch dir {
		case "N":
			if !e.occupied[elf.position.Up(1)] && !e.occupied[elf.position.Up(1).Left(1)] && !e.occupied[elf.position.Up(1).Right(1)] {
				proposedMove := elf.position.Up(1)
				elf.proposedMove = &proposedMove
			}
		case "S":
			if !e.occupied[elf.position.Down(1)] && !e.occupied[elf.position.Down(1).Left(1)] && !e.occupied[elf.position.Down(1).Right(1)] {
				proposedMove := elf.position.Down(1)
				elf.proposedMove = &proposedMove
			}
		case "W":
			if !e.occupied[elf.position.Left(1)] && !e.occupied[elf.position.Left(1).Up(1)] && !e.occupied[elf.position.Left(1).Down(1)] {
				proposedMove := elf.position.Left(1)
				elf.proposedMove = &proposedMove
			}
		case "E":
			if !e.occupied[elf.position.Right(1)] && !e.occupied[elf.position.Right(1).Up(1)] && !e.occupied[elf.position.Right(1).Down(1)] {
				proposedMove := elf.position.Right(1)
				elf.proposedMove = &proposedMove
			}
		}
		if elf.proposedMove != nil {
			return elf.proposedMove
		}
	}

	return elf.proposedMove
}

func (e *ElfMob) RotateMovePriority() {
	e.movePriority = append(e.movePriority, e.movePriority[0])
	e.movePriority = e.movePriority[1:]
}

func NewElfMob(input []string) ElfMob {
	elfMob := ElfMob{
		movePriority: []string{"N", "S", "W", "E"},
		occupied:     make(map[helpers.Coords]bool),
	}

	var elfCnt int
	for row, line := range input {
		for col, c := range line {
			if c == '#' {
				elfCnt++
				elfMob.elves = append(elfMob.elves, NewElf(elfCnt, col, row))
				elfMob.occupied[helpers.NewCoords(col, row)] = true
			}
		}
	}

	return elfMob
}
