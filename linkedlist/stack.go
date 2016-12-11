package linkedlist

//implements stack logic for the linked list. First in Last out (FILO)

// pops last value
func Pop(head *Node) Node {
    cur := head
    var prev *Node
    for cur.Next != nil {
        prev = cur
        cur = cur.Next
    }
    prev.Next = nil
    return *cur
}

//pushes value to back
func Push (head *Node, v interface{}) {
    newNode := Node{Value: v, Next: nil}
    tmp := head
    for tmp.Next != nil {
        tmp = tmp.Next
    }
    tmp.Next = &newNode
}

