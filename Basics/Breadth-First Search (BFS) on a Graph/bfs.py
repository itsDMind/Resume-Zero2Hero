from collections import defaultdict, deque

class Graph:
    def __init__(self):
        self.adj_list = defaultdict(list)

    def add_vertex(self, vertex):
        self.adj_list[vertex]

    def add_edge(self, vertex1, vertex2):
        self.adj_list[vertex1].append(vertex2)
        self.adj_list[vertex2].append(vertex1)  # for an undirected graph

    def bfs(self, starting_vertex):
        visited = set()
        queue = deque()

        queue.append(starting_vertex)
        visited.add(starting_vertex)

        while queue:
            current_vertex = queue.popleft()
            print(current_vertex)

            for neighbor in self.adj_list[current_vertex]:
                if neighbor not in visited:
                    queue.append(neighbor)
                    visited.add(neighbor)

# Example usage:
my_graph = Graph()
my_graph.add_vertex('A')
my_graph.add_vertex('B')
my_graph.add_vertex('C')
my_graph.add_vertex('D')
my_graph.add_vertex('E')

my_graph.add_edge('A', 'B')
my_graph.add_edge('B', 'C')
my_graph.add_edge('B', 'D')
my_graph.add_edge('D', 'E')

print('BFS Traversal:')
my_graph.bfs('A')
