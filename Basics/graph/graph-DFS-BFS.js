class Graph {
    constructor() {
      this.adjacencyList = new Map();
    }
  
    addVertex(vertex) {
      this.adjacencyList.set(vertex, []);
    }
  
    addEdge(vertex1, vertex2) {
      this.adjacencyList.get(vertex1).push(vertex2);
      this.adjacencyList.get(vertex2).push(vertex1);
    }
  
    dfs(startVertex, visited = new Set()) {
      console.log(startVertex);
      visited.add(startVertex);
  
      for (const neighbor of this.adjacencyList.get(startVertex)) {
        if (!visited.has(neighbor)) {
          this.dfs(neighbor, visited);
        }
      }
    }
  
    bfs(startVertex) {
      const queue = [startVertex];
      const visited = new Set();
  
      while (queue.length > 0) {
        const currentVertex = queue.shift();
        console.log(currentVertex);
        visited.add(currentVertex);
  
        for (const neighbor of this.adjacencyList.get(currentVertex)) {
          if (!visited.has(neighbor) && !queue.includes(neighbor)) {
            queue.push(neighbor);
          }
        }
      }
    }
  }
  
  // Example usage:
  const myGraph = new Graph();
  myGraph.addVertex("A");
  myGraph.addVertex("B");
  myGraph.addVertex("C");
  myGraph.addEdge("A", "B");
  myGraph.addEdge("B", "C");
  
  console.log("DFS:");
  myGraph.dfs("A");
  
  console.log("BFS:");
  myGraph.bfs("A");
  