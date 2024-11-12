package main

import "fmt"

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) InsertBinaryTree(data int) {
	newNode := &BinaryNode{data: data, left: nil, right: nil}
	if t.root == nil {
		t.root = newNode
		return
	}
	currentNode := t.root
	for {
		if data < currentNode.data {
			if currentNode.left == nil {
				currentNode.left = newNode
				return
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = newNode
				return
			}
			currentNode = currentNode.right
		}
	}
}

func (t *BinaryTree) InsertNode(data int) {
	newNode := &BinaryNode{data: data, left: nil, right: nil}
	if t.root == nil {
		t.root = newNode
		return
	}
	queue := []*BinaryNode{t.root}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:] // remove the first element from the queue
		if currentNode.left == nil {
			currentNode.left = newNode
			return
		} else if currentNode.right == nil {
			currentNode.right = newNode
			return
		} else {
			queue = append(queue, currentNode.left, currentNode.right)
		}
	}
}

func (t *BinaryTree) PrintTree() {
	if t.root == nil {
		fmt.Println("Tree is Empty!!")
		return
	}
	queue := []*BinaryNode{t.root}
	level := 0
	for len(queue) > 0 {
		fmt.Println("Level", level)
		currentLevelSize := len(queue)
		for i := 0; i < currentLevelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			fmt.Print(currentNode.data, " ")
			if currentNode.left != nil {
				queue = append(queue, currentNode.left)
			}
			if currentNode.right != nil {
				queue = append(queue, currentNode.right)
			}
		}
	}
}

func (t *BinaryTree) DFS(node *BinaryNode) {
	if node == nil {
		return
	}
	fmt.Print(node.data, " ")
	t.DFS(node.left)
	t.DFS(node.right)
}

func (t *BinaryTree) BFS() {
	if t.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	queue := []*BinaryNode{t.root}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		fmt.Print(currentNode.data, " ")
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}

}
func main() {
	tree := BinaryTree{}
	tree.InsertNode(4)
	tree.InsertNode(6)
	tree.InsertNode(1)
	tree.InsertNode(3)
	tree.InsertNode(10)
	tree.InsertNode(11)
	tree.InsertNode(12)
	tree.PrintTree()
	fmt.Println()
	fmt.Print("DFS: ")
	tree.DFS(tree.root)
	fmt.Println()
	fmt.Print("BFS: ")
	tree.BFS()

	binarytree := BinaryTree{}
	binarytree.InsertBinaryTree(4)
	binarytree.InsertBinaryTree(6)
	binarytree.InsertBinaryTree(1)
	binarytree.InsertBinaryTree(3)
	binarytree.InsertBinaryTree(10)
	binarytree.InsertBinaryTree(11)
	binarytree.InsertBinaryTree(12)
	binarytree.PrintTree()
}
