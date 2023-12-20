struct Queue {
    items: Vec<i32>,
}

impl Queue {
    fn new() -> Self {
        Queue { items: Vec::new() }
    }

    fn enqueue(&mut self, item: i32) {
        self.items.push(item);
    }

    fn dequeue(&mut self) -> Option<i32> {
        if !self.is_empty() {
            Some(self.items.remove(0))
        } else {
            None
        }
    }

    fn is_empty(&self) -> bool {
        self.items.is_empty()
    }
}
