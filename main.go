package main

import "fmt"
import "golang-cs-fundamentals/list"
import "golang-cs-fundamentals/stack"
import "golang-cs-fundamentals/queue"

func main() {
    head := list.Node{Value: 0, Next: nil}
    for i := 1; i <=10; i++ {
         stack.Push(&head, i)
    }
    fmt.Print("---Origional List---\n")
    list.PrintAll(&head)

    fmt.Print("\n\n---STACK Logic---\n\n")
    popped := stack.Pop(&head)
    fmt.Printf("Popped Value: %v\n", popped)
    fmt.Print("---List After Pop---\n")
    list.PrintAll(&head)

    fmt.Print("\n\n---QUEUE Logic---\n\n")
    dequeued := queue.Dequeue(&head)
    fmt.Printf("Dequeued Value: %v\n", dequeued)
    list.PrintAll(&head)
}
