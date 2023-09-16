package list

type List[T any] struct {
  items []T
}

func NewList[T any]() *List[T] {
  return &List[T]{}
}

func(this *List[T]) add(item T) {
  // add a given item
}

func(this *List[T]) get(index int){
  // returns the element in that index
}

func(this *List[T]) set(index int, item T){

}

func (this *List[T]) remove(index int) {
}

func (this *List[T]) size() int {
}

func (this *List[T]) isEmpty() bool {
}

func (this *List[T]) clear() {
}

func (this *List[T]) contains(item T) bool {
}

func (this *List[T]) indexOf(item T) int {
}

func (this *List[T]) lastIndexOf(item T) int {
}

func (this *List[T]) subList(start, end int) *List[T] {
}

func (this *List[T]) sort() {
}

func (this *List[T]) reverse() {
}

// More to be added 
