package stack

import "golang-cs-fundamentals/list"

//implements stack logic for the linked list. First in Last out (FILO)

// pops last value
func Pop(head *list.Node) list.Node {
    cur := head
    var prev *list.Node
    for cur.Next != nil {
        prev = cur
        cur = cur.Next
    }
    prev.Next = nil
    return *cur
}

//pushes value to back
func Push (head *list.Node, v interface{}) {
    newNode := list.Node{Value: v, Next: nil}
    tmp := head
    for tmp.Next != nil {
        tmp = tmp.Next
    }
    tmp.Next = &newNode
}

