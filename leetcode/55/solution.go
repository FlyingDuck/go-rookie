package main

import (
	"fmt"
)

/*
55. 跳跃游戏

给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
*/
func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	result := canJump(nums)
	fmt.Println(result)
}

// 直接遍历所有点，如果当前点比前面所有点可以达到的点的最大距离还要大就返回false，如果能遍历到最后一个点，说明最后一个点可以到达，返回true
func canJump(nums []int) bool {
	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}
		k = max(i+nums[i], k)
	}
	return true
}

//func canJump(nums []int) bool {
//	dp := make([]int, len(nums))
//	dp[0] = 1
//
//	for i := 1; i < len(nums); i++ {
//		for j := 0; j < i; j++ {
//			if nums[j]+j >= i {
//				dp[i] += dp[j]
//			}
//		}
//	}
//
//	return dp[len(nums)-1] > 0
//}
