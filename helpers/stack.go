package helpers

import "strings"

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
