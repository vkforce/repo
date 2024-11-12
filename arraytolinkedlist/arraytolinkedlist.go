package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func ArrayToLinkedList(arr []int) *Node {
	var head, tail *Node
	for _, val := range arr {
		newNode := &Node{
			data: val,
			next: nil,
		}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.next = newNode
			tail = newNode
		}
	}
	return head
}

func PrintLinkedList(head *Node) {
	//currentNode := head
	//for currentNode != nil {
	//	fmt.Printf("%d ", currentNode.data)
	//	currentNode = currentNode.next
	//}
	//fmt.Println()
	if head != nil {
		fmt.Printf("%d ", head.data)
		PrintLinkedList(head.next)
	} else {
		fmt.Println()
	}
}

func DeleteNode(head *Node, value int) *Node {
	if head != nil && head.data == value {
		return head.next
	}
	currentNode := head
	var prevNode *Node

	for currentNode != nil && currentNode.data != value {
		prevNode = currentNode
		currentNode = currentNode.next
	}
	if currentNode != nil {
		prevNode.next = currentNode.next
	}
	return head
}

func AddBeforeXElement(head *Node, data int, existingValue int) *Node {
	newNode := &Node{
		data: data,
		next: nil,
	}
	var previousNode *Node
	currentNode := head
	for currentNode != nil {
		if currentNode.data == existingValue {
			if currentNode == head {
				newNode.next = currentNode
				newNode = head
			} else {
				previousNode.next = newNode
				newNode.next = currentNode
			}
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return head
}

func AddAfterXElement(head *Node, data int, existingValue int) *Node {
	newNode := &Node{
		data: data,
		next: nil,
	}
	currentNode := head
	for currentNode != nil {
		if currentNode.data == existingValue {
			if currentNode == head {
				newNode.next = currentNode
				newNode = head
			} else {
				newNode.next = currentNode.next
				currentNode.next = newNode
			}
		}
		currentNode = currentNode.next
	}
	return head
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	linkedList := ArrayToLinkedList(arr)
	PrintLinkedList(linkedList)
	//DeleteNode(linkedList, 30)
	AddBeforeXElement(linkedList, 22, 30)
	//AddAfterXElement(linkedList, 32, 30)
	PrintLinkedList(linkedList)
}
