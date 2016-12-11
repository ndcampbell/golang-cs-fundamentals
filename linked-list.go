package main

import "fmt"

type Node struct {
    Value interface{}
    Next *Node
}

func addLast (head *Node, v interface{}) {

    newNode := Node{Value: v, Next: nil}
    tmp := head
    for tmp.Next != nil {
        tmp = tmp.Next
    }
    tmp.Next = &newNode
}

func printAll(head *Node) {
    tmp := head
    for tmp.Next != nil {
        fmt.Printf("%v\n", tmp)
        tmp = tmp.Next
    }
    fmt.Printf("%v\n", tmp)
    
}
func main() {
    head := Node{Value: 0, Next: nil}

    for i := 1; i <= 10; i++ {
        addLast(&head, i)
    }
    printAll(&head)

}
