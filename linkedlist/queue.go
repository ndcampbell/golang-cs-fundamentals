package linkedlist

func dequeue(head *Node) *Node {
    cur := head
    var prev *Node
    for cur.Next != nil {
        prev = cur
        cur = cur.Next
    }
    prev.Next = nil
    return cur
}

func enqueue (head *Node, v interface{}) {
    newNode := Node{Value: v, Next: nil}
    tmp := head
    for tmp.Next != nil {
        tmp = tmp.Next
    }
    tmp.Next = &newNode
}