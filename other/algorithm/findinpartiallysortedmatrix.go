package algorithm

/*
在一个二维数组中，每行都是按照从左往右递增的顺序排序，每列也是都是按照从上到下递增的顺序排序的。现从这样一个数组中查找一个整数。
*/

func FindInpartiallySortedMatrix(array [][]int, target int) bool {
	found := false

	if nil != array {
		totalRows := len(array)
		totalCols := len(array[0])

		// 我们从二维数组的右上角上的数开始查找
		row := 0
		col := totalCols - 1

		for row < totalRows && col >= 0 {
			if array[row][col] > target { // 当前数大于目标数，那么这一列其他数必定也大于目标数，故剔除整列
				col--
			} else if array[row][col] < target { // 当前数小于目标数，那么这一行其他数必定也小于目标数，故剔除整行
				row++
			} else {
				found = true
				break
			}
		}
	}

	return found
}
