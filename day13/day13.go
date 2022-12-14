package day13

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/deefstes/AdventOfCode2022/helpers"
)

type PacketType int

const (
	value PacketType = iota
	list
	empty
)

type Packet struct {
	pType PacketType
	value int
	list  []Packet
}

func (p *Packet) String() string {
	var ret strings.Builder
	switch p.pType {
	case value:
		fmt.Fprintf(&ret, "%d", p.value)
	case list:
		fmt.Fprintf(&ret, "[")
		for _, v := range p.list {
			fmt.Fprintf(&ret, "%s,", v.String())
		}
		fmt.Fprintf(&ret, "]")
	}

	return strings.ReplaceAll(ret.String(), ",]", "]")
}

func NewPacket(in string) Packet {
	var ret Packet

	if in == "" {
		ret.pType = empty
		return ret
	}

	if in[0] == '[' {
		ret.pType = list
		// Parse elements
		elemStart := 1
		for pos := elemStart; pos < len(in); pos++ {
			switch in[pos] {
			case ',', ']':
				newP := NewPacket(in[elemStart:pos])
				if newP.pType != empty {
					ret.list = append(ret.list, newP)
				}
				elemStart = pos + 1
			case '[':
				var closingIndex int
				var bracketStack helpers.Stack[rune]
				bracketStack.Push('[')
				for i := pos + 1; bracketStack.Size() != 0; i++ {
					switch in[i] {
					case '[':
						bracketStack.Push('[')
					case ']':
						bracketStack.Pop(1)
					}

					if bracketStack.Size() == 0 {
						closingIndex = i + 1
					}
				}
				ret.list = append(ret.list, NewPacket(in[pos:closingIndex]))
				pos = closingIndex
				elemStart = pos + 1
			}
		}
	} else {
		ret.pType = value
		if in != "" {
			ret.value, _ = strconv.Atoi(in)
		}
	}

	return ret
}

type PacketPair struct {
	left  Packet
	right Packet
}

func (p *PacketPair) String() string {
	var ret strings.Builder
	fmt.Fprintf(&ret, "%s\n", p.left.String())
	fmt.Fprintf(&ret, "%s\n", p.right.String())

	return ret.String()
}

func Compare(a, b Packet) int {
	if a.pType == value && b.pType == value {
		if a.value > b.value {
			return -1
		} else if b.value > a.value {
			return 1
		}
		return 0
	}

	if a.pType == value {
		a.list = []Packet{a}
	}
	if b.pType == value {
		b.list = []Packet{b}
	}
	max := helpers.Max(len(a.list), len(b.list))
	for i := 0; i < max; i++ {
		if i >= len(a.list) {
			return 1
		}
		if i >= len(b.list) {
			return -1
		}
		if sub := Compare(a.list[i], b.list[i]); sub != 0 {
			return sub
		}
	}
	return 0
}

type Signal struct {
	pairs []PacketPair
}

func (s *Signal) RightOrderedPairs() ([]int, []PacketPair) {
	var orderedPairs []PacketPair
	var orderedIndices []int

	for i, p := range s.pairs {
		sub := Compare(p.left, p.right)
		if sub >= 0 {
			orderedPairs = append(orderedPairs, p)
			orderedIndices = append(orderedIndices, i+1)
		}
	}

	return orderedIndices, orderedPairs
}

func (s *Signal) String() string {
	var ret strings.Builder

	for _, p := range s.pairs {
		fmt.Fprintf(&ret, "%s\n", p.String())
	}

	return ret.String()
}

func NewSignal(lines []string) Signal {
	var signal Signal

	for i := 0; i < len(lines); i = i + 3 {
		signal.pairs = append(signal.pairs, PacketPair{
			left:  NewPacket(lines[i]),
			right: NewPacket(lines[i+1]),
		})
	}
	return signal
}

func NewRawPackets(lines []string) []Packet {
	var retval []Packet

	for _, line := range lines {
		if line == "" {
			continue
		}
		retval = append(retval, NewPacket(line))
	}

	return retval
}
