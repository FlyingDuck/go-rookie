package other

/*
环节点的走法数

题目描述
描述信息
一个环上有10个点，编号为0-9，从0点出发，每步可以顺时针到下一个点，也可以逆时针到上一个点，
求：经过 n 步又回到0点有多少种不同的走法？

举例：
如果n = 1，则从0出发只能到1或者9，不可能回到0，共0种走法
如果n = 2，则从0出发有4条路径：0->1->2, 0->1->0, 0->9->8, 0->9->0，其中有两条回到了0点，故一共有2种走法。
*/

// solution: https://leetcode.cn/circle/discuss/TWO4Z5/
// 动态规划
// 定义问题：走 i 步回到原点，第 i-1 步存在两种情况，在位置 1 或者在 位置 9，那么第 i 步回到原点的走法为：
// 第 i-1 步在位置 1 的走法 + 第 i-1 步在位置 9 的走法
//
// 问题分析
// dp[i][j] 为走 i 步在位置 j 的走法数量，那么 dp[i][j] = dp[i-1][(j+1)%length] + dp[i-1][(j-1+length)%length]
//
// 子问题
// dp[0][0] = 1,
func stepBack2Zero(length int, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, length)
	}
	dp[0][0] = 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j < length; j++ {
			dp[i][j] = dp[i-1][(j+1)%length] + dp[i-1][(j-1+length)%length]
		}
	}

	return dp[n][0]
}
