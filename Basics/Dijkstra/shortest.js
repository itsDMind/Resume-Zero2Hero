class Graph {
    constructor() {
      this.nodes = {};
    }
  
    addNode(node) {
      this.nodes[node] = [];
    }
  
    addEdge(node1, node2, weight) {
      this.nodes[node1].push({ node: node2, weight });
      this.nodes[node2].push({ node: node1, weight });
    }
  
    dijkstra(startNode) {
      const distances = {};
      const visited = {};
      const priorityQueue = new PriorityQueue();
  
      // Initialize distances and priority queue
      for (const node in this.nodes) {
        distances[node] = node === startNode ? 0 : Infinity;
        priorityQueue.enqueue(node, distances[node]);
      }
  
      while (!priorityQueue.isEmpty()) {
        const currentNode = priorityQueue.dequeue();
        visited[currentNode] = true;
  
        for (const neighbor of this.nodes[currentNode]) {
          if (!visited[neighbor.node]) {
            const potentialDistance = distances[currentNode] + neighbor.weight;
            if (potentialDistance < distances[neighbor.node]) {
              distances[neighbor.node] = potentialDistance;
              priorityQueue.enqueue(neighbor.node, potentialDistance);
            }
          }
        }
      }
  
      return distances;
    }
  }
  
  class PriorityQueue {
    constructor() {
      this.queue = [];
    }
  
    enqueue(node, priority) {
      this.queue.push({ node, priority });
      this.sort();
    }
  
    dequeue() {
      return this.queue.shift().node;
    }
  
    isEmpty() {
      return this.queue.length === 0;
    }
  
    sort() {
      this.queue.sort((a, b) => a.priority - b.priority);
    }
  }
  
  // Example usage:
  const myGraph = new Graph();
  myGraph.addNode("A");
  myGraph.addNode("B");
  myGraph.addNode("C");
  myGraph.addEdge("A", "B", 1);
  myGraph.addEdge("B", "C", 2);
  myGraph.addEdge("A", "C", 4);
  
  const distances = myGraph.dijkstra("A");
  console.log(distances); // Output: { A: 0, B: 1, C: 3 }
  