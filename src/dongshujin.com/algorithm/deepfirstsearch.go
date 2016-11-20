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
