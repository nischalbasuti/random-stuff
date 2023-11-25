package main

import ( "fmt" )

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

func print (l *List) {
    current := l.head

    fmt.Print("[")
    for current != nil {
        fmt.Print(current.data, " ")
        current = current.next
    }
    fmt.Print("]")
}

func main() {
    var l List = List{}
    l.add(1)
    l.add(2)
    l.add(3)

    print(&l)
}
