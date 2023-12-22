class Graph {
    constructor() {
      this.adjList = new Map();
    }
  
    addVertex(vertex) {
      this.adjList.set(vertex, []);
    }
  
    addEdge(vertex1, vertex2) {
      this.adjList.get(vertex1).push(vertex2);
      this.adjList.get(vertex2).push(vertex1); // for an undirected graph
    }
  
    dfs(startingVertex) {
      const visited = new Set();
  
      const dfsHelper = (vertex) => {
        visited.add(vertex);
        console.log(vertex);
  
        for (const neighbor of this.adjList.get(vertex)) {
          if (!visited.has(neighbor)) {
            dfsHelper(neighbor);
          }
        }
      };
  
      dfsHelper(startingVertex);
    }
  }
  
  // Example usage:
  const myGraph = new Graph();
  myGraph.addVertex('A');
  myGraph.addVertex('B');
  myGraph.addVertex('C');
  myGraph.addVertex('D');
  myGraph.addVertex('E');
  
  myGraph.addEdge('A', 'B');
  myGraph.addEdge('B', 'C');
  myGraph.addEdge('B', 'D');
  myGraph.addEdge('D', 'E');
  
  console.log('DFS Traversal:');
  myGraph.dfs('A');
  