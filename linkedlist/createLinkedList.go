package linkedlist

import "reflect"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head: &Node[T]{},
		size: 0,
	}
}

func (this *LinkedList[T]) Add(item T) {
	if this.size == 0 {
		this.head.data = item
	} else {
		newNode := Node[T]{
			data: item,
		}
		// traverse from head and append this
		temp := this.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &newNode
	}
	this.size += 1
}
func (this *LinkedList[T]) Get(index int) T {
	temp := this.head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp.data
}

func (this *LinkedList[T]) Remove(index int) {
	if index >= this.size || index < 0 {
		panic("index out of range")
	}
	temp := this.head
	temp2 := this.head
	for i := 0; i < index; i++ {
		temp2 = temp
		temp = temp.next
	}
	temp2.next = temp.next
	this.size -= 1
}

func (this *LinkedList[T]) Size() int {
	return this.size
}

func (this *LinkedList[T]) IsEmpty() bool {
	return this.size == 0
}

func (this *LinkedList[T]) Contains(data T) bool {
	temp := this.head
	for i := 0; i < this.size; i++ {
		if reflect.DeepEqual(temp.data, data) {
			return true
		}
	}
	return false
}
