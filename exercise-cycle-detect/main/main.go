package main

import (
	"fmt"
	"errors"
)

type Node struct {
	value int
	next *Node
}

func main() {
	var firstNode = Node{}
	var secondNode = Node{1, &firstNode}
	var thirdNode = Node{2, &secondNode}
	var forthNode = Node{3, &thirdNode}

	firstNode = Node{10, &forthNode}

	var loopedNode, noLoop = detectCycle(&forthNode)

	if noLoop != nil {
		fmt.Println(noLoop)
	} else {
		fmt.Println("Found", loopedNode)
	}
}

func detectCycle(rootValues *Node) (Node, error) {
	var registeredNode = map[*Node]*Node{}
	var currentNode = rootValues

	for currentNode.next != nil {
		if registeredNode[currentNode.next] != nil {
			return *currentNode, nil
		}

		registeredNode[currentNode] = currentNode
		currentNode = currentNode.next
	}

	return Node{}, errors.New("No cycle found")
}