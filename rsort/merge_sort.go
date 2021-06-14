package rsort

import "fmt"

func MergeSort(arr []int64) {
	fmt.Println(arr)
	if len(arr) <= 1 {
		return
	}

	mergeSubsequence(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func mergeSubsequence(arr []int64, startIdx, endIdx int) {
	if startIdx >= endIdx {
		return
	}

	midIdx := startIdx + (endIdx-startIdx)/2 // (startIdx+endIdx)/2
	mergeSubsequence(arr, startIdx, midIdx)
	mergeSubsequence(arr, midIdx+1, endIdx)
	merge(arr, startIdx, midIdx, endIdx)
}

func merge(arr []int64, startIdx, midIdx, endIdx int) {
	var tmpSubArr []int64

	i, j := startIdx, midIdx+1
	for i <= midIdx && j <= endIdx {
		if arr[i] < arr[j] {
			tmpSubArr = append(tmpSubArr, arr[i])
			i++
		} else {
			tmpSubArr = append(tmpSubArr, arr[j])
			j++
		}
	}

	if i > midIdx {
		tmpSubArr = append(tmpSubArr, arr[j:endIdx+1]...)
	}
	if j > endIdx {
		tmpSubArr = append(tmpSubArr, arr[i:midIdx+1]...)
	}

	for i, j := 0, startIdx; i < len(tmpSubArr); {
		arr[j] = tmpSubArr[i]
		i++
		j++
	}
}
