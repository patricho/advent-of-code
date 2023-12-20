package shared

import "container/list"

type Point struct {
	Y, X int
}

var (
	Up         Point   = Point{Y: -1, X: 0}
	Down       Point   = Point{Y: 1, X: 0}
	Right      Point   = Point{Y: 0, X: 1}
	Left       Point   = Point{Y: 0, X: -1}
	Directions []Point = []Point{
		Up,
		Right,
		Down,
		Left,
	}
)

// Queue is a type safe-ish wrapper around container/list
type Queue[T any] struct {
	list *list.List
}

func CreateQueue[T any]() Queue[T] {
	return Queue[T]{
		list: list.New(),
	}
}

func (q Queue[T]) Len() int {
	return q.list.Len()
}

func (q Queue[T]) Enqueue(item T) {
	q.list.PushBack(item)
}

func (q Queue[T]) Dequeue() (item T) {
	item = q.list.Front().Value.(T)
	q.list.Remove(q.list.Front())
	return item
}
