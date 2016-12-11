package queue

import "golang-cs-fundamentals/list"

//Queue logic. First in, First out (FIFO)

func Dequeue(head *list.Node) list.Node {
    cur := *head
    *head = *head.Next
    return cur
}

func Enqueue (head *list.Node, v interface{}) {
    newNode := list.Node{Value: v, Next: nil}
    tmp := head
    for tmp.Next != nil {
        tmp = tmp.Next
    }
    tmp.Next = &newNode
}
