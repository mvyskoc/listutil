package listutil

import (
	"container/list"
)

type LessFunc func(a any, b any) bool

func BubbleSort(l *list.List, IsLess LessFunc) {
	var before_element *list.Element

	for i := l.Front(); i != nil; {
		unsorted := i.Next()
		before_element = nil
		for j := i.Prev(); j != nil; j = j.Prev() {

			if IsLess(i.Value, j.Value) {
				before_element = j
			} else {
				break
			}
		}
		if before_element != nil {
			l.MoveBefore(i, before_element)
		}
		i = unsorted
	}
}

func MergeSort(l *list.List, IsLess LessFunc) {
	mergeSort(l, ListView{l.Front(), l.Back()}, IsLess)
}

// Merge list src into the list dst
// Both lists src, dst must be already sorted
func MergeSortedLists(dst, src *list.List, IsLess LessFunc) {
	if dst.Len() == 0 {
		dst.PushBackList(src)
	} else if src.Len() > 0 {
		lw1 := ListView{dst.Front(), dst.Back()}
		dst.PushBackList(src)
		lw2 := ListView{lw1.Back.Next(), dst.Back()}
		mergeSortedView(dst, lw1, lw2, IsLess)
	}
}

func mergeSort(l *list.List, lw ListView, IsLess LessFunc) ListView {
	if (lw.Front == nil) || (lw.Next(lw.Front) == nil) {
		return ListView{Front: lw.Front, Back: lw.Back}
	}

	mid_element := getMiddle(lw)
	second_start := lw.Next(mid_element)
	sorted_first := mergeSort(l, ListView{lw.Front, mid_element}, IsLess)
	sorted_second := mergeSort(l, ListView{second_start, lw.Back}, IsLess)
	result := mergeSortedView(l, sorted_first, sorted_second, IsLess)
	return result
}

// Return middle element of the list
func getMiddle(lw ListView) *list.Element {
	slow := lw.Front
	if slow == nil {
		return nil
	}
	fast := lw.Next(slow)
	for fast != nil {
		fast = lw.Next(fast)
		if fast != nil {
			slow = lw.Next(slow)
			fast = lw.Next(fast)
		}
	}
	return slow
}

func mergeSortedView(l *list.List, lw1, lw2 ListView, IsLess LessFunc) ListView {
	if (lw1.Front == nil) && (lw2.Front == nil) {
		return ListView{nil, nil}
	}
	if lw1.Front == nil {
		return lw2
	}
	if lw2.Front == nil {
		return lw1
	}

	if !IsLess(lw1.Front.Value, lw2.Front.Value) {
		lw1, lw2 = lw2, lw1
	}

	back := lw1.Front
	e1 := lw1.Next(back)
	e2 := lw2.Front
	for (e1 != nil) && (e2 != nil) {
		if IsLess(e1.Value, e2.Value) {
			back = e1
			e1 = lw1.Next(e1)
		} else {
			new_e2 := lw2.Next(e2)
			l.MoveAfter(e2, back)
			back = e2
			e2 = new_e2
		}
	}

	if e2 != nil {
		for e2 != nil {
			new_e2 := lw2.Next(e2)
			l.MoveAfter(e2, back)
			back = e2
			e2 = new_e2
		}
		lw1.Back = back
	}
	return lw1
}
