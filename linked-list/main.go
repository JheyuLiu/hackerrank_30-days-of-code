package main

import "fmt"

type Node struct {
	value int
	next *Node
}

type List struct {
    head *Node
}

func (list *List) insert(val int) {
    node := &Node{value: val, next: nil}
    if list.head == nil {
    	list.head = node
    } else {
    	tmp := list.head
    	for tmp.next != nil {
    		tmp = tmp.next
    	}
    	tmp.next = node
    }
}

func (list *List) display() {
    node := list.head
    for node != nil {
        fmt.Print(node.value, " ")
        node = node.next
    }
} 

func main() {
	var num, val int
	fmt.Scan(&num)

    list := List{}
	for num > 0 {
        fmt.Scan(&val)
        list.insert(val)
        num--
	}
	list.display()
}