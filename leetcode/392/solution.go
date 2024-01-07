package main

import "fmt"

/*
392. 判断子序列

给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

示例 1：
输入：s = "abc", t = "ahbgdc"
输出：true

示例 2：
输入：s = "axc", t = "ahbgdc"
输出：false
*/
func main() {
	s := "abc"
	t := "ahbgdc"
	result := isSubsequence(s, t)
	fmt.Println(result)
}

/*
思路及算法

本题询问的是，sss 是否是 ttt 的子序列，因此只要能找到任意一种 sss 在 ttt 中出现的方式，即可认为 sss 是 ttt 的子序列。
而当我们从前往后匹配，可以发现每次贪心地匹配靠前的字符是最优决策。

这样，我们初始化两个指针 iii 和 jjj，分别指向 sss 和 ttt 的初始位置。每次贪心地匹配，匹配成功则 iii 和 jjj 同时右移，匹配 sss 的下一个位置，匹配失败则 jjj 右移，iii 不变，尝试用 ttt 的下一个字符匹配 sss。
最终如果 iii 移动到 sss 的末尾，就说明 sss 是 ttt 的子序列。
*/
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == len(s)
}
