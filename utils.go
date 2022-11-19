package listutil

import "container/list"

func ToList[T any](arr []T) *list.List {
	result := list.New()
	for _, v := range arr {
		result.PushBack(v)
	}
	return result
}

func ToSlice[T any](l *list.List) []T {
	result := make([]T, l.Len())
	i := 0
	for element := l.Front(); element != nil; element = element.Next() {
		result[i] = element.Value.(T)
		i++
	}
	return result
}

func NextN(el *list.Element, index int) *list.Element {
	for (el != nil) && (index > 0) {
		index--
		el = el.Next()
	}
	return el
}
