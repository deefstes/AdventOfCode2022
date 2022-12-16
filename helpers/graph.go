package helpers

import (
	"math"
)

type Node struct {
	Name    string
	Value   int
	Through *Node
}

func (n *Node) StepCount() int {
	if n.Through == nil {
		return 0
	}

	return 1 + n.Through.StepCount()
}

func (n *Node) TotalCost() int {
	if n.Through == nil {
		return n.Value
	}

	return n.Value + n.Through.TotalCost()
}

type Edge struct {
	Node   *Node
	Weight int
}

type WeightedGraph struct {
	Nodes []*Node
	Edges map[string][]*Edge
}

func NewGraph() *WeightedGraph {
	return &WeightedGraph{
		Edges: make(map[string][]*Edge),
	}
}

func (g *WeightedGraph) GetNode(name string) (node *Node) {
	for _, n := range g.Nodes {
		if n.Name == name {
			node = n
		}
	}
	return
}

func (g *WeightedGraph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func AddNodes(graph *WeightedGraph, names ...string) (nodes map[string]*Node) {
	nodes = make(map[string]*Node)
	for _, name := range names {
		n := &Node{name, math.MaxInt, nil}
		graph.AddNode(n)
		nodes[name] = n
	}
	return
}

func (g *WeightedGraph) AddEdge(n1, n2 *Node, weight int, bidirectional bool) {
	g.Edges[n1.Name] = append(g.Edges[n1.Name], &Edge{n2, weight})
	if bidirectional {
		g.Edges[n2.Name] = append(g.Edges[n2.Name], &Edge{n1, weight})
	}
}

func (g *WeightedGraph) Dijkstra(start string) {
	visited := make(map[string]bool)
	heap := &Heap{}

	startNode := g.GetNode(start)
	startNode.Value = 0
	heap.Push(startNode)

	for heap.Size() > 0 {
		current := heap.Pop()
		visited[current.Name] = true
		edges := g.Edges[current.Name]
		for _, edge := range edges {
			if !visited[edge.Node.Name] {
				heap.Push(edge.Node)
				if current.Value+edge.Weight < edge.Node.Value {
					edge.Node.Value = current.Value + edge.Weight
					edge.Node.Through = current
				}
			}
		}
	}
}
