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
  
    bfs(startingVertex) {
      const visited = new Set();
      const queue = [];
  
      queue.push(startingVertex);
      visited.add(startingVertex);
  
      while (queue.length !== 0) {
        const currentVertex = queue.shift();
        console.log(currentVertex);
  
        for (const neighbor of this.adjList.get(currentVertex)) {
          if (!visited.has(neighbor)) {
            queue.push(neighbor);
            visited.add(neighbor);
          }
        }
      }
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
  
  console.log('BFS Traversal:');
  myGraph.bfs('A');
  