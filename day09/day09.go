package day09

import (
	"strconv"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type Rope struct {
	knots  []helpers.Coords
	visits map[string]bool
}

func NewRope(len int) Rope {
	visits := make(map[string]bool)
	visits["0,0"] = true
	var knots []helpers.Coords
	for i := 0; i < len; i++ {
		knots = append(knots, helpers.NewCoords(0, 0))
	}

	return Rope{
		knots:  knots,
		visits: visits,
	}
}

func (r *Rope) Move(move string) {
	components := strings.Split(move, " ")
	dir := components[0]
	dist, _ := strconv.Atoi(components[1])

	for i := 0; i < dist; i++ {
		// Move head
		switch dir {
		case "U":
			r.MoveKnot(0, r.knots[0].X, r.knots[0].Y-1)
		case "D":
			r.MoveKnot(0, r.knots[0].X, r.knots[0].Y+1)
		case "L":
			r.MoveKnot(0, r.knots[0].X-1, r.knots[0].Y)
		case "R":
			r.MoveKnot(0, r.knots[0].X+1, r.knots[0].Y)
		}
	}
}

func (r *Rope) MoveKnot(knot, x, y int) {
	r.knots[knot].X = x
	r.knots[knot].Y = y

	// Move next knot
	if len(r.knots)-1 <= knot {
		r.visits[r.knots[knot].String()] = true
		return
	}

	head := r.knots[knot]
	tail := r.knots[knot+1]
	if head.SimpleDist(tail) <= 1 {
		return
	}

	newx := tail.X
	newy := tail.Y
	if head.X == tail.X && head.Y != tail.Y {
		// vertical movement
		if head.Y > tail.Y {
			newy = tail.Y + 1
		}
		if head.Y < tail.Y {
			newy = tail.Y - 1
		}
	} else if head.X != tail.X && head.Y == tail.Y {
		// horizontal movement
		if head.X > tail.X {
			newx = tail.X + 1
		}
		if head.X < tail.X {
			newx = tail.X - 1
		}
	} else if head.X != tail.X && head.Y != tail.Y {
		// diagonal movement
		if head.Y > tail.Y {
			newy = tail.Y + 1
		}
		if head.Y < tail.Y {
			newy = tail.Y - 1
		}
		if head.X > tail.X {
			newx = tail.X + 1
		}
		if head.X < tail.X {
			newx = tail.X - 1
		}
	}

	r.MoveKnot(knot+1, newx, newy)
}

func (r *Rope) NumTailVisits() int {
	return len(r.visits)
}
