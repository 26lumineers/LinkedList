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
func (l linkList) printListData(){
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ",toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}
func (l *linkList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l *linkList) deleteWithValue(value int) {
	if l.length == 0 {
		return 
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	// fmt.Println("hello")
	myList := linkList{}
	node1 := &node{data: 48} //head
	node2 := &node{data: 12}
	node3 := &node{data: 55}
	node4 := &node{data: 87}
	node5 := &node{data: 62}
	node6 := &node{data: 35}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	// fmt.Println(myList)
	myList.printListData()
	myList.deleteWithValue(48)
	myList.printListData()
	emptyList := linkList{}
	emptyList.deleteWithValue(10)
	emptyList.printListData()

}
