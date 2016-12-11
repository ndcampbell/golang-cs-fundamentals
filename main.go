package main

import "fmt"
import "go-cs-fundamentals/linkedlist"

func main() {
    head := linkedlist.Node{Value: 0, Next: nil}
    for i := 1; i <=10; i++ {
         linkedlist.Push(&head, i)
    }
    fmt.Print("---Origional List---\n")
    linkedlist.PrintAll(&head)

    fmt.Print("\n\n---STACK Logic---\n\n")
    popped := linkedlist.Pop(&head)
    fmt.Printf("Popped Value: %v\n", popped)
    fmt.Print("---List After Pop---\n")
    linkedlist.PrintAll(&head)

    fmt.Print("\n\n---QUEUE Logic---\n\n")
    dequeued := linkedlist.Dequeue(&head)
    fmt.Printf("Dequeued Value: %v\n", dequeued)
    linkedlist.PrintAll(&head)
}
