package main

import (
	"fmt"
	"errors"
	"math/rand"
)

type Tree struct {
	node Node
}

type Node struct {
	value int
	sx *Node
	dx *Node
}

func initNode(value int) (Node) {
	return Node {
		value: value,
		sx: &Node{},
		dx: &Node{},
	}
}

func main() {
	tree := &Tree{}
	lenValue := rand.Intn(30)

	for i:=0; i < lenValue; i++ {
		value := rand.Intn(100)
		fmt.Printf("add: %d", value)
		add(Node{value, &Node{}, &Node{}}, tree)
	}

	println("")
	printTree(*tree)
}

func add(node Node, tree *Tree) (Node, error) {
	return addToTree(node, &tree.node)
}

func printTree(tree Tree) {
	recursivePrint(tree.node)
}

func recursivePrint(node Node) {
	fmt.Println("node", node)

	if node.value != 0 {
		recursivePrint(*node.sx)
		recursivePrint(*node.dx)
	}
}

func addToTree(node Node, current *Node) (Node, error) {
	if current.value == 0 {
		current.value = node.value
		current.sx = &Node{}
		current.dx = &Node{}
		return node, nil
	} else if current.value == node.value {
		return Node{}, errors.New("found duplicate")
	}

	if node.value > current.value {
		return addToTree(node, current.dx)
	} else {
		return addToTree(node, current.sx)
	}
}