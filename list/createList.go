package list

import "reflect"

type List[T any] struct {
	items []T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (this *List[T]) Add(item T) {
	this.items = append(this.items, item)
}

func (this *List[T]) Get(index int) T {
	if index < 0 || index >= len(this.items) {
		panic("index out of range")
	}
	return this.items[index]
}

func (this *List[T]) Set(index int, item T) {
	if index < 0 || index >= len(this.items) {
		panic("index out of range")
	}
	this.items[index] = item
}

func (this *List[T]) Remove(index int) {
	if index < 0 || index >= len(this.items) {
		panic("index out of range")
	}
	this.items = append(this.items[:index], this.items[index+1:]...)
}

func (this *List[T]) Size() int {
	return len(this.items)
}

func (this *List[T]) IsEmpty() bool {
	return len(this.items) == 0
}

func (this *List[T]) Clear() {
	this.items = []T{}
}

func (this *List[T]) Contains(item T) bool {
	for _, v := range this.items {
		if reflect.DeepEqual(v, item) {
			return true
		}
	}
	return false
}

func (this *List[T]) IndexOf(item T) int {
	for i, v := range this.items {
		if reflect.DeepEqual(v, item) {
			return i
		}
	}
	return -1
}

func (this *List[T]) LastIndexOf(item T) int {
	for i := len(this.items) - 1; i >= 0; i-- {
		if reflect.DeepEqual(this.items[i], item) {
			return i
		}
	}
	return -1
}

func (this *List[T]) SubList(start, end int) *List[T] {
	if start < 0 || start >= len(this.items) {
		panic("start index out of range")
	}
	if end < 0 || end >= len(this.items) {
		panic("end index out of range")
	}
	if start > end {
		panic("start index is greater than end index")
	}
	return &List[T]{items: this.items[start:end]}
}
