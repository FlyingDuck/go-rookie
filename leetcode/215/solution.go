package main

import "fmt"

/*
215. 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4
*/
func main() {
	nums := []int{3, 2, 1, 5, 6, 4, 7}
	k := 2
	result := findKthLargest(nums, k)
	fmt.Println(result)
}

func findKthLargest(nums []int, k int) int {

	left, right := 0, len(nums)-1
	for {
		p := partition(nums, left, right)
		if p > k-1 {
			right = p - 1
		} else if p < k-1 {
			left = p + 1
		} else {
			return nums[p]
		}
	}
}

//func swap(nums []int, i, j int) {
//	tmp := nums[i]
//	nums[i] = nums[j]
//	nums[j] = tmp
//}

func partition(nums []int, startIdx, endIdx int) (idx int) {
	pivot := startIdx - 1
	pivotVal := nums[endIdx]
	for i := startIdx; i < endIdx; i++ {
		if nums[i] > pivotVal {
			pivot++
			nums[i], nums[pivot] = nums[pivot], nums[i]
		}
	}

	nums[pivot+1], nums[endIdx] = nums[endIdx], nums[pivot+1]
	return pivot + 1
}
