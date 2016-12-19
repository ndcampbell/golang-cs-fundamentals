package binarysearchtree

import "fmt"

type Node struct {
    Key interface{}
    Value interface{}
    Parent *Node
    Left *Node
    Right *Node
}

type BST struct {
    root *Node
}

//Public BST Functions

func (bst *BST) Search(key interface{}) Node {
    foundnode := search_node(bst.root, key)
    return *foundnode
}

//Public entry for inserting a key/value node onto tree
func (bst *BST) Insert(key interface{}, v interface{}) {
    bst.root = insert(bst.root, bst.root, key, v)
}


func (bst *BST) Delete(key interface{}) {
    delete_node(bst.root, key)
}

//Prints all values in order in the tree
func (bst *BST) PrintAll() {
    traverse(bst.root, print_node)
}

//Private functions

func replace_node(parent *Node, child *Node) {
    if child != nil {
        if parent.Left != nil {
            parent.Left = child
        } else {
            parent.Right = child
        }
    }
}

func delete_node(node *Node, key interface{}) {
    if key.(int) > node.Key.(int) {
        delete_node(node.Right, key)
    } else if key.(int) < node.Key.(int) {
        delete_node(node.Left, key)
    } else {
        if node.Left != nil && node.Right != nil {
            fmt.Print("No logic for deleting node with two children")
        } else if node.Left != nil {
            replace_node(node.Parent, node.Left)
        } else if node.Right !=nil {
            replace_node(node.Parent, node.Right)
        } else {
            replace_node(node.Parent, nil)
        }
    }
}

func search_node(node *Node, key interface{}) *Node {
    if node.Key == key {
        return node
    }
    if node == nil {
        return node
    }
    if key.(int) < node.Key.(int) {
        return search_node(node.Left, key)
    } else {
        return search_node(node.Right, key)
    }
}

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
func insert(node *Node, parent *Node, key interface{}, v interface{}) *Node {
    if node == nil {
        node = &Node{Parent: parent, Key: key, Value: v}
    } else if key.(int) < node.Key.(int) {
        node.Left = insert(node.Left, node, key, v)
    } else {
        node.Right = insert(node.Right, node, key, v)
    }
    return node
}
