package stack

import "fmt"

//implements stack logic using linked list. First in Last out (FILO)

type Node struct {
    Value interface{}
    Next *Node
}

type Stack struct {
    head *Node
}

// pops last value from stack. Returns whole node
func (s *Stack) Pop() Node {
    cur := s.head
    var prev *Node
    for cur.Next != nil {
        prev = cur
        cur = cur.Next
    }
    prev.Next = nil
    return *cur
}

//pushes value on top of Stack.
func (s *Stack) Push(v interface{}) {
    newNode := Node{Value: v, Next: nil}
    if s.head == nil {
        s.head = &newNode
    } else {
        tmp := s.head
        for tmp.Next != nil {
            tmp = tmp.Next
        }
        tmp.Next = &newNode
    }
}

// Prints entire stack
func (s *Stack) PrintAll() {
    tmp := s.head
    for tmp.Next != nil {
        fmt.Printf("%v\n", tmp)
        tmp = tmp.Next
    }
    fmt.Printf("%v\n", tmp)
}
