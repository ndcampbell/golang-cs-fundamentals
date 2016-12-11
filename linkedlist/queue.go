package linkedlist

//Queue logic. First in, First out (FIFO)

func Dequeue(head *Node) Node {
    cur := *head
    *head = *head.Next
    return cur
}

func Enqueue (head *Node, v interface{}) {
    newNode := Node{Value: v, Next: nil}
    tmp := head
    for tmp.Next != nil {
        tmp = tmp.Next
    }
    tmp.Next = &newNode
}
