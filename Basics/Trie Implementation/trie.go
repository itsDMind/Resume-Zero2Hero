package main

import "fmt"

type TrieNode struct {
	children       map[rune]*TrieNode
	isEndOfWord    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children:    make(map[rune]*TrieNode),
		isEndOfWord: false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true
}

// Example usage:
func main() {
	trie := NewTrie()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))     // Output: true
	fmt.Println(trie.Search("app"))       // Output: false
	fmt.Println(trie.StartsWith("app"))   // Output: true
	trie.Insert("app")
	fmt.Println(trie.Search("app"))       // Output: true
}
