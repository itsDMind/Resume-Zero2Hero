from collections import defaultdict

class Graph:
    def __init__(self):
        self.adj_list = defaultdict(list)

    def add_vertex(self, vertex):
        self.adj_list[vertex]

    def add_edge(self, vertex1, vertex2):
        self.adj_list[vertex1].append(vertex2)
        self.adj_list[vertex2].append(vertex1)  # for an undirected graph

    def dfs(self, starting_vertex):
        visited = set()

        def dfs_helper(vertex):
            visited.add(vertex)
            print(vertex)

            for neighbor in self.adj_list[vertex]:
                if neighbor not in visited:
                    dfs_helper(neighbor)

        dfs_helper(starting_vertex)

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

print('DFS Traversal:')
my_graph.dfs('A')
