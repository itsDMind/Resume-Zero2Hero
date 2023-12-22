use std::fmt;

#[derive(Debug)]
struct TreeNode {
    value: i32,
    left: Option<Box<TreeNode>>,
    right: Option<Box<TreeNode>>,
}

impl TreeNode {
    fn new(value: i32) -> Self {
        TreeNode {
            value,
            left: None,
            right: None,
        }
    }
}

#[derive(Debug)]
struct BinaryTree {
    root: Option<Box<TreeNode>>,
}

impl BinaryTree {
    fn new() -> Self {
        BinaryTree { root: None }
    }

    fn insert(&mut self, value: i32) {
        self.root = self._insert(self.root.take(), value);
    }

    fn _insert(&mut self, node: Option<Box<TreeNode>>, value: i32) -> Option<Box<TreeNode>> {
        match node {
            None => Some(Box::new(TreeNode::new(value))),
            Some(mut n) => {
                if value < n.value {
                    n.left = self._insert(n.left.take(), value);
                } else {
                    n.right = self._insert(n.right.take(), value);
                }
                Some(n)
            }
        }
    }

    fn in_order_traversal(&self) -> Vec<i32> {
        let mut result = Vec::new();
        self._in_order_traversal(self.root.as_ref(), &mut result);
        result
    }

    fn _in_order_traversal(&self, node: Option<&Box<TreeNode>>, result: &mut Vec<i32>) {
        if let Some(n) = node {
            self._in_order_traversal(n.left.as_ref(), result);
            result.push(n.value);
            self._in_order_traversal(n.right.as_ref(), result);
        }
    }

    fn pre_order_traversal(&self) -> Vec<i32> {
        let mut result = Vec::new();
        self._pre_order_traversal(self.root.as_ref(), &mut result);
        result
    }

    fn _pre_order_traversal(&self, node: Option<&Box<TreeNode>>, result: &mut Vec<i32>) {
        if let Some(n) = node {
            result.push(n.value);
            self._pre_order_traversal(n.left.as_ref(), result);
            self._pre_order_traversal(n.right.as_ref(), result);
        }
    }

    fn post_order_traversal(&self) -> Vec<i32> {
        let mut result = Vec::new();
        self._post_order_traversal(self.root.as_ref(), &mut result);
        result
    }

    fn _post_order_traversal(&self, node: Option<&Box<TreeNode>>, result: &mut Vec<i32>) {
        if let Some(n) = node {
            self._post_order_traversal(n.left.as_ref(), result);
            self._post_order_traversal(n.right.as_ref(), result);
            result.push(n.value);
        }
    }
}

// Example usage:
fn main() {
    let mut my_tree = BinaryTree::new();
    my_tree.insert(5);
    my_tree.insert(3);
    my_tree.insert(7);
    my_tree.insert(2);
    my_tree.insert(4);
    my_tree.insert(6);
    my_tree.insert(8);

    println!("In-order traversal: {:?}", my_tree.in_order_traversal());
    println!("Pre-order traversal: {:?}", my_tree.pre_order_traversal());
    println!("Post-order traversal: {:?}", my_tree.post_order_traversal());
}
