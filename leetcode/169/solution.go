package main

import "fmt"

/*
169, 多数元素
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	result := majorityElement(nums)
	fmt.Println(result)
}

func majorityElement(nums []int) int {
	majority := nums[0]
	votes := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == majority {
			votes++
		} else {
			votes--
			if votes <= 0 {
				majority = nums[i]
				votes = 1
			}
		}
	}
	return majority
}
