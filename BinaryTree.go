package main

import "fmt"

var val interface{}

type Node struct {
	val   interface{}
	left  *Node
	right *Node
}

func addNode(root *Node, val interface{}) {
	if root == nil {
		root = &Node{val: val}
		return
	}
	if root.left == nil {
		root.left = &Node{val: val}
		return
	}
	if root.right == nil {
		root.right = &Node{val: val}
		return
	}
	addNode(root.left, val)
}

func main() {
	root := &Node{val: 1}
	addNode(root, 2)
	addNode(root, 3)
	addNode(root, 4)

	fmt.Println(root.left.left.val)

}
