package main

import (
    "fmt"
    "math"
)

type Node struct {
    data int
    next *Node
}

type List struct {
    head *Node
}

func (this *List) add(value int) {
    var newNodeInstance = Node{data: value}
    var newNodePointer *Node = &newNodeInstance

    if this.head == nil {
        this.head = newNodePointer
        return
    }

    current := this.head

    for current.next != nil {
        current = current.next
    }

    current.next = newNodePointer

    return
}

func (l *List) asNumber() int {
    var current *Node = l.head
    var number int = 0

    var length int = 0
    for current != nil {
        current = current.next
        length += 1
    }

    current = l.head
    var index int = 0
    for current != nil {
        index += 1
        place := float64(length - index)

        number += current.data * int(math.Pow(10, place))

        current = current.next
    }
    

    return number
}

func (l *List) print() {
    current := l.head

    fmt.Print("[")
    for current != nil {
        fmt.Print(current.data, " ")
        current = current.next
    }
    fmt.Print("]\n")
}

func main() {
    var n1 List = List{}
    n1.add(1)
    n1.add(0)
    n1.add(0)
    n1.add(0)

    var n2 List = List{}
    n2.add(0)
    n2.add(0)
    n2.add(0)
    n2.add(1)

    n1.print()
    n2.print()

    fmt.Print(n1.asNumber())
    fmt.Print(" x ")
    fmt.Print(n2.asNumber())
    fmt.Print(" = ")
    fmt.Println(n1.asNumber() * n2.asNumber())

    fmt.Println("----end----")
}
