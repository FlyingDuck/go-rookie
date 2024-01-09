package main

import "fmt"

/*
504. 七进制数

给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。

示例 1:
输入: num = 100
输出: "202"

示例 2:
输入: num = -7
输出: "-10"
*/
func main() {
	result := convertToBase7(100)
	fmt.Println(result)
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	numbers := []byte{'0', '1', '2', '3', '4', '5', '6'}

	flag := false
	if num < 0 {
		flag = true
		num = num * -1
	}
	results := make([]byte, 0)
	for num != 0 {
		mod := num % 7
		results = append(results, numbers[mod])
		num = num / 7
	}
	if flag {
		results = append(results, '-')
	}

	for i, j := 0, len(results)-1; i < j; {
		results[i], results[j] = results[j], results[i]
		i++
		j--
	}
	return string(results)
}
