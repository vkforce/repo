package main

import "fmt"

type Node struct {
	data  int
	right *Node
	left  *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (d *DoubleLinkedList) Insert(data int) {
	newNode := &Node{data: data, left: nil, right: nil}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.left = d.tail
		d.tail.right = newNode
		d.tail = newNode
	}
}

func (d *DoubleLinkedList) PrintForward() {
	currentNode := d.head
	for currentNode != nil {
		fmt.Printf("%d\t", currentNode.data)
		currentNode = currentNode.right
	}
	fmt.Println()
}

func (d *DoubleLinkedList) AddNodeAtEnd(data int) {
	newNode := &Node{data: data, left: nil, right: nil}
	newNode.left = d.tail
	d.tail.right = newNode
	d.tail = newNode
}

func (d *DoubleLinkedList) AddNodeAtFront(data int) {
	newNode := &Node{
		data:  data,
		right: nil,
		left:  nil,
	}
	newNode.right = d.head
	d.head.left = newNode
	d.head = newNode
}

func (d *DoubleLinkedList) PrintReverse() {
	currentNode := d.tail
	for currentNode != nil {
		fmt.Printf("%d\t", currentNode.data)
		currentNode = currentNode.left
	}
	fmt.Println()
}

func main() {
	d := DoubleLinkedList{}
	d.Insert(10)
	d.Insert(20)
	d.Insert(30)
	d.Insert(40)
	d.Insert(50)
	d.AddNodeAtEnd(99)
	d.AddNodeAtFront(0)
	fmt.Println("Doubly linked list  (forward):")
	d.PrintForward()

	fmt.Println("Doubly linked list (reverse):")
	d.PrintReverse()

}
