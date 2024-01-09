package main

import "fmt"

/*
6. Z 字形变换
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
*/
func main() {
	//s := "PAYPALISHIRING"
	//numRows := 3
	//result := convert(s, numRows)
	//fmt.Println(result)

	s := "AB"
	numRows := 1
	result := convert(s, numRows)
	fmt.Println(result)
}

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	rows := make([][]byte, numRows)
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]byte, 0)
	}

	flag := -1
	rowIdx := 0
	for i := 0; i < len(s); i++ {
		rows[rowIdx] = append(rows[rowIdx], s[i])
		if rowIdx == 0 || rowIdx == len(rows)-1 {
			flag = -flag
		}
		rowIdx = rowIdx + flag
	}
	result := ""
	for _, row := range rows {
		result = result + string(row)
	}
	return result
}
