import heapq

class Graph:
    def __init__(self):
        self.nodes = {}

    def add_node(self, node):
        self.nodes[node] = []

    def add_edge(self, node1, node2, weight):
        self.nodes[node1].append((node2, weight))
        self.nodes[node2].append((node1, weight))

    def dijkstra(self, start_node):
        distances = {node: float('inf') for node in self.nodes}
        distances[start_node] = 0
        priority_queue = [(0, start_node)]

        while priority_queue:
            current_distance, current_node = heapq.heappop(priority_queue)

            for neighbor, weight in self.nodes[current_node]:
                distance = current_distance + weight
                if distance < distances[neighbor]:
                    distances[neighbor] = distance
                    heapq.heappush(priority_queue, (distance, neighbor))

        return distances

# Example usage:
my_graph = Graph()
my_graph.add_node("A")
my_graph.add_node("B")
my_graph.add_node("C")
my_graph.add_edge("A", "B", 1)
my_graph.add_edge("B", "C", 2)
my_graph.add_edge("A", "C", 4)

distances = my_graph.dijkstra("A")
print(distances)  # Output: {'A': 0, 'B': 1, 'C': 3}
