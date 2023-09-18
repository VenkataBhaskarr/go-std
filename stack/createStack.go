package stack

type Stack[T any] struct {
	items   []T
	pointer int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items:   make([]T, 0),
		pointer: 0,
	}
}

func (this *Stack[T]) Add(item T) {
	this.items = append(this.items, item)
	this.pointer += 1
}

func (this *Stack[T]) Pop() T {
	if this.pointer == 0 {
		panic("empty stack")
	}
	this.pointer -= 1
	return this.items[this.pointer]
}

func (this *Stack[T]) Peek() T {
	if this.pointer == 0 {
		panic("empty stack")
	}
	return this.items[this.pointer-1]
}

func (this *Stack[T]) Size() int {
	return this.pointer
}
