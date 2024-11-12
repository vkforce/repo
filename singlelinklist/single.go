package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkList struct {
	head *Node
}

func (list *LinkList) InsertNode(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.next = newNode
	}
}

func (list *LinkList) Display() {
	currentNode := list.head
	if currentNode != nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Print("Linked list: ")
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}
func main() {
	list := LinkList{}
	list.InsertNode(10)
	list.InsertNode(20)
	list.InsertNode(30)
	list.InsertNode(40)
	list.Display()
}
