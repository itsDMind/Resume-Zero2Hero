class Node {
    constructor(data) {
        this.data = data;
        this.next = null;
    }
}

class LinkedList {
    constructor() {
        this.head = null;
    }

    insert(data) {
        const newNode = new Node(data);
        newNode.next = this.head;
        this.head = newNode;
    }

    delete(key) {
        let current = this.head;

        if (current !== null && current.data === key) {
            this.head = current.next;
            return;
        }

        let prev = null;
        while (current !== null && current.data !== key) {
            prev = current;
            current = current.next;
        }

        if (current === null) {
            return;
        }

        prev.next = current.next;
    }
}

// Example usage:
// const ll = new LinkedList();
// ll.insert(3);
// ll.insert(7);
// ll.delete(3);
