package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}

func (ll *LinkedList) insert(data int) {
    newNode := &Node{data: data, next: ll.head}
    ll.head = newNode
}

func (ll *LinkedList) delete(key int) {
    current := ll.head

    if current != nil && current.data == key {
        ll.head = current.next
        return
    }

    var prev *Node
    for current != nil && current.data != key {
        prev = current
        current = current.next
    }

    if current == nil {
        return
    }

    prev.next = current.next
}

// Example usage:
// ll := LinkedList{}
// ll.insert(3)
// ll.insert(7)
// ll.delete(3)
