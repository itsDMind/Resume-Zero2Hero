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

func (g *Graph) BFS(startingVertex string) {
	visited := make(map[string]bool)
	queue := make([]string, 0)

	queue = append(queue, startingVertex)
	visited[startingVertex] = true

	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]
		fmt.Println(currentVertex)

		for _, neighbor := range g.adjList[currentVertex] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
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

	fmt.Println("BFS Traversal:")
	myGraph.BFS("A")
}
