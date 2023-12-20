class Queue {
    constructor() {
        this.items = [];
    }

    enqueue(item) {
        this.items.push(item);
    }

    dequeue() {
        if (!this.isEmpty()) {
            return this.items.shift();
        } else {
            throw new Error("dequeue from an empty queue");
        }
    }

    isEmpty() {
        return this.items.length === 0;
    }
}
