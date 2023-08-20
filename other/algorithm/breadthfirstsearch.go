package algorithm

import "fmt"

type node struct {
	x int // x坐标
	y int // y坐标
	f int // 父节点在队列中的编号
	s int // 步数
}

func BFS() {
	fmt.Println("BFS")

	var queue []node = make([]node, 50) // 队列

	var head, tail int // 分别标记队列头和尾

	var treasureMap = [5][4]int{ // 地图
		{0, 0, 1, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 1},
	}

	var locationMarks [5][4]int // 标记地图上哪些点已经走过

	var derections = [4][2]int{ // 4个移动方向
		{1, 0},  // 向右
		{0, 1},  // 向下
		{-1, 0}, // 向左
		{0, -1}, // 向上

	}

	var targetX, targetY int = 2, 3
	var flag bool = false

	/*--------------------- 开始寻宝 ---------------------------*/

	/*------------ 从起点开始 ---------------*/
	// 初始化队列
	head = 0
	tail = 0

	// 标示起点
	// 起点入队 这里我们默认从(X0, Y0)开始
	queue[tail].x = 0
	queue[tail].y = 0
	queue[tail].s = 0

	tail++                  // 尾指针后移一位
	locationMarks[0][0] = 1 // 标记已经走过的点

	/*------------ 走起 ---------------*/

	for head < tail { // 队列不为空

		// 枚举4个方向
		for i := 0; i < len(derections); i++ {
			// 计算下一个点坐标
			tx := queue[head].x + derections[i][0]
			ty := queue[head].y + derections[i][1]

			if tx > 3 || tx < 0 || ty > 4 || ty < 0 {
				continue
			}

			// 坐标不是障碍物 并且 没有标记过
			if treasureMap[ty][tx] == 0 && locationMarks[ty][tx] == 0 {
				locationMarks[ty][tx] = 1 // 标记已经走过

				// 将新的坐标入队
				queue[tail].x = tx
				queue[tail].y = ty
				queue[tail].f = head
				queue[tail].s = queue[head].s + 1
				tail++
			}

			// 如果已经到达藏宝点
			if tx == targetX && ty == targetY {

				flag = true
				break
			}

		}

		if flag {
			break
		} else {
			head++ // 移动头指针，遍历下一层
		}

	}

	fmt.Println("i  {x, y, f, s}")
	for i := 0; i < tail; i++ {
		fmt.Println(i, " ", queue[i])
	}

	for i := tail - 1; (queue[i].x != 0 || queue[i].y != 0) && i >= 0; {
		fmt.Println("Step ", queue[i].s, " (", queue[i].x, queue[i].y, ")")
		i = queue[i].f
	}

}
