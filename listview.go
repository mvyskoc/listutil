package listutil

import "container/list"

type ListView struct {
	Front *list.Element
	Back  *list.Element
}

func (l *ListView) Len() int {
	var count int = 0

	for el := l.Front; el != nil; el = l.Next(el) {
		count++
	}
	return count
}

func (l *ListView) Next(el *list.Element) *list.Element {
	if el != l.Back {
		return el.Next()
	}
	return nil
}

func (l *ListView) Prev(el *list.Element) *list.Element {
	if el != l.Front {
		return el.Prev()
	}
	return nil
}

func ListViewToSlice[T any](lw ListView) []T {
	result := make([]T, 0, lw.Len())
	for el := lw.Front; el != nil; el = lw.Next(el) {
		result = append(result, el.Value.(T))
	}
	return result
}
