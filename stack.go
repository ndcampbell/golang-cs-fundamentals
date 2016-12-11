package main

import "fmt"

type Node struct {
    Value interface{}
    Next *Node
}

// pops last value
func pop(head *Node) *Node {
    cur := head
    var prev *Node
    for cur.Next != nil {
        prev = cur
        cur = cur.Next
    }
    prev.Next = nil
    return cur
}

//pushes value to back
func push (head *Node, v interface{}) {
    newNode := Node{Value: v, Next: nil}
    tmp := head
    for tmp.Next != nil {
        tmp = tmp.Next
    }
    tmp.Next = &newNode
}

// Resurive function for printing all nodes
func printAll(head *Node) {
    fmt.Printf("%v\n", head)
    if head.Next != nil {
        printAll(head.Next)
    }
}

func main() {
    head := Node{Value: 0, Next: nil}

    for i := 1; i <= 10; i++ {
        push(&head, i)
    }
    fmt.Print("---List Before Pop---\n")
    printAll(&head)
    popped := pop(&head)
    fmt.Printf("Popped Value: %v\n", popped)
    fmt.Print("---List After Pop---\n")
    printAll(&head)

}
