# Golang listutil

Function for work with builtin list.List:
  - MergeSort
  - BubbleSort
  - MergeSortedLists

For more information look at sortlist_test.go

# Example

````
package main

import (
	"fmt"

	"github.com/mvyskoc/listutil"
)

func intIsLess(a any, b any) bool {
	return a.(int) < b.(int)
}

func main() {
	list1 := listutil.ToList([]int{7, 5, 3, 1})
	listutil.MergeSort(list1, intIsLess)

	fmt.Printf("Sorted list: %v", listutil.ToSlice[int](list1))
}
````