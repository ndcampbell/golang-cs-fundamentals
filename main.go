package main

import "fmt"
import "golang-cs-fundamentals/stack"
import "golang-cs-fundamentals/queue"
import "golang-cs-fundamentals/binarysearchtree"

// Runs a demo of how the stack work with push and pop
func demo_stack() {
    fmt.Print("\n---STACK Logic---\n\n")
    s := stack.Stack{}

    for i := 0; i <= 5; i++ {
         s.Push(i)
    }
    fmt.Print("---Stack Before Pop---\n")
    s.PrintAll()

    popped := s.Pop()
    fmt.Printf("\nPopped Value: %v\n", popped)
    fmt.Print("---Stack After Pop---\n")
    s.PrintAll()
}

// Runs a demo of how the queue works, with enqueue and dequeue
func demo_queue() {
    fmt.Print("\n---QUEUE Logic---\n\n")
    q := queue.Queue{}

    for i := 0; i <= 5; i++ {
         q.Enqueue(i)
    }
    fmt.Print("---Queue Before Dequeue---\n")
    q.PrintAll()
    dequeued := q.Dequeue()
    fmt.Printf("Dequeued Value: %v\n", dequeued)
    fmt.Print("---Queue After Dequeue---\n")
    q.PrintAll()
}

// Runs a demo of the Binary Search Tree. WIP
func demo_bst() {
    fmt.Print("\n---Binary Search Tree Demo---\n\n")
    bst := binarysearchtree.BST{}
    bst.Insert(20, 0)
    bst.Insert(15, 0)
    bst.Insert(30, 0)
    bst.Insert(2, 0)
    fmt.Print("---All Nodes in Tree---\n")
    bst.PrintAll()
}

func main() {
    demo_stack()
    demo_queue()
    demo_bst()
    }
