package leetcode

func Convert(s string, numRows int) string {
	var result = ""
	step := 2 * (numRows - 1)
	if 0 == step {
		step = 1
	}
	for row := 0; row < numRows; row++ {
		for index, length := row, len(s); index < length; index += step {
			result += string(s[index])
			if row > 0 && row < numRows-1 {
				second := index + 2*(numRows-row-1)
				if second < length {
					result += string(s[second])
				}
			}
		}
	}
	return result
}

func Convert2(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	rowSte := make([]string, numRows)
	direction := 1
	row := 0
	for index, length := 0, len(s); index < length; index++ {
		rowSte[row] += string(s[index])

		row += direction
		if row >= numRows {
			direction = -1
			row -= 2
		}

		if row < 0 {
			direction = 1
			row = 1
		}

	}

	var result = ""
	for index := 0; index < len(rowSte); index++ {
		result += rowSte[index]
	}

	return result
}
