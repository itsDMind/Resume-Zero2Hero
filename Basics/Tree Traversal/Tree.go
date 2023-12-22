package main

import "fmt"

type TreeNode struct {
	value       int
	left, right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil}
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{nil}
}

func (bt *BinaryTree) Insert(value int) {
	bt.root = bt.insert(bt.root, value)
}

func (bt *BinaryTree) insert(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return NewTreeNode(value)
	}

	if value < root.value {
		root.left = bt.insert(root.left, value)
	} else {
		root.right = bt.insert(root.right, value)
	}

	return root
}

func (bt *BinaryTree) InOrderTraversal() []int {
	result := []int{}
	bt.inOrderTraversal(bt.root, &result)
	return result
}

func (bt *BinaryTree) inOrderTraversal(node *TreeNode, result *[]int) {
	if node != nil {
		bt.inOrderTraversal(node.left, result)
		*result = append(*result, node.value)
		bt.inOrderTraversal(node.right, result)
	}
}

func (bt *BinaryTree) PreOrderTraversal() []int {
	result := []int{}
	bt.preOrderTraversal(bt.root, &result)
	return result
}

func (bt *BinaryTree) preOrderTraversal(node *TreeNode, result *[]int) {
	if node != nil {
		*result = append(*result, node.value)
		bt.preOrderTraversal(node.left, result)
		bt.preOrderTraversal(node.right, result)
	}
}

func (bt *BinaryTree) PostOrderTraversal() []int {
	result := []int{}
	bt.postOrderTraversal(bt.root, &result)
	return result
}

func (bt *BinaryTree) postOrderTraversal(node *TreeNode, result *[]int) {
	if node != nil {
		bt.postOrderTraversal(node.left, result)
		bt.postOrderTraversal(node.right, result)
		*result = append(*result, node.value)
	}
}

// Example usage:
func main() {
	myTree := NewBinaryTree()
	myTree.Insert(5)
	myTree.Insert(3)
	myTree.Insert(7)
	myTree.Insert(2)
	myTree.Insert(4)
	myTree.Insert(6)
	myTree.Insert(8)

	fmt.Println("In-order traversal:", myTree.InOrderTraversal())
	fmt.Println("Pre-order traversal:", myTree.PreOrderTraversal())
	fmt.Println("Post-order traversal:", myTree.PostOrderTraversal())
}
