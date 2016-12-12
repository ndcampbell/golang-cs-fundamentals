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

func (bst *BST) Search(key interface{}) Node {
    return Node{}

}

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

func (bst *BST) Insert(key interface{}, v interface{}) {
    node := insert(bst.root, key, v)
    fmt.Print(node)
}


func (bst *BST) Delete(key interface{}) {
    return

}
