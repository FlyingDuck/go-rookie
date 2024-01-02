package main

import (
	"fmt"
	"github.com/FlyingDuck/go-rookie/leetcode/util"
)

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
*/
func main() {
	nums := []int{2, 7, 9, 3, 1}
	result := rob(nums)

	fmt.Println(nums, result)

	result, routine := robAndRecord(nums)
	fmt.Println(nums, result, routine)

}

// 问题定义：
// dp[i]：表示打劫前 i 家房屋，所获得的现金的最高金额
//
// 问题分析：
// 还是顶层思维，解决不确定问题。当打劫到第 i 家房屋时，存在不确定性的情况是：打劫 或 不打劫：
// - 打劫：根据题目设定，那么第 i-1 家不能打劫，最大金额为：dp[i-2] + nums[i-1]；
// - 不打劫：那么最大金额为打劫前 i-1 家的最大金额：dp[i-1]；
//
// 子问题：
// dp[0] = 0, dp[1] = nums[0], dp[2] = max(nums[0], nums[1])
func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = util.MaxInt(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}

// 计算金额并记录路径
func robAndRecord(nums []int) (int, []int) {
	type record struct {
		val    int
		robbed bool
	}

	dp := make([]record, len(nums)+1)
	dp[0] = record{
		val:    0,
		robbed: false,
	}
	dp[1] = record{
		val:    nums[0],
		robbed: true,
	}
	for i := 2; i <= len(nums); i++ {
		maxVal, first := util.MaxIntEnhanced(dp[i-2].val+nums[i-1], dp[i-1].val)
		dp[i] = record{
			val:    maxVal,
			robbed: first,
		}
	}

	routine := make([]int, 0, len(nums))
	for i := len(nums); i > 0; {
		if dp[i].robbed {
			routine = append(routine, i-1)
			i = i - 2
		} else {
			i--
		}
	}

	return dp[len(nums)].val, routine
}
