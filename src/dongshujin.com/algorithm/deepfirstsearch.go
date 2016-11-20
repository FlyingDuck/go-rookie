package algorithm

import (
	"fmt"
)

/*
问题描述：
输入一个数n，输出1～n的全排列？
*/

var n int = 3                        // N个卡片
var marks []int = make([]int, n)     // 用于标记N个卡片是否已经使用
var locations []int = make([]int, n) // N个位置
var sum int = 0

func ResetCondition(num int) {
	n = num
	marks = make([]int, n)
	locations = make([]int, n)
	sum = 0
}

// step 用于记录当前所在位置,这里从0开始
func DeepFirstSearch(step int) {

	if step == n { // 如果走到了位置N＋1，则表示得到了一次全排
		fmt.Println("N=", n, " ", locations)
		return
	}

	for card := 0; card < n; card++ {
		if 0 == marks[card] { // 如果卡片i还没有被使用
			locations[step] = card    // 将卡片i放入位置 locations[step]
			marks[card] = 1           // 标记卡片为已使用
			DeepFirstSearch(step + 1) // 进入下一步
			marks[card] = 0           // ** 将刚才尝试过的卡片i收回，才能进行下一次尝试 **
		}
	}
	return
}

func DeepFirstSearch2(step int) {
	if step == n {
		if ((locations[0]+locations[3])*100 + (locations[1]+locations[4])*10 + locations[2] + locations[5]) == (locations[6]*100 + locations[7]*10 + locations[8]) {
			fmt.Println(locations[0:3], " + ", locations[3:6], " = ", locations[6:9])
		}
		return
	}

	for i := 0; i < n; i++ {
		if 0 == marks[i] {
			locations[step] = i
			marks[i] = 1
			DeepFirstSearch2(step + 1)
			marks[i] = 0
		}
	}
	return
}

var treasureMap = [5][4]int{
	{0, 0, 1, 0},
	{0, 0, 0, 0},
	{0, 0, 1, 0},
	{0, 1, 0, 0},
	{0, 0, 0, 1},
}

//var locationMarks = [5][5]int{
//	{1, 0, 0, 0},
//	{0, 0, 0, 0},
//	{0, 0, 0, 0},
//	{0, 0, 0, 0},
//	{0, 0, 0, 0},
//}

var locationMarks [5][4]int

var dereoctions = [4][2]int{
	{1, 0},  // 向右
	{0, 1},  // 向下
	{-1, 0}, // 向左
	{0, -1}, // 向上
}

var minDistance int = 20
var targetX, targetY int = 2, 3

func DeepFirstSearch3(x, y, step int) {
	locationMarks[y][x] = 1

	if targetX == x && targetY == y {
		if step < minDistance {
			minDistance = step
			fmt.Println("Now distance is ", minDistance)
			for i := 0; i < 5; i++ {
				fmt.Println(locationMarks[i])
			}
		}
		return
	}

	for d := 0; d < 4; d++ { // 遍历下一步的方向
		// 计算下一个坐标
		tx := x + dereoctions[d][0]
		ty := y + dereoctions[d][1]

		// 查看是否越界
		if tx < 0 || tx > 3 || ty < 0 || ty > 4 {
			continue
		}

		// 查看下一个点是否是障碍物 或者 已经经过
		if treasureMap[ty][tx] == 0 && locationMarks[ty][tx] == 0 {
			//locationMarks[ty][tx] = 1
			DeepFirstSearch3(tx, ty, step+1)
			locationMarks[ty][tx] = 0
		}
	}
	return
}
