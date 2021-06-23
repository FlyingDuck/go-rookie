package rsearch

func BinarySearch(arr []int64, targetVal int64) int {
	if len(arr) == 0 {
		return -1
	}

	for startIdx, endIdx := 0, len(arr)-1; startIdx <= endIdx; {
		midIdx := startIdx + ((endIdx - startIdx) >> 1) // + 的优先级 高于 >>
		if arr[midIdx] > targetVal {
			endIdx = midIdx - 1
		} else if arr[midIdx] < targetVal {
			startIdx = midIdx + 1
		} else {
			return midIdx
		}
	}
	return -1
}

// 查找第一个值等于给定值的元素
func BinarySearchFirst(arr []int64, targetVal int64) int {
	if len(arr) == 0 {
		return -1
	}

	for startIdx, endIdx := 0, len(arr)-1; startIdx <= endIdx; {
		midIdx := startIdx + ((endIdx - startIdx) >> 1) // + 的优先级 高于 >>
		if arr[midIdx] > targetVal {
			endIdx = midIdx - 1
		} else if arr[midIdx] < targetVal {
			startIdx = midIdx + 1
		} else {
			if midIdx == 0 || arr[midIdx-1] != targetVal {
				return midIdx
			}
			endIdx = midIdx - 1
		}
	}
	return -1
}

// 查找最后一个值等于给定值的元素
func BinarySearchLast(arr []int64, targetVal int64) int {
	if len(arr) == 0 {
		return -1
	}

	for startIdx, endIdx := 0, len(arr)-1; startIdx <= endIdx; {
		midIdx := startIdx + ((endIdx - startIdx) >> 1) // + 的优先级 高于 >>
		if arr[midIdx] > targetVal {
			endIdx = midIdx - 1
		} else if arr[midIdx] < targetVal {
			startIdx = midIdx + 1
		} else {
			if midIdx == len(arr)-1 || arr[midIdx+1] != targetVal {
				return midIdx
			}
			startIdx = midIdx + 1
		}
	}
	return -1
}

// 查找第一个大于等于给定值的元素
func BinarySearchFirstGE(arr []int64, targetVal int64) int {
	if len(arr) == 0 {
		return -1
	}

	for startIdx, endIdx := 0, len(arr)-1; startIdx <= endIdx; {
		midIdx := startIdx + ((endIdx - startIdx) >> 1) // + 的优先级 高于 >>
		if arr[midIdx] >= targetVal {
			if midIdx == 0 || arr[midIdx-1] < targetVal {
				return midIdx
			}
			endIdx = midIdx - 1
		} else {
			startIdx = midIdx + 1
		}
	}
	return -1
}

// 查找最后一个小于等于给定值的元素
func BinarySearchLastLE(arr []int64, targetVal int64) int {
	if len(arr) == 0 {
		return -1
	}

	for startIdx, endIdx := 0, len(arr)-1; startIdx <= endIdx; {
		midIdx := startIdx + ((endIdx - startIdx) >> 1) // + 的优先级 高于 >>
		if arr[midIdx] > targetVal {
			endIdx = midIdx - 1
		} else {
			if midIdx == len(arr)-1 || arr[midIdx+1] > targetVal {
				return midIdx
			}
			startIdx = midIdx + 1
		}
	}
	return -1
}
