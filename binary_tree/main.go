package main

import (
	"fmt"
)
type Node struct{
	left *Node
	right *Node
	key int
}
type Tree struct{
	root *Node
}
func (t *Tree)append(data int) {
	node := Node{left: nil, right: nil, key: data}
	if t.root == nil{
		t.root = &node
		return
	}

	current := t.root
	for current != nil{
		if data <= current.key{
			if current.left == nil{
				current.left = &node 
				return
			}
			current = current.left
		}else{
			if current.right == nil{
				current.right = &node 
				return	
			}
			current = current.right
		}
	}
}

func (t *Tree) print(node *Node){
	if node == nil{
		return
	}
	t.print(node.left)
	fmt.Println(node.key)

	t.print(node.right)
}

func (t *Tree) find(data int) *Node{
	current := t.root
	for current != nil{
		if current.key == data{
			return current
		}
		if data < current.key{
			current = current.left
		}else{
			current = current.right
		}
	}
	return nil

}
func (t *Tree) erase(data int){
	current := t.root
	var prev *Node = nil

	for current != nil && current.key != data {
		prev = current
		if(data < current.key) {
			current = current.left
		} else {
			current = current.right
		}
	}

	if(current == nil) {
		return
	}

	if(current.left == nil) {
		if(prev.left == current) {
			prev.left = current.right
			return
		} else {
			prev.right = current.right
			return
		}
	}
	if(current.right == nil) {
		if(prev.left == current) {
			prev.left = current.left
			return
		} else {
			prev.right = current.left
			return
		}
	}
	mini := t.min(current.right)
	t.erase(mini.key)
	current.key = mini.key


}
func (t *Tree) min(node *Node) *Node{
	if node.left != nil{
		return t.min(node.left)	
	}
	
	return node

	/* for node.left != nil{
		node = node.left
	}
	return node */
}

func main(){
	tree := Tree{root: nil}
	tree.append(5)
	tree.append(6)
	tree.append(4)

	tree.append(3)
	tree.append(10)
	tree.append(20)
	tree.append(17)
	tree.append(8)
	tree.append(7)
	tree.append(9)
	

	tree.erase(10)
	tree.print(tree.root)


}