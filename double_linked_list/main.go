package main

import "fmt"

type Node struct {
	key  int
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (l *DoubleLinkedList) append(key int) {
	node := &Node{key: key}

	if l.tail == nil && l.head == nil {
		l.tail = node
		l.head = node
		return
	}

	l.tail.next = node
	node.prev = l.tail
	l.tail = node

}

func (l *DoubleLinkedList) prepend(key int) {
	node := &Node{key: key}

	if l.tail == nil && l.head == nil {
		l.tail = node
		l.head = node
		return
	}
	
	l.head.prev = node
	node.next = l.head
	l.head = node

}

func (l *DoubleLinkedList) pop() {
	l.tail = l.tail.prev
	l.tail.next = nil
}
func (l *DoubleLinkedList) shift() {
	l.head = l.head.next
	l.head.prev = nil
}

func (l *DoubleLinkedList) find(key int) *Node {
	current := l.head
	
	for current != nil{
		if current.key == key{
			return current
		}
		current = current.next	
	}

	return nil

}

func (l *DoubleLinkedList) delete(key int) {
	node := l.find(key)
	if l.head == node{
		l.head = node.next
	}
	if l.tail == node{
		l.tail = node.prev
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	
}


func (l *DoubleLinkedList) print() {
	current := l.head

	for current != nil {
		fmt.Println(current.key)
		current = current.next
	}

}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func main() {
	var list *DoubleLinkedList = NewDoubleLinkedList()
	list.append(8)
	list.append(10)
	list.append(9)
	list.prepend(50)
	list.prepend(60)
/* 	list.pop()
	list.shift() */
	list.delete(8)
	
	list.print()
	

}