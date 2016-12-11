package main

import "fmt"

type LinkedList struct {
    Value interface{}
    Next *LinkedList
}

func addLast (l *LinkedList, v interface{}) LinkedList {
    newNode := LinkedList{Value: v, Next: nil}
    l.Next = &newNode
    return newNode
}

func main() {
    head := LinkedList{Value: 0, Next: nil}
    newNode := addLast(&head, 1)

    fmt.Printf("%v\n", head.Next)
    fmt.Printf("%v\n", newNode.Next)

}
