package algorithm

import "fmt"

func BubbleSorter(array *[]int) {
	fmt.Println("BubbleSorter Start ", *array)
	len := len(*array)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if (*array)[j] > (*array)[j+1] {
				(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
			}
		}
	}
	fmt.Println("BubbleSorter End ", *array)
}

func QuickSorter(array *[]int, start int, end int) {
	if start > end {
		return
	}

	pivot, head, tail := start, start, end

	for head < tail {
		// 移动尾指针 寻找小于基准值的数 （这里基准值选取的是头部数，所以先移动尾指针）
		for (*array)[tail] >= (*array)[pivot] && head < tail {
			tail--
		}
		// 移动头指针 寻找大于基准值的数
		for (*array)[head] <= (*array)[pivot] && head < tail {
			head++
		}

		if head < tail {
			(*array)[head], (*array)[tail] = (*array)[tail], (*array)[head]
		} else {
			// 将基准值归位
			(*array)[pivot], (*array)[head] = (*array)[head], (*array)[pivot]
		}
	}

	QuickSorter(array, start, head-1)
	QuickSorter(array, head+1, end)
}
