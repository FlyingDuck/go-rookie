package main

import (
	"fmt"
)

/*
给定一个二维数组 array，请返回「螺旋遍历」该数组的结果。
螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，直到提取完所有元素。

示例 1：
输入：array = [

	[1,2,3],
	[8,9,4],
	[7,6,5]

]
输出：[1,2,3,4,5,6,7,8,9]

示例 2：
输入：array  = [

	[1,2,3,4],
	[12,13,14,5],
	[11,16,15,6],
	[10,9,8,7]

]
输出：[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16]
*/
func main() {
	array := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}
	result := spiralArray(array)
	fmt.Println(array, result)

}

func spiralArray(array [][]int) []int {
	if len(array) < 1 {
		return nil
	}

	top := 0
	bottom := len(array) - 1
	left := 0
	right := len(array[0]) - 1

	result := make([]int, 0)
	for {
		for i := left; i <= right; i++ {
			result = append(result, array[top][i])
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			result = append(result, array[i][right])
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			result = append(result, array[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		for i := bottom; i >= top; i-- {
			result = append(result, array[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return result
}
