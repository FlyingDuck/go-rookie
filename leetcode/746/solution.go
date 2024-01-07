package main

import "fmt"

/*
746. 使用最小花费爬楼梯

给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。
*/
func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	result := minCostClimbingStairs(cost)
	fmt.Println(result)
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1) // 爬到第 i 个台阶时的最低花费
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i], dp[i-2]+cost[i])
	}
	return min(dp[len(cost)-1], dp[len(cost)-2])
}

//func min(num1, num2 int) int {
//	if num1 > num2 {
//		return num2
//	}
//	return num1
//}
