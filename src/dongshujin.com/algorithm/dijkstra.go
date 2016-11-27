package algorithm

import "fmt"

func Dijkstra() {
	// 999表示顶点之间没有连通
	var theMap = [6][6]int{
		{0, 1, 12, 999, 999, 999},
		{999, 0, 9, 3, 999, 999},
		{999, 999, 0, 999, 5, 999},
		{999, 999, 4, 0, 13, 15},
		{999, 999, 999, 999, 0, 4},
		{999, 999, 999, 999, 999, 0},
	}

	var marks = [6]int{1, 0, 0, 0, 0, 0} // 1，表示该顶点最短路径为确定值；0，表示该顶点的最短路径为估计值

	var dis [6]int

	// 初始化A顶点的最小路径数组
	for i := 0; i < 6; i++ {
		dis[i] = theMap[0][i]
	}

	fmt.Println("Dijkstra")
	// Dijkstra
	for i := 0; i < 5; i++ { // 这里为6个顶点，所以总共要进行5次 “松弛”
		minDistance := 1000 // 记录一次松弛中“估计值”中的最小距离
		currentPoint := 0   // 记录一次松弛中“估计值”中的顶点
		// 遍历最短距离数组，找到“估计值”中距离A顶点最近的顶点
		for j := 0; j < len(dis); j++ { //
			if marks[j] == 0 && minDistance > dis[j] {
				minDistance = dis[j]
				currentPoint = j
			}
		}

		marks[currentPoint] = 1 // 标记最小“估计值”为“确认值”

		// 遍历该顶点的出边并进行松弛
		for k := 0; k < 6; k++ {
			if theMap[currentPoint][k] < 999 && dis[k] > (dis[currentPoint]+theMap[currentPoint][k]) {
				dis[k] = dis[currentPoint] + theMap[currentPoint][k]
			}
		}

		fmt.Println("Step ", i, " ", dis)

	}

	fmt.Println("Dijkstra A -> ... ", dis)

}
