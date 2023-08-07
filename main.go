package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkList struct {
	head   *node
	length int
}

func (l *linkList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	fmt.Println("hello")
	myList := linkList{}
	node1 := &node{data: 48}
	node2 := &node{data: 12}
	node3 := &node{data: 55}
	node4 := &node{data: 87}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	fmt.Println(myList)
}
