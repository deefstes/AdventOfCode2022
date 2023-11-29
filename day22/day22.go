package day22

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type Direction int

const (
	east Direction = iota
	south
	west
	north
)

func (d Direction) Left(number int) Direction {
	if number == 0 {
		return d
	}

	d = (4 + d - 1) % 4
	return d.Left(number - 1)
}

func (d Direction) Right(number int) Direction {
	if number == 0 {
		return d
	}
	d = (d + 1) % 4
	return d.Right(number - 1)
}

type Path struct {
	steps []any
}

func newPath(input string) Path {
	var path Path

	regex := regexp.MustCompile(`(\d+|R|L)`)
	comps := regex.FindAllStringSubmatch(input, -1)
	for _, c := range comps {
		if c[0] == "R" || c[0] == "L" {
			path.steps = append(path.steps, c[0])
		} else {
			num, _ := strconv.Atoi(c[0])
			path.steps = append(path.steps, num)
		}
	}

	return path
}

type Transformation struct {
	face      int
	rotations int
}

type CubeXForm struct {
	face       int
	neighbours map[Direction]Transformation
}

type Board struct {
	grid       helpers.Grid
	cube       bool
	cubeLayout [][]int
	cubeXForms map[int]CubeXForm
	start      helpers.Coords
	path       Path
}

func (b *Board) Move(cell helpers.Coords, dir Direction, dist int) (helpers.Coords, Direction) {
	newCoords := cell
	newDir := dir

	for i := 0; i < dist; i++ {
		var canMove bool
		curCoords := newCoords
		newCoords, newDir, canMove = b.Advance(newCoords, newDir)
		if !canMove {
			return curCoords, newDir
		}
	}

	return newCoords, newDir
}

func (b *Board) FaceSize() (int, int) {
	faceHeight := b.grid.Height / len(b.cubeLayout)
	faceWidth := b.grid.Width / len(b.cubeLayout[0])

	return faceWidth, faceHeight
}

func (b *Board) FindFace(c helpers.Coords) int {
	faceWidth, faceHeight := b.FaceSize()
	return b.cubeLayout[c.Y/faceHeight][c.X/faceWidth]
}

func (b *Board) FaceCoords(c helpers.Coords) helpers.Coords {
	faceWidth, faceHeight := b.FaceSize()
	return helpers.NewCoords(c.X%faceWidth, c.Y%faceHeight)
}

func (b *Board) GlobalCoords(fc helpers.Coords, face int) helpers.Coords {
	fw, fh := b.FaceSize()
	var fx, fy int

doubleloop:
	for y := 0; y < len(b.cubeLayout); y++ {
		for x := 0; x < len(b.cubeLayout[y]); x++ {
			if b.cubeLayout[y][x] == face {
				fx, fy = x, y
				break doubleloop
			}
		}
	}

	return helpers.NewCoords(fc.X+fx*fw, fc.Y+fy*fh)
}

func (b *Board) Advance(cell helpers.Coords, dir Direction) (newCell helpers.Coords, newDir Direction, canMove bool) {
	switch dir {
	case north:
		newCell = cell.Up(1)
	case east:
		newCell = cell.Right(1)
	case south:
		newCell = cell.Down(1)
	case west:
		newCell = cell.Left(1)
	}
	newCell = newCell.Wrap(b.grid.Width, b.grid.Height)
	newDir = dir
	curFace := b.FindFace(cell)
	if b.FindFace(newCell) != curFace {
		xform := b.cubeXForms[curFace]
		newFace := xform.neighbours[newDir]
		fw, fh := b.FaceSize()
		faceCoords := b.FaceCoords(cell)
		faceCoords = faceCoords.Rotate(-newFace.rotations, fw, fh)
		if newFace.rotations > 0 {
			newDir = newDir.Left(newFace.rotations)
		}
		if newFace.rotations < 0 {
			newDir = newDir.Right(-newFace.rotations)
		}

		switch newDir {
		case north:
			faceCoords = faceCoords.Up(1)
		case east:
			faceCoords = faceCoords.Right(1)
		case south:
			faceCoords = faceCoords.Down(1)
		case west:
			faceCoords = faceCoords.Left(1)
		}
		faceCoords = faceCoords.Wrap(fw, fh)
		newCell = b.GlobalCoords(faceCoords, newFace.face)
	}

	if b.grid.GetCell(newCell) == '#' {
		return cell, dir, false
	}
	return newCell, newDir, true
}

func (b *Board) Password() int64 {
	dir := east
	pos := b.start

	for _, step := range b.path.steps {
		switch step {
		case "R":
			dir = dir.Right(1)
		case "L":
			dir = dir.Left(1)
		default:
			pos, dir = b.Move(pos, dir, step.(int))
		}
	}

	return (int64(pos.Y)+1)*1000 + (int64(pos.X)+1)*4 + int64(dir)
}

func (b *Board) MakeCubeLayout() [][]int {
	// These values are hard coded for the sample input and puzzle input specifically.
	// A better developer than me would have calculated these in code from the input.
	var layout [][]int
	if b.grid.Width == 16 {
		// small map (sample.txt)
		layout = append(layout, []int{0, 0, 1, 0})
		layout = append(layout, []int{2, 3, 4, 0})
		layout = append(layout, []int{0, 0, 5, 6})
	} else {
		// large map (input.txt)
		layout = append(layout, []int{0, 1, 2})
		layout = append(layout, []int{0, 3, 0})
		layout = append(layout, []int{5, 4, 0})
		layout = append(layout, []int{6, 0, 0})
	}

	return layout
}

func (b *Board) MakeCubeXForms() map[int]CubeXForm {
	// These values are hard coded for the sample input and puzzle input specifically.
	// A better developer than me would have calculated these in code from the input.
	xforms := make(map[int]CubeXForm)
	if b.grid.Width == 16 {
		// small map (sample.txt)
		if b.cube {
			xforms[1] = CubeXForm{face: 1, neighbours: map[Direction]Transformation{
				north: {2, 2},
				east:  {6, 2},
				south: {4, 0},
				west:  {3, 1},
			}}
			xforms[2] = CubeXForm{face: 2, neighbours: map[Direction]Transformation{
				north: {1, 2},
				east:  {3, 0},
				south: {5, 2},
				west:  {6, -1},
			}}
			xforms[3] = CubeXForm{face: 3, neighbours: map[Direction]Transformation{
				north: {1, -1},
				east:  {4, 0},
				south: {5, 1},
				west:  {2, 0},
			}}
			xforms[4] = CubeXForm{face: 4, neighbours: map[Direction]Transformation{
				north: {1, 0},
				east:  {6, -1},
				south: {5, 0},
				west:  {3, 0},
			}}
			xforms[5] = CubeXForm{face: 5, neighbours: map[Direction]Transformation{
				north: {4, 0},
				east:  {6, 0},
				south: {2, 2},
				west:  {3, -1},
			}}
			xforms[6] = CubeXForm{face: 6, neighbours: map[Direction]Transformation{
				north: {4, 1},
				east:  {1, 2},
				south: {2, 1},
				west:  {5, 0},
			}}
		} else {
			xforms[1] = CubeXForm{face: 1, neighbours: map[Direction]Transformation{
				north: {5, 0},
				east:  {1, 0},
				south: {4, 0},
				west:  {1, 0},
			}}
			xforms[2] = CubeXForm{face: 2, neighbours: map[Direction]Transformation{
				north: {2, 0},
				east:  {3, 0},
				south: {2, 0},
				west:  {4, 0},
			}}
			xforms[3] = CubeXForm{face: 3, neighbours: map[Direction]Transformation{
				north: {3, 0},
				east:  {4, 0},
				south: {3, 0},
				west:  {2, 0},
			}}
			xforms[4] = CubeXForm{face: 4, neighbours: map[Direction]Transformation{
				north: {1, 0},
				east:  {2, 0},
				south: {5, 0},
				west:  {3, 0},
			}}
			xforms[5] = CubeXForm{face: 5, neighbours: map[Direction]Transformation{
				north: {4, 0},
				east:  {6, 0},
				south: {1, 0},
				west:  {6, 0},
			}}
			xforms[6] = CubeXForm{face: 6, neighbours: map[Direction]Transformation{
				north: {6, 0},
				east:  {5, 0},
				south: {6, 0},
				west:  {5, 0},
			}}
		}
	} else {
		// large map (input.txt)
		if b.cube {
			xforms[1] = CubeXForm{face: 1, neighbours: map[Direction]Transformation{
				north: {6, -1},
				east:  {2, 0},
				south: {3, 0},
				west:  {5, 2},
			}}
			xforms[2] = CubeXForm{face: 2, neighbours: map[Direction]Transformation{
				north: {6, 0},
				east:  {4, 2},
				south: {3, -1},
				west:  {1, 0},
			}}
			xforms[3] = CubeXForm{face: 3, neighbours: map[Direction]Transformation{
				north: {1, 0},
				east:  {2, 1},
				south: {4, 0},
				west:  {5, 1},
			}}
			xforms[4] = CubeXForm{face: 4, neighbours: map[Direction]Transformation{
				north: {3, 0},
				east:  {2, 2},
				south: {6, -1},
				west:  {5, 0},
			}}
			xforms[5] = CubeXForm{face: 5, neighbours: map[Direction]Transformation{
				north: {3, -1},
				east:  {4, 0},
				south: {6, 0},
				west:  {1, 2},
			}}
			xforms[6] = CubeXForm{face: 6, neighbours: map[Direction]Transformation{
				north: {5, 0},
				east:  {4, 1},
				south: {2, 0},
				west:  {1, 1},
			}}
		} else {
			xforms[1] = CubeXForm{face: 1, neighbours: map[Direction]Transformation{
				north: {4, 0},
				east:  {2, 0},
				south: {3, 0},
				west:  {2, 0},
			}}
			xforms[2] = CubeXForm{face: 2, neighbours: map[Direction]Transformation{
				north: {2, 0},
				east:  {1, 0},
				south: {2, 0},
				west:  {1, 0},
			}}
			xforms[3] = CubeXForm{face: 3, neighbours: map[Direction]Transformation{
				north: {1, 0},
				east:  {3, 0},
				south: {4, 0},
				west:  {3, 0},
			}}
			xforms[4] = CubeXForm{face: 4, neighbours: map[Direction]Transformation{
				north: {3, 0},
				east:  {5, 0},
				south: {1, 0},
				west:  {5, 0},
			}}
			xforms[5] = CubeXForm{face: 5, neighbours: map[Direction]Transformation{
				north: {6, 0},
				east:  {4, 0},
				south: {6, 0},
				west:  {4, 0},
			}}
			xforms[6] = CubeXForm{face: 6, neighbours: map[Direction]Transformation{
				north: {5, 0},
				east:  {6, 0},
				south: {5, 0},
				west:  {6, 0},
			}}
		}

	}

	return xforms
}

func NewBoard(input []string, cube bool) Board {
	board := Board{cube: cube}

	var instructionsIndex int
	var width int
	for i, row := range input {
		width = helpers.Max(width, len(row))
		if row == "" {
			instructionsIndex = i + 1
			break
		}
	}

	// Pad all rows to maximum width
	for row := 0; row < instructionsIndex; row++ {
		input[row] = fmt.Sprintf("%s%s", input[row], strings.Repeat(" ", width-len(input[row])))
	}
	board.grid = helpers.NewGrid(input[:instructionsIndex-1], func(c rune) any {
		return c
	})

	// Make board's cube mappings
	board.cubeLayout = board.MakeCubeLayout()
	board.cubeXForms = board.MakeCubeXForms()

	board.path = newPath(input[instructionsIndex])

	for col := 0; col < board.grid.Width; col++ {
		if board.grid.GetCell(helpers.NewCoords(col, 0)) == '.' {
			board.start = helpers.NewCoords(col, 0)
			break
		}
	}

	return board
}
