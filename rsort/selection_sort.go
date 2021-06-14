package rsort

import "fmt"

func SelectionSort(arr []int64) {
	fmt.Println(arr)
	if len(arr) <= 1 {
		return
	}

	for i := range arr {
		min := arr[i]
		minIdx := i
		for j:=i+1; j<len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	fmt.Println(arr)
}
