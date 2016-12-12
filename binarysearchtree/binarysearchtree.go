package binarysearchtree

import "fmt"

type Node struct {
    Key interface{}
    Value interface{}
    Left *Node
    Right *Node
}

type BST struct {
    root *Node
}

//Public BST Functions

func (bst *BST) Search(key interface{}) Node {
    return Node{}

}

//Public entry for inserting a key/value node onto tree
func (bst *BST) Insert(key interface{}, v interface{}) {
    bst.root = insert(bst.root, key, v)
}


func (bst *BST) Delete(key interface{}) {
    return
}

//Prints all values in order in the tree
func (bst *BST) PrintAll() {
    traverse(bst.root, print_node)
}

//Private functions

//Define a callback function for nodes
type node_callback func(*Node)

//Function to print out the node. Used as a callback
func print_node(node *Node) {
    fmt.Printf("%v\n", node)
}

//recursively traverses the tree and runs a callback against each node.
func traverse(node *Node, callback node_callback) {
    if node == nil {
        return
    }
    traverse(node.Left, callback)
    callback(node)
    traverse(node.Right, callback)
}

// Recursively adds a node to the tree
func insert(node *Node, key interface{}, v interface{}) *Node {
    if node == nil {
        node = &Node{Key: key, Value: v}
    } else if key.(int) < node.Key.(int) {
        node.Left = insert(node.Left, key, v)
    } else {
        node.Right = insert(node.Right, key, v)
    }
    return node
}
