package main

import "fmt"
import "golang-cs-fundamentals/stack"
import "golang-cs-fundamentals/queue"

func main() {
    fmt.Print("\n\n---STACK Logic---\n\n")
    s := stack.Stack{}

    for i := 0; i < 10; i++ {
         s.Push(i)
    }
    fmt.Print("---Stack Before Pop---\n")
    s.PrintAll()

    popped := s.Pop()
    fmt.Printf("\nPopped Value: %v\n", popped)
    fmt.Print("---Stack After Pop---\n")
    s.PrintAll()

    fmt.Print("\n\n---QUEUE Logic---\n\n")
    q := queue.Queue{}

    for i := 0; i < 10; i++ {
         q.Enqueue(i)
    }
    fmt.Print("---Queue Before Dequeue---\n")
    q.PrintAll()
    dequeued := q.Dequeue()
    fmt.Printf("Dequeued Value: %v\n", dequeued)
    fmt.Print("---Queue After Dequeue---\n")
    q.PrintAll()
}
