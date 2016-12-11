package linkedlist

import "fmt"

type Node struct {
    Value interface{}
    Next *Node
}

// Recursive function for printing all nodes
func PrintAll(head *Node) {
    fmt.Printf("%v\n", head)
    if head.Next != nil {
        PrintAll(head.Next)
    }
}

