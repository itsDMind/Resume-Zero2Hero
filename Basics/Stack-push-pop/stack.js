class Stack {
    constructor() {
        this.items = [];
    }

    push(item) {
        this.items.push(item);
    }

    pop() {
        if (!this.isEmpty()) {
            return this.items.pop();
        } else {
            throw new Error("pop from an empty stack");
        }
    }

    isEmpty() {
        return this.items.length === 0;
    }
}
