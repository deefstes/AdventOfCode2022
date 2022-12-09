package helpers

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadInputFile() []string {
	var inputFileName string
	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.Parse()
	fmt.Printf("input file name: %s\n", inputFileName)

	// Read input file
	file, err := os.ReadFile(inputFileName)
	if err != nil {
		log.Fatalf("reading input file: %v", err)
	}
	lines := strings.Split(strings.TrimRight(string(file), "\n"), "\n")

	return lines
}

func IsNumber(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

func OnedFromTwod(x, y, width int) int {
	return y*width + x
}

type Coords struct {
	X int
	Y int
}

func NewCoords(x, y int) Coords {
	return Coords{X: x, Y: y}
}

func (c *Coords) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func (c1 *Coords) SimpleDist(c2 Coords) int {
	dx := Abs(c1.X - c2.X)
	dy := Abs(c1.Y - c2.Y)
	d := math.Sqrt(float64(dx*dx + dy*dy))

	return int(d)
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
