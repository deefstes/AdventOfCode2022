package day16

import (
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	Name  string
	Rate  uint16
	Edges string
}

type Cave struct {
	valves []*Valve
	start  uint16
	nodes  [][2]uint16
	edges  []uint16
}

func NewGraph(valves []*Valve) map[*Valve]map[*Valve]uint16 {
	graph := make(map[*Valve]map[*Valve]uint16)
	for _, v1 := range valves {
		graph[v1] = make(map[*Valve]uint16)
		for _, v2 := range valves {
			if v1 == v2 {
				graph[v1][v2] = 0
			} else if strings.Contains(v1.Edges, v2.Name) {
				graph[v1][v2] = 1
			} else {
				graph[v1][v2] = 0xff
			}
		}
	}

	for _, i1 := range valves {
		for _, i2 := range valves {
			for _, i3 := range valves {
				if graph[i2][i3] > graph[i2][i1]+graph[i1][i3] {
					graph[i2][i3] = graph[i2][i1] + graph[i1][i3]
				}
			}
		}
	}

	return graph
}

func NewCave(lines []string) Cave {
	var cave Cave

	regex := regexp.MustCompile(`Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? ((?:[A-Z]{2}(?:, )?)+)`)

	for _, line := range lines {
		comps := regex.FindStringSubmatch(line)
		i, _ := strconv.Atoi(comps[2])
		cave.valves = append(cave.valves, &Valve{Name: comps[1], Rate: uint16(i), Edges: comps[3]})
	}

	graph := NewGraph(cave.valves)

	var candidates []*Valve
	for _, v := range cave.valves {
		if v.Rate > 0 || v.Name == "AA" {
			candidates = append(candidates, v)
		}
	}

	bitfield := make(map[*Valve]uint16)
	for i, v := range candidates {
		bitfield[v] = 1 << i
	}

	for _, v := range candidates {
		if v.Name == "AA" {
			cave.start = bitfield[v]
			break
		}
	}

	cave.edges = make([]uint16, 0xffff)
	for _, v1 := range candidates {
		for _, v2 := range candidates {
			cave.edges[bitfield[v1]|bitfield[v2]] = graph[v1][v2]
		}
	}

	cave.nodes = make([][2]uint16, len(candidates))
	for idx, v := range candidates {
		cave.nodes[idx] = [2]uint16{bitfield[v], v.Rate}
	}

	return cave
}

func (c *Cave) Search(target, pressure, minute, on, node uint16) uint16 {
	max := pressure
	for _, w := range c.nodes {
		if node == w[0] || w[0] == c.start || w[0]&on != 0 {
			continue
		}
		l := c.edges[node|w[0]] + 1
		if minute+l > target {
			continue
		}
		if next := c.Search(target, pressure+(target-minute-l)*w[1], minute+l, on|w[0], w[0]); next > max {
			max = next
		}
	}
	return max
}

func (c *Cave) Paths(target, pressure, minute, on, node, path uint16) [][2]uint16 {
	paths := [][2]uint16{{pressure, path}}
	for _, w := range c.nodes {
		if w[0] == node || w[0] == c.start || w[0]&on != 0 {
			continue
		}
		l := c.edges[node|w[0]] + 1
		if minute+l > target {
			continue
		}
		paths = append(paths, c.Paths(target, pressure+(target-minute-l)*w[1], minute+l, on|w[0], w[0], path|w[0])...)
	}
	return paths
}

func (c *Cave) Solve(withElephants bool) int {
	solitary := c.Search(30, 0, 0, 0, c.start)
	if !withElephants {
		return int(solitary)
	}

	var paths [][2]uint16
	for _, p := range c.Paths(26, 0, 0, 0, c.start, 0) {
		if p[0] > solitary/2 {
			paths = append(paths, p)
		}
	}

	var max uint16 = 0
	for i1 := 0; i1 < len(paths); i1 += 1 {
		for i2 := i1 + 1; i2 < len(paths); i2 += 1 {
			if paths[i1][1]&paths[i2][1] != 0 {
				continue
			}
			if m := paths[i1][0] + paths[i2][0]; m > max {
				max = m
			}
		}
	}

	return int(max)
}
