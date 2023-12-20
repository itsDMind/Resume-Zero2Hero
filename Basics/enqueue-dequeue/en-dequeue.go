package main

import "fmt"

type Queue struct {
    items []int
}

func (q *Queue) enqueue(item int) {
    q.items = append(q.items, item)
}

func (q *Queue) dequeue() int {
    if !q.isEmpty() {
        item := q.items[0]
        q.items = q.items[1:]
        return item
    } else {
        panic("dequeue from an empty queue")
    }
}

func (q *Queue) isEmpty() bool {
    return len(q.items) == 0
}
