package algorithm

import "fmt"

func Floyd() {
	// 999 表示两点之间不连通
	var theMap = [4][4]int{
		{0, 2, 6, 4},
		{999, 0, 3, 999},
		{7, 999, 0, 1},
		{5, 999, 12, 0},
	}

	fmt.Println("Floyd")

	for point := 0; point < 4; point++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if theMap[i][j] > (theMap[i][point] + theMap[point][j]) {
					theMap[i][j] = theMap[i][point] + theMap[point][j]
				}
			}
		}
		fmt.Println(" Cover point()", point, ") ", theMap)
	}
}
