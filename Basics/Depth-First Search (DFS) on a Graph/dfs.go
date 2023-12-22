package main

import "fmt"

type Graph struct {
	adjList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[string][]string),
	}
}

func (g *Graph) AddVertex(vertex string) {
	g.adjList[vertex] = []string{}
}

func (g *Graph) AddEdge(vertex1, vertex2 string) {
	g.adjList[vertex1] = append(g.adjList[vertex1], vertex2)
	g.adjList[vertex2] = append(g.adjList[vertex2], vertex1) // for an undirected graph
}

func (g *Graph) DFS(startingVertex string) {
	visited := make(map[string]bool)

	var dfsHelper func(vertex string)
	dfsHelper = func(vertex string) {
		visited[vertex] = true
		fmt.Println(vertex)

		for _, neighbor := range g.adjList[vertex] {
			if !visited[neighbor] {
				dfsHelper(neighbor)
			}
		}
	}

	dfsHelper(startingVertex)
}

// Example usage:
func main() {
	myGraph := NewGraph()
	myGraph.AddVertex("A")
	myGraph.AddVertex("B")
	myGraph.AddVertex("C")
	myGraph.AddVertex("D")
	myGraph.AddVertex("E")

	myGraph.AddEdge("A", "B")
	myGraph.AddEdge("B", "C")
	myGraph.AddEdge("B", "D")
	myGraph.AddEdge("D", "E")

	fmt.Println("DFS Traversal:")
	myGraph.DFS("A")
}
