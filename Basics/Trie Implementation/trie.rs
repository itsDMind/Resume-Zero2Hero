use std::collections::HashMap;

struct TrieNode {
    children: HashMap<char, TrieNode>,
    is_end_of_word: bool,
}

impl TrieNode {
    fn new() -> Self {
        TrieNode {
            children: HashMap::new(),
            is_end_of_word: false,
        }
    }
}

struct Trie {
    root: TrieNode,
}

impl Trie {
    fn new() -> Self {
        Trie { root: TrieNode::new() }
    }

    fn insert(&mut self, word: &str) {
        let mut node = &mut self.root;
        for char in word.chars() {
            node = node.children.entry(char).or_insert(TrieNode::new());
        }
        node.is_end_of_word = true;
    }

    fn search(&self, word: &str) -> bool {
        let mut node = &self.root;
        for char in word.chars() {
            if let Some(next_node) = node.children.get(&char) {
                node = next_node;
            } else {
                return false;
            }
        }
        node.is_end_of_word
    }

    fn starts_with(&self, prefix: &str) -> bool {
        let mut node = &self.root;
        for char in prefix.chars() {
            if let Some(next_node) = node.children.get(&char) {
                node = next_node;
            } else {
                return false;
            }
        }
        true
    }
}

// Example usage:
fn main() {
    let mut trie = Trie::new();
    trie.insert("apple");
    println!("{}", trie.search("apple"));  // Output: true
    println!("{}", trie.search("app"));    // Output: false
    println!("{}", trie.starts_with("app"));  // Output: true
    trie.insert("app");
    println!("{}", trie.search("app"));    // Output: true
}
