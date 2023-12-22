use std::collections::{HashSet, HashMap, VecDeque};

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

    fn bfs(&self, starting_vertex: char) {
        let mut visited = HashSet::new();
        let mut queue = VecDeque::new();

        queue.push_back(starting_vertex);
        visited.insert(starting_vertex);

        while let Some(current_vertex) = queue.pop_front() {
            println!("{}", current_vertex);

            if let Some(neighbors) = self.adj_list.get(&current_vertex) {
                for &neighbor in neighbors {
                    if !visited.contains(&neighbor) {
                        queue.push_back(neighbor);
                        visited.insert(neighbor);
                    }
                }
            }
        }
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

    println!("BFS Traversal:");
    my_graph.bfs('A');
}
