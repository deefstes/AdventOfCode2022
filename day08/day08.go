package day08

import (
	"fmt"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type Forest struct {
	width  int
	height int
	trees  []int
}

func (f *Forest) String() string {
	var retval string
	for i := 0; i < len(f.trees); i++ {
		if i%f.width == 0 {
			retval = fmt.Sprintf("%s\n", retval)
		}
		retval = fmt.Sprintf("%s%d", retval, f.trees[i])
	}

	return strings.TrimLeft(retval, "\n")
}

func (f *Forest) IsVisible(x, y int) bool {
	refTree := f.trees[helpers.OnedFromTwod(x, y, f.width)]

	// Check for northern visibility
	visCheck := true
	for yy := y - 1; yy >= 0; yy-- {
		if f.trees[helpers.OnedFromTwod(x, yy, f.width)] >= refTree {
			visCheck = false
		}
	}
	if visCheck {
		return true
	}

	// Check for southern visibility
	visCheck = true
	for yy := y + 1; yy < f.height; yy++ {
		if f.trees[helpers.OnedFromTwod(x, yy, f.width)] >= refTree {
			visCheck = false
		}
	}
	if visCheck {
		return true
	}

	// Check for western visibility
	visCheck = true
	for xx := x - 1; xx >= 0; xx-- {
		if f.trees[helpers.OnedFromTwod(xx, y, f.width)] >= refTree {
			visCheck = false
		}
	}
	if visCheck {
		return true
	}

	// Check for eastern visibility
	visCheck = true
	for xx := x + 1; xx < f.width; xx++ {
		if f.trees[helpers.OnedFromTwod(xx, y, f.width)] >= refTree {
			visCheck = false
		}
	}
	if visCheck {
		return true
	}

	return false
}

func (f *Forest) ScenicScore(x, y int) int64 {
	refTree := f.trees[helpers.OnedFromTwod(x, y, f.width)]

	// Look north
	var distN int64
	for yy := y - 1; yy >= 0; yy-- {
		distN = distN + 1
		if f.trees[helpers.OnedFromTwod(x, yy, f.width)] >= refTree {
			break
		}
	}

	// Look south
	var distS int64
	for yy := y + 1; yy < f.height; yy++ {
		distS = distS + 1
		if f.trees[helpers.OnedFromTwod(x, yy, f.width)] >= refTree {
			break
		}
	}

	// Look west
	var distW int64
	for xx := x - 1; xx >= 0; xx-- {
		distW = distW + 1
		if f.trees[helpers.OnedFromTwod(xx, y, f.width)] >= refTree {
			break
		}
	}

	// Look east
	var distE int64
	for xx := x + 1; xx < f.width; xx++ {
		distE = distE + 1
		if f.trees[helpers.OnedFromTwod(xx, y, f.width)] >= refTree {
			break
		}
	}

	return distN * distS * distE * distW
}

func (f *Forest) BestScenicScore() int64 {
	var bestScore int64
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			score := f.ScenicScore(x, y)
			if score > bestScore {
				bestScore = score
			}
		}
	}

	return bestScore
}

func (f *Forest) CountVisible() int {
	var cnt int
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height; y++ {
			if f.IsVisible(x, y) {
				cnt = cnt + 1
			}
		}
	}

	return cnt
}

func MakeForest(input []string) Forest {
	h := len(input)
	w := len(input[0])
	var trees []int

	for _, row := range input {
		for _, c := range row {
			trees = append(trees, int(c)-48)
		}
	}

	return Forest{
		width:  w,
		height: h,
		trees:  trees,
	}
}
