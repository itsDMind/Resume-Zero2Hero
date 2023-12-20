struct Node {
    data: i32,
    next: Option<Box<Node>>,
}

struct LinkedList {
    head: Option<Box<Node>>,
}

impl LinkedList {
    fn new() -> Self {
        LinkedList { head: None }
    }

    fn insert(&mut self, data: i32) {
        let new_node = Node {
            data,
            next: self.head.take(),
        };
        self.head = Some(Box::new(new_node));
    }

    fn delete(&mut self, key: i32) {
        let mut current = &mut self.head;

        while let Some(node) = current {
            if node.data == key {
                let next = node.next.take();
                *current = next;
                return;
            }
            current = &mut node.next;
        }
    }
}

// Example usage:
// let mut ll = LinkedList::new();
// ll.insert(3);
// ll.insert(7);
// ll.delete(3);
