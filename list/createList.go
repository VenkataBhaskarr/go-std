package list

type List[T any] struct {
	items []T
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (this *List[T]) Add(item T) {

}

func (this *List[T]) Get(index int) {

}

func (this *List[T]) Set(index int, item T) {

}

func (this *List[T]) Remove(index int) {
}

func (this *List[T]) Size() int {
}

func (this *List[T]) IsEmpty() bool {
}

func (this *List[T]) Clear() {
}

func (this *List[T]) Contains(item T) bool {
}

func (this *List[T]) IndexOf(item T) int {
}

func (this *List[T]) LastIndexOf(item T) int {
}

func (this *List[T]) SubList(start, end int) *List[T] {
}

func (this *List[T]) Sort() {
}

func (this *List[T]) Reverse() {
}
