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

func OnedFromXY(x, y, width int) int {
	return y*width + x
}

func OnedFromCoords(c Coords, width int) int {
	return c.Y*width + c.X
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

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Pop(cnt int) []T {
	retval := s.stack[len(s.stack)-cnt:]
	s.stack = s.stack[:len(s.stack)-cnt]
	return retval
}

func (s *Stack[T]) Peek() T {
	return s.stack[len(s.stack)-1]
}

func (s *Stack[T]) Push(items []T) {
	s.stack = append(s.stack, items...)
}

func (s *Stack[T]) Reverse() {
	var retval []T
	for i := len(s.stack) - 1; i >= 0; i-- {
		retval = append(retval, s.stack[i])
	}

	s.stack = retval
}

func (s *Stack[T]) Contains(val T, eq func(t1, t2 T) bool) bool {
	for _, v := range s.stack {
		if eq(v, val) {
			return true
		}
	}

	return false
}

func (s *Stack[T]) String(delim string, convFunc func(val T) string) string {
	var strStack []string
	for _, v := range s.stack {
		strStack = append(strStack, convFunc(v))
	}
	return strings.Join(strStack, delim)
}

type Queue[T any] []T

func (q *Queue[T]) Add(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Get() T {
	ret := (*q)[0]
	*q = (*q)[1:]

	return ret
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

func ReverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
