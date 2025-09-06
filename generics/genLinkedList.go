package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

type LinkedList[T comparable] struct {
	head *List[T]
}

func (list *LinkedList[T]) add(val T) {
	current := list.head

	if current == nil {
		list.head = &List[T]{val: val}
    return;
	}

	for {
		if current.next == nil {
			current.next = &List[T]{val: val}
			break
		}

		current = current.next
	}

}

func (list *LinkedList[T]) String() string {
	if list.head == nil {
		return "[]"
	}

	current := list.head
	items := make([]T, 0)

	for current != nil {
		items = append(items, current.val)
		current = current.next
	}

	return fmt.Sprintf("%v", items)

}

func main() {
	list1 := &LinkedList[string]{}

	fmt.Println(list1)

	list1.add("apple")
	list1.add("banana")
	list1.add("orange")
	fmt.Println(list1)

	numList := &LinkedList[int64]{}
	fmt.Println(numList)

  numList.add(23)
  numList.add(23)
  numList.add(23)
  numList.add(23)
	fmt.Println(numList)


}
