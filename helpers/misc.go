package helpers

import (
	"flag"
	"fmt"
	"log"
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

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func Abs64(val int64) int64 {
	if val < 0 {
		return -val
	}
	return val
}

func Max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
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

func SetCharInString(in string, r rune, i int) string {
	if i >= len(in) {
		return in
	}
	retval := []rune(in)
	retval[i] = r
	return string(retval)
}
