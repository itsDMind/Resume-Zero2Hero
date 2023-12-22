use std::collections::{HashMap, HashSet, VecDeque};

struct Graph {
    adjacency_list: HashMap<String, Vec<String>>,
}

impl Graph {
    fn new() -> Self {
        Graph {
            adjacency_list: HashMap::new(),
        }
    }

    fn add_vertex(&mut self, vertex: &str) {
        self.adjacency_list.insert(vertex.to_string(), Vec::new());
    }

    fn add_edge(&mut self, vertex1: &str, vertex2: &str) {
        self.adjacency_list
            .get_mut(vertex1)
            .unwrap()
            .push(vertex2.to_string());
        self.adjacency_list
            .get_mut(vertex2)
            .unwrap()
            .push(vertex1.to_string());
    }

    fn dfs(&self, start_vertex: &str, visited: &mut HashSet<String>) {
        println!("{}", start_vertex);
        visited.insert(start_vertex.to_string());

        for neighbor in self.adjacency_list.get(start_vertex).unwrap() {
            if !visited.contains(neighbor) {
                self.dfs(neighbor, visited);
            }
        }
    }

    fn bfs(&self, start_vertex: &str) {
        let mut queue = VecDeque::new();
        let mut visited = HashSet::new();

        queue.push_back(start_vertex.to_string());

        while let Some(current_vertex) = queue.pop_front() {
            println!("{}", current_vertex);
            visited.insert(current_vertex.clone());

            for neighbor in self.adjacency_list.get(&current_vertex).unwrap() {
                if !visited.contains(neighbor) && !queue.contains(neighbor) {
                    queue.push_back(neighbor.to_string());
                }
            }
        }
    }
}

fn main() {
    let mut my_graph = Graph::new();
    my_graph.add_vertex("A");
    my_graph.add_vertex("B");
    my_graph.add_vertex("C");
    my_graph.add_edge("A", "B");
    my_graph.add_edge("B", "C");

    println!("DFS:");
    my_graph.dfs("A", &mut HashSet::new());

    println!("BFS:");
    my_graph.bfs("A");
}
