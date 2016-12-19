package queue

import "fmt"

//Queue logic. First in, First out (FIFO)

type Node struct {
    Value interface{}
    Next *Node
}

type Queue struct {
    head *Node
    last *Node
}

func (q *Queue) Dequeue() Node {
    cur := *q.head
    *q.head = *q.head.Next
    return cur
}

func (q *Queue) Enqueue(v interface{}) {
    newNode := Node{Value: v, Next: nil}
    if q.head == nil {
        q.head = &newNode
        q.last = &newNode
    } else {
        tmp := q.last
        tmp.Next = &newNode
        q.last = &newNode
    }
}

func (q *Queue) PrintAll() {
    tmp := q.head
    for tmp.Next != nil {
        fmt.Printf("%v\n", tmp)
        tmp = tmp.Next
    }
    fmt.Printf("%v\n", tmp)
}
