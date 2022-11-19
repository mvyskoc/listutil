package listutil

import (
	"container/list"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

type sortFunc func(*list.List, LessFunc)

func compareInts(a any, b any) bool {
	return a.(int) <= b.(int)
}

func compareList(l *list.List, int_array []int) error {
	if l.Len() != len(int_array) {
		return fmt.Errorf("length of list %d, but expected %d", l.Len(), len(int_array))
	}

	i := 0
	for el := l.Front(); el != nil; el = el.Next() {
		if el.Value.(int) != int_array[i] {
			return fmt.Errorf("unexpected element value (list[%d] = %d) != %d", i, el.Value.(int), int_array[i])
		}
		i++
	}

	return nil
}

func testSortList(t *testing.T, sortlist sortFunc, int_array []int) {
	sorted_list := ToList(int_array)
	sort.Ints(int_array)

	sortlist(sorted_list, compareInts)
	err := compareList(sorted_list, int_array)
	if err != nil {
		if len(int_array) < 20 {
			t.Errorf("%v\n"+
				"  sorted list:    %v\n"+
				"  expected value: %v",
				err, ToSlice[int](sorted_list), int_array)
		} else {
			t.Error(err)
		}
	}
}

func TestMergeSort(t *testing.T) {
	testSortList(t, MergeSort, []int{})
	testSortList(t, MergeSort, []int{1})
	testSortList(t, MergeSort, []int{2, 1})
	testSortList(t, MergeSort, []int{3, 1, 0})
	testSortList(t, MergeSort, []int{5, 1, 3, 0, 4})
	testSortList(t, MergeSort, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	rand_arr := rand.Perm(10000)
	t.Log("Start long array sort")
	testSortList(t, MergeSort, rand_arr)
	t.Log("Stop long array sort")
}

func TestBubbleSort(t *testing.T) {
	testSortList(t, BubbleSort, []int{})
	testSortList(t, BubbleSort, []int{1})
	testSortList(t, BubbleSort, []int{2, 1})
	testSortList(t, BubbleSort, []int{3, 1, 0})
	testSortList(t, BubbleSort, []int{5, 1, 3, 0, 4})
	testSortList(t, BubbleSort, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	rand_arr := rand.Perm(10000)
	t.Log("Start long array sort")
	testSortList(t, BubbleSort, rand_arr)
	t.Log("Stop long array sort")
}

func checkGetMiddle(t *testing.T, elements []int, expected int) {
	base_list := ToList(elements)
	lw := ListView{Front: base_list.Front(), Back: base_list.Back()}
	middle_element := getMiddle(lw)
	if len(elements) == 0 {
		if middle_element != nil {
			t.Errorf("Return element for empty list")
		}
	} else if middle_element == nil {
		t.Errorf("Return nil element, expected %v\n"+
			"  elements = %v", expected, elements)
	} else {
		if middle_element.Value.(int) != expected {
			t.Errorf("Return %v, expected %v\n"+
				"  elements = %v", middle_element.Value.(int), expected, elements)
		}
	}

}

func TestGetMiddle(t *testing.T) {
	checkGetMiddle(t, []int{}, 0)
	checkGetMiddle(t, []int{1}, 1)
	checkGetMiddle(t, []int{1, 2}, 1)
	checkGetMiddle(t, []int{1, 2, 3}, 2)
	checkGetMiddle(t, []int{1, 2, 3, 4}, 2)
	checkGetMiddle(t, []int{1, 2, 3, 4, 5}, 3)
	checkGetMiddle(t, []int{2, 5, 6, 7, 1, 3, 4, 8}, 7)
}

func TestGetNextN(t *testing.T) {
	l := ToList([]int{0, 1, 2, 3, 4, 5, 6})

	if el := NextN(l.Front(), 0); el.Value.(int) != 0 {
		t.Errorf("Expected 0 element, GetNextN(0) = %v", el.Value)
	}
	if el := NextN(l.Front(), 2); el.Value.(int) != 2 {
		t.Errorf("Expected 2 element, GetNextN(0) = %v", el.Value)
	}
	if el := NextN(l.Front(), 100); el != nil {
		t.Errorf("Expected nil element, GetNextN(100) = %v", el.Value)
	}
}

// arr1, arr2 are sorted slices
func testSortMergeView(t *testing.T, int_array1, int_array2 []int) {
	var (
		sorted_list  *list.List
		sorted_view1 = ListView{nil, nil}
		sorted_view2 = ListView{nil, nil}
	)

	if len(int_array1) > 0 {
		sorted_list = ToList(int_array1)
		sorted_view1 = ListView{sorted_list.Front(), sorted_list.Back()}
	} else {
		sorted_list = ToList(int_array2)
		sorted_view2 = ListView{sorted_list.Front(), sorted_list.Back()}
	}

	if (len(int_array1) > 0) && (len(int_array2) > 0) {
		sorted_list.PushBackList(ToList(int_array2))
		sorted_view2 = ListView{sorted_view1.Back.Next(), sorted_list.Back()}
	}

	merge_view := mergeSortedView(sorted_list, sorted_view1, sorted_view2, compareInts)

	int_merge_expected := make([]int, 0, len(int_array1)+len(int_array2))
	int_merge_expected = append(int_merge_expected, int_array1...)
	int_merge_expected = append(int_merge_expected, int_array2...)
	sort.Ints(int_merge_expected)

	int_merged := ListViewToSlice[int](merge_view)
	if !reflect.DeepEqual(int_merge_expected, int_merged) {
		if len(int_merge_expected) < 10 {
			t.Errorf("merge view does not match\n"+
				"  list1:     %v\n"+
				"  list2:     %v\n"+
				"  merged:  %v\n"+
				"  expected:  %v", int_array1, int_array2, int_merged, int_merge_expected)
		} else {
			t.Error("merge view does not match")
		}
	} else {
		if len(int_merge_expected) < 10 {
			t.Logf("Merge view OK:\n"+
				"  list1:     %v\n"+
				"  list2:     %v\n"+
				"  merged:  %v", int_array1, int_array2, int_merged)
		}
	}

}

func TestSortMergeView(t *testing.T) {
	testSortMergeView(t, []int{1}, []int{})
	testSortMergeView(t, []int{}, []int{1})
	testSortMergeView(t, []int{1}, []int{2})
	testSortMergeView(t, []int{2}, []int{1})
	testSortMergeView(t, []int{1, 5}, []int{3})
	testSortMergeView(t, []int{3}, []int{1, 5})
	testSortMergeView(t, []int{2, 4}, []int{1, 3})
	testSortMergeView(t, []int{1, 3}, []int{2, 4})
}

func testSortMergeLists(t *testing.T, int_array1, int_array2 []int) {
	merge_list := ToList(int_array1)
	int_list2 := ToList(int_array2)
	expected_merge := make([]int, 0, len(int_array1)+len(int_array2))
	expected_merge = append(expected_merge, int_array1...)
	expected_merge = append(expected_merge, int_array2...)
	sort.Ints(expected_merge)

	MergeSortedLists(merge_list, int_list2, compareInts)
	err := compareList(merge_list, expected_merge)
	if err != nil {
		t.Errorf("%v\n"+
			"  list1:    %v\n"+
			"  list2:    %v\n"+
			"  merged:   %v\n"+
			"  expected: %v",
			err, int_array1, int_array2, ToSlice[int](merge_list), expected_merge)
	} else {
		t.Logf("Mergelist OK\n"+
			"  list1:    %v\n"+
			"  list2:    %v\n"+
			"  merged:   %v\n"+
			"  expected: %v",
			int_array1, int_array2, ToSlice[int](merge_list), expected_merge)
	}

}

func TestMergeSortedLists(t *testing.T) {
	testSortMergeLists(t, []int{1}, []int{})
	testSortMergeLists(t, []int{}, []int{1})
	testSortMergeLists(t, []int{1}, []int{2})
	testSortMergeLists(t, []int{2}, []int{1})
	testSortMergeLists(t, []int{1, 5}, []int{3})
	testSortMergeLists(t, []int{3}, []int{1, 5})
	testSortMergeLists(t, []int{2, 4}, []int{1, 3})
	testSortMergeLists(t, []int{1, 3}, []int{2, 4})
}

func BenchmarkMergeSort(b *testing.B) {
	sorted_list := list.New()
	for i := 0; i < b.N; i++ {
		sorted_list.PushBack(b.N - i)
	}
	MergeSort(sorted_list, compareInts)

}

func BenchmarkBubbleSort(b *testing.B) {
	sorted_list := list.New()
	for i := 0; i < b.N; i++ {
		sorted_list.PushBack(b.N - i)
	}
	BubbleSort(sorted_list, compareInts)

}

func BenchmarkSliceSort(b *testing.B) {
	sorted_int := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		sorted_int[i] = b.N - i
	}
	sort.Ints(sorted_int)
}
