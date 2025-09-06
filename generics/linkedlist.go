package main

import (
	"fmt"
)

// type Node[T comparable] struct {
//   val T
//   next *Node[T]
// }
//
// type LinkedList[T comparable] struct {
//   next *Node[T]
// }
//
// func NewLinkedList[T comparable](items []T) *LinkedList[T] {
//   list := &LinkedList[T]{}
//
//   head := Node[t]{}
//
//   for i, value := range items {
//      if(i == 0){
//
//      }
//   }
//
//
//
//
// }

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

// PrintElements is not required since we implemented Striger method,
// so fmt package would know how to print the linkedlist
func (list *LinkedList) PrintElements() {
	current := list.head

	if current == nil {
		fmt.Println("No Elements found")
		return
	}

	ItemsInList := []int{}

	for current != nil {
		ItemsInList = append(ItemsInList, current.val)
		current = current.next
	}

	fmt.Println(ItemsInList)

}

func (list *LinkedList) String() string {
	current := list.head
	output := []int{}

	if current == nil {
		return "[]"
	}

	for current != nil {
		output = append(output, current.val)
		current = current.next
	}

	return fmt.Sprintf("%v", output)
}

func (list *LinkedList) Add(num int) {
	current := list.head

	if current == nil {
		list.head = &Node{val: num}
		return
	}

	for current.next != nil {
		current = current.next
	}

	current.next = &Node{val: num}

}

func (list *LinkedList) Remove(num int) {

	if list.head == nil {
		return
	}

	current := list.head
  var prev *Node = nil

	for current != nil {
    if current.val == num {
			// Removing head
			if prev == nil {
				list.head = current.next
				break
			}

			prev.next = current.next
			break
		}

		prev = current
		current = current.next

	}

}

func InitLinkedList() *LinkedList {
	return &LinkedList{}
}

func main() {

	list1 := InitLinkedList()

	fmt.Println(list1)
	list1.PrintElements()

	list1.Add(87)
	list1.Add(3)
	list1.Add(8)
	list1.Add(3)
	list1.Add(28)
	list1.Add(63)

	fmt.Println(list1)

  fmt.Println("Removing number 3")
	list1.Remove(63)
	fmt.Println(list1)

}
