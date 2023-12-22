use std::collections::{HashSet, HashMap};

struct Graph {
    adj_list: HashMap<char, Vec<char>>,
}

impl Graph {
    fn new() -> Self {
        Graph {
            adj_list: HashMap::new(),
        }
    }

    fn add_vertex(&mut self, vertex: char) {
        self.adj_list.entry(vertex).or_insert(Vec::new());
    }

    fn add_edge(&mut self, vertex1: char, vertex2: char) {
        self.adj_list.entry(vertex1).or_insert(Vec::new()).push(vertex2);
        self.adj_list.entry(vertex2).or_insert(Vec::new()).push(vertex1); // for an undirected graph
    }

    fn dfs(&self, starting_vertex: char) {
        let mut visited = HashSet::new();

        fn dfs_helper(graph: &Graph, vertex: char, visited: &mut HashSet<char>) {
            visited.insert(vertex);
            println!("{}", vertex);

            if let Some(neighbors) = graph.adj_list.get(&vertex) {
                for &neighbor in neighbors {
                    if !visited.contains(&neighbor) {
                        dfs_helper(graph, neighbor, visited);
                    }
                }
            }
        }

        dfs_helper(self, starting_vertex, &mut visited);
    }
}

// Example usage:
fn main() {
    let mut my_graph = Graph::new();
    my_graph.add_vertex('A');
    my_graph.add_vertex('B');
    my_graph.add_vertex('C');
    my_graph.add_vertex('D');
    my_graph.add_vertex('E');

    my_graph.add_edge('A', 'B');
    my_graph.add_edge('B', 'C');
    my_graph.add_edge('B', 'D');
    my_graph.add_edge('D', 'E');

    println!("DFS Traversal:");
    my_graph.dfs('A');
}
