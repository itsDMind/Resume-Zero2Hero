use std::collections::{BinaryHeap, HashMap};
use std::cmp::Ordering;

#[derive(Eq, PartialEq)]
struct Node {
    name: String,
    distance: u32,
}

impl Ord for Node {
    fn cmp(&self, other: &Self) -> Ordering {
        other.distance.cmp(&self.distance)
    }
}

impl PartialOrd for Node {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

struct Graph {
    nodes: HashMap<String, Vec<(String, u32)>>,
}

impl Graph {
    fn new() -> Self {
        Graph {
            nodes: HashMap::new(),
        }
    }

    fn add_node(&mut self, node: &str) {
        self.nodes.insert(node.to_string(), Vec::new());
    }

    fn add_edge(&mut self, node1: &str, node2: &str, weight: u32) {
        self.nodes.get_mut(node1).unwrap().push((node2.to_string(), weight));
        self.nodes.get_mut(node2).unwrap().push((node1.to_string(), weight));
    }

    fn dijkstra(&self, start_node: &str) -> HashMap<String, u32> {
        let mut distances: HashMap<String, u32> = self.nodes.keys().map(|k| (k.clone(), u32::MAX)).collect();
        let mut priority_queue = BinaryHeap::new();

        distances.insert(start_node.to_string(), 0);
        priority_queue.push(Node { name: start_node.to_string(), distance: 0 });

        while let Some(Node { name, distance }) = priority_queue.pop() {
            for (neighbor, weight) in &self.nodes[&name] {
                let new_distance = distance + weight;
                if new_distance < distances[neighbor] {
                    distances.insert(neighbor.clone(), new_distance);
                    priority_queue.push(Node { name: neighbor.clone(), distance: new_distance });
                }
            }
        }

        distances
    }
}

fn main() {
    let mut my_graph = Graph::new();
    my_graph.add_node("A");
    my_graph.add_node("B");
    my_graph.add_node("C");
    my_graph.add_edge("A", "B", 1);
    my_graph.add_edge("B", "C", 2);
    my_graph.add_edge("A", "C", 4);

    let distances = my_graph.dijkstra("A");
    println!("{:?}", distances); // Output: {"A": 0, "B": 1, "C": 3}
}
