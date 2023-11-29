package day17

import (
	"fmt"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type Direction int

const (
	left Direction = iota
	right
	down
)

type Rock struct {
	shape  []string
	coords []helpers.Coords
}

func (r *Rock) Height() int {
	return len(r.shape)
}

func NewRock(input []string) Rock {
	var rock Rock
	rock.shape = append(rock.shape, input...)

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if input[y][x] == '#' {
				rock.coords = append(rock.coords, helpers.NewCoords(x, y))
			}
		}
	}

	return rock
}

type Chamber struct {
	linesToStore int
	width        int
	settled      []string
	archived     int64
	rockSequence []Rock
	windSequence []rune

	fallingRock *Rock
	fallingPos  helpers.Coords
	rockIndex   int
	windIndex   int
}

func (c *Chamber) String(pretty bool) string {
	if len(c.settled) == 0 {
		return ""
	}

	var retval strings.Builder
	for row := len(c.settled) - 1; row >= 0; row-- {
		if pretty {
			if row == len(c.settled)-1 {
				fmt.Fprintf(&retval, " %13d → ", c.Height())
			} else {
				fmt.Fprintf(&retval, "                 ")
			}
		}
		fmt.Fprintf(&retval, "│%s│", c.settled[row])
		if pretty {
			fmt.Fprintln(&retval)
		}
	}
	if pretty {
		if c.archived > 0 {
			fmt.Fprintf(&retval, " %13d ╮ │◌◌◌◌◌◌◌│\n", c.archived)
			fmt.Fprintf(&retval, "     more rows ┝ │◌◌◌◌◌◌◌│\n")
			fmt.Fprintf(&retval, "     not shown ╯ │◌◌◌◌◌◌◌│\n")
		}
		fmt.Fprintf(&retval, "                 ╘═══════╛\n")
	}

	return retval.String()
}

func (c *Chamber) CanMove(dir Direction) bool {
	var hmover, vmover int
	switch dir {
	case left:
		hmover = -1
	case right:
		hmover = 1
	case down:
		vmover = -1
	}

	for _, coords := range c.fallingRock.coords {
		testX := c.fallingPos.X + coords.X + hmover
		testY := c.fallingPos.Y - coords.Y + vmover // negative because rock map y-axis is inverted from chamber map

		// Check for collision with chamber sides
		if testX < 0 || testX >= c.width {
			return false
		}
		if testY < 0 {
			return false
		}

		// Check for collision with settled rocks
		if len(c.settled) > testY {
			if c.settled[testY][testX] == '#' {
				return false
			}
		}
	}

	return true
}

func (c *Chamber) IntroduceRock() {
	c.fallingRock = &c.rockSequence[c.rockIndex]
	c.rockIndex = (c.rockIndex + 1) % len(c.rockSequence)
	c.fallingPos.X = 2
	c.fallingPos.Y = len(c.settled) + 2 + c.fallingRock.Height()
}

func (c *Chamber) SettleRock() {
	// Add lines to chamber for settled rock if needed
	var newLines []string
	for i := len(c.settled); i <= c.fallingPos.Y; i++ {
		newLines = append(newLines, ".......")
	}
	c.settled = append(c.settled, newLines...)

	// Add rock to settled lines
	for _, coords := range c.fallingRock.coords {
		row := c.fallingPos.Y - coords.Y
		col := c.fallingPos.X + coords.X
		c.settled[row] = helpers.SetCharInString(c.settled[row], '#', col)
	}
}

func (c *Chamber) DropRock() {
	c.IntroduceRock()

	var settled bool
	for !settled {
		// Apply wind
		var windDir Direction
		if c.windSequence[c.windIndex] == '<' {
			windDir = left
		} else {
			windDir = right
		}
		c.windIndex = (c.windIndex + 1) % len(c.windSequence)

		if c.CanMove(windDir) {
			if windDir == left {
				c.fallingPos.X = c.fallingPos.X - 1
			} else {
				c.fallingPos.X = c.fallingPos.X + 1
			}
		}

		// Apply gravity
		if c.CanMove(down) {
			c.fallingPos.Y = c.fallingPos.Y - 1
		} else {
			settled = true
			c.SettleRock()
		}
	}

	c.trimTop()
	c.trimBottom()
}

func (c *Chamber) trimTop() {
	for c.settled[len(c.settled)-1] == "......." {
		c.settled = c.settled[:len(c.settled)-1]
	}
}

func (c *Chamber) trimBottom() {
	truncate := helpers.Max(len(c.settled)-c.linesToStore, 0)
	c.settled = c.settled[truncate:]
	c.archived += int64(truncate)
}

func (c *Chamber) Height() int64 {
	return int64(len(c.settled)) + c.archived
}

type cacheElem struct {
	height     int64
	iterations int64
}

func (c *Chamber) Simulate(iterations int64) {
	cache := make(map[string]cacheElem)
	for iterations > 0 {
		hash := c.String(false)
		if elem, ok := cache[hash]; ok {
			repeatingHeight := c.Height() - elem.height
			repeatingLength := elem.iterations - iterations
			c.archived = c.archived + (iterations/repeatingLength)*repeatingHeight
			iterations = iterations % repeatingLength
			break
		} else {
			cache[hash] = cacheElem{height: c.Height(), iterations: iterations}
			c.DropRock()
			iterations--
		}
	}

	for iterations > 0 {
		c.DropRock()
		iterations--
	}
}

func NewChamber(input string) Chamber {
	chamber := Chamber{width: 7, linesToStore: 50}

	chamber.rockSequence = append(chamber.rockSequence, NewRock([]string{
		"####",
	}))
	chamber.rockSequence = append(chamber.rockSequence, NewRock([]string{
		".#.",
		"###",
		".#.",
	}))
	chamber.rockSequence = append(chamber.rockSequence, NewRock([]string{
		"..#",
		"..#",
		"###",
	}))
	chamber.rockSequence = append(chamber.rockSequence, NewRock([]string{
		"#",
		"#",
		"#",
		"#",
	}))
	chamber.rockSequence = append(chamber.rockSequence, NewRock([]string{
		"##",
		"##",
	}))

	for _, c := range input {
		chamber.windSequence = append(chamber.windSequence, c)
	}

	return chamber
}
