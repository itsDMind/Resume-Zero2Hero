package main

import "fmt"

type Stack struct {
    items []int
}

func (s *Stack) push(item int) {
    s.items = append(s.items, item)
}

func (s *Stack) pop() int {
    if !s.isEmpty() {
        item := s.items[len(s.items)-1]
        s.items = s.items[:len(s.items)-1]
        return item
    } else {
        panic("pop from an empty stack")
    }
}

func (s *Stack) isEmpty() bool {
    return len(s.items) == 0
}
