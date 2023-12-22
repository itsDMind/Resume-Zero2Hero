package main

import (
	"fmt"
	"sync"
)

type Graph struct {
	adjacencyList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[string][]string)}
}

func (g *Graph) AddVertex(vertex string) {
	g.adjacencyList[vertex] = []string{}
}

func (g *Graph) AddEdge(vertex1, vertex2 string) {
	g.adjacencyList[vertex1] = append(g.adjacencyList[vertex1], vertex2)
	g.adjacencyList[vertex2] = append(g.adjacencyList[vertex2], vertex1)
}

func (g *Graph) DFS(startVertex string, visited map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(startVertex)
	visited[startVertex] = true

	for _, neighbor := range g.adjacencyList[startVertex] {
		if !visited[neighbor] {
			wg.Add(1)
			go g.DFS(neighbor, visited, wg)
		}
	}
}

func (g *Graph) BFS(startVertex string) {
	queue := []string{startVertex}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		currentVertex := queue[0]
		fmt.Println(currentVertex)
		visited[currentVertex] = true
		queue = queue[1:]

		for _, neighbor := range g.adjacencyList[currentVertex] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {
	myGraph := NewGraph()
	myGraph.AddVertex("A")
	myGraph.AddVertex("B")
	myGraph.AddVertex("C")
	myGraph.AddEdge("A", "B")
	myGraph.AddEdge("B", "C")

	fmt.Println("DFS:")
	var wg sync.WaitGroup
	wg.Add(1)
	go myGraph.DFS("A", make(map[string]bool), &wg)
	wg.Wait()

	fmt.Println("BFS:")
	myGraph.BFS("A")
}
