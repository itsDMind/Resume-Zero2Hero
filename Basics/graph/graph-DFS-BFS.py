from collections import defaultdict

class Graph:
    def __init__(self):
        self.adjacency_list = defaultdict(list)

    def add_vertex(self, vertex):
        self.adjacency_list[vertex] = []

    def add_edge(self, vertex1, vertex2):
        self.adjacency_list[vertex1].append(vertex2)
        self.adjacency_list[vertex2].append(vertex1)

    def dfs(self, start_vertex, visited=None):
        if visited is None:
            visited = set()
        print(start_vertex)
        visited.add(start_vertex)

        for neighbor in self.adjacency_list[start_vertex]:
            if neighbor not in visited:
                self.dfs(neighbor, visited)

    def bfs(self, start_vertex):
        queue = [start_vertex]
        visited = set()

        while queue:
            current_vertex = queue.pop(0)
            print(current_vertex)
            visited.add(current_vertex)

            for neighbor in self.adjacency_list[current_vertex]:
                if neighbor not in visited and neighbor not in queue:
                    queue.append(neighbor)

# Example usage:
my_graph = Graph()
my_graph.add_vertex("A")
my_graph.add_vertex("B")
my_graph.add_vertex("C")
my_graph.add_edge("A", "B")
my_graph.add_edge("B", "C")

print("DFS:")
my_graph.dfs("A")

print("BFS:")
my_graph.bfs("A")
