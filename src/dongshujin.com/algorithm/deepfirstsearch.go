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
