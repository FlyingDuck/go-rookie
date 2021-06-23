package rsort

import "fmt"

func QuickSort(arr []int64) {
	fmt.Println(arr)
	if len(arr) <= 1 {
		return
	}

	quickSortSubsequence(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func quickSortSubsequence(arr []int64, start, end int) {
	if end <= start {
		return
	}

	pivot := partition(arr, start, end)

	quickSortSubsequence(arr, start, pivot-1)
	quickSortSubsequence(arr, pivot+1, end)
}

func partition(arr []int64, start, end int) int {
	pivotVal := arr[end]
	pivot := start - 1

	for i := start; i < end; i++ {
		if arr[i] < pivotVal {
			arr[pivot+1], arr[i] = arr[i], arr[pivot+1]
			pivot++
		}
	}
	arr[pivot+1], arr[end] = arr[end], arr[pivot+1]
	return pivot + 1
}
