package day05

import (
	"fmt"
	"strconv"
	"strings"
)

type Storage struct {
	stacks   []Stack
	moveList []Movement
}

func (s *Storage) String() string {
	retval := ""
	for i, s := range s.stacks {
		retval = fmt.Sprintf("%s%d: %s\n", retval, i+1, s.String())
	}

	return retval
}

func (s *Storage) DoWork(singles bool) string {
	for _, movement := range s.moveList {
		if singles {
			for cnt := 0; cnt < movement.Count; cnt++ {
				s.stacks[movement.Destination-1].Push(s.stacks[movement.Origin-1].Pop(1))
			}
		} else {
			s.stacks[movement.Destination-1].Push(s.stacks[movement.Origin-1].Pop(movement.Count))
		}
	}

	var retval string
	for _, stack := range s.stacks {
		retval = fmt.Sprintf("%s%s", retval, stack.Peek())
	}

	return retval
}

func MakeStorage(input []string) Storage {
	var stacks []Stack
	var movements []Movement
	for _, line := range input {
		if line == "" {
			continue
		}
		line = strings.ReplaceAll(strings.ReplaceAll(line, "[", " "), "]", " ")
		if line[0] == ' ' && line[1] != '1' {
			// This is a line defining stack entries
			if len(stacks) == 0 {
				stacks = make([]Stack, (len(line)+1)/4)
			}
			for i := 0; i < len(stacks); i++ {
				if line[i*4+1] != ' ' {
					stacks[i].Push([]string{string(line[i*4+1])})
				}
			}
		}

		if strings.HasPrefix(line, "move") {
			// This is a line defining a movement
			movements = append(movements, MakeMovement(line))
		}
	}

	// Reverse all stacks
	var revStacks []Stack
	for _, s := range stacks {
		s.Reverse()
		revStacks = append(revStacks, s)
	}

	return Storage{
		stacks:   revStacks,
		moveList: movements,
	}
}

type Movement struct {
	Origin      int
	Destination int
	Count       int
}

func MakeMovement(input string) Movement {
	input = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(input, "move ", ""), "from ", ""), "to ", "")
	elements := strings.Split(input, " ")
	count, _ := strconv.ParseInt(elements[0], 10, 64)
	origin, _ := strconv.ParseInt(elements[1], 10, 64)
	destination, _ := strconv.ParseInt(elements[2], 10, 64)

	return Movement{
		Origin:      int(origin),
		Destination: int(destination),
		Count:       int(count),
	}
}

type Stack struct {
	stack []string
}

func (s *Stack) Pop(cnt int) []string {
	retval := s.stack[len(s.stack)-cnt:]
	s.stack = s.stack[:len(s.stack)-cnt]
	return retval
}

func (s *Stack) Peek() string {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Push(items []string) {
	s.stack = append(s.stack, items...)
}

func (s *Stack) Reverse() {
	var retval []string
	for i := len(s.stack) - 1; i >= 0; i-- {
		retval = append(retval, s.stack[i])
	}

	s.stack = retval
}

func (s *Stack) String() string {
	return strings.Join(s.stack, ",")
}
