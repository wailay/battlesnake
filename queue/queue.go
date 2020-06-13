package queue
import "container/list"
import . "../utils"
import "errors"
func New() *list.List {
	l := list.New()
	return l
}

func Push(l *list.List, x Point){
	l.PushBack(x)
}

func Pop(l *list.List) (Point, error) {
	var p Point
	front := l.Front()
	if front == nil { return p , errors.New("empty")}
	l.Remove(front)
	p = front.Value.(Point)
	return p, nil
}

func Empty(l *list.List) bool {
	return l.Len() < 1
}