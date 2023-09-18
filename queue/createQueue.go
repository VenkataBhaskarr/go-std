package queue

type Queue[T any] struct {
	items   []T
	pointer int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items:   make([]T, 0),
		pointer: 0,
	}
}

func (this *Queue[T]) Add(item T) {
	this.items = append(this.items, item)
}

func (this *Queue[T]) Poll() T {
	if len(this.items) == 0 {
		panic("empty queue")
	}
	this.items = this.items[1:]
	return this.items[0]
}

func (this *Queue[T]) Peek() T {
	if len(this.items) == 0 {
		panic("empty queue")
	}
	return this.items[0]
}

func (this *Queue[T]) Size() int {
	return len(this.items)
}
