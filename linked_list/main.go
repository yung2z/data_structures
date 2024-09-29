package main

import (
	"fmt"
)
 
type Node[T string | int] struct{
	next *Node[T]
	value T
}

type LinkedList[T string | int] struct{
	head *Node[T]
}

func (list *LinkedList[T]) append(data T){
	var newData Node[T] = Node[T]{value: data, next: nil}
	if list.head == nil{
		list.head = &newData
		return
	}

	current := list.head
	for current.next != nil{
		current = current.next
	}
	current.next = &newData
}
func (list *LinkedList[T]) prepend(data T){

	tmp := list.head
	var newData Node[T] = Node[T]{value: data, next: tmp}

	list.head = &newData  

}

func (list *LinkedList[T]) print(node *Node[T]){
	if node == nil {
		return
	}
	fmt.Println(node.value)
	list.print(node.next)
}
func (list *LinkedList[T]) erase (data T){
	current := list.head
	for current.next != nil{
		if current.next.value == data{
			current.next = current.next.next
		}
		current = current.next
	}
}

func (list *LinkedList[T]) find (data T) *Node[T]{
	current := list.head
	for current.next != nil{
		if current.next.value == data{
			return current.next
		}
		current = current.next
	}
	return nil 
}

func NewLinkedList[T string | int]() *LinkedList[T] {
	return &LinkedList[T]{
		head: nil,
	};
}

func main() {
	var list *LinkedList[string] = &LinkedList[string]{head: nil};
	list.append("123")
	list.prepend("1111111")
	list.print(list.head)
}