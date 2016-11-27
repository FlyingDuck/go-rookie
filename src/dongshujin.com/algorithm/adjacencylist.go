package algorithm

import "fmt"

// 顶点编号：0，1，2，3 表示顶点 A，B，C，D
// 边编号：0，1，2，3，4
var u = [5]int{0, 3, 0, 1, 0}
var v = [5]int{3, 2, 1, 3, 2}
var w = [5]int{9, 8, 5, 6, 7}

func AdjacencyList() ([4]int, [5]int) {

	for i := 0; i < 5; i++ {
		fmt.Println(u[i], " -> ", v[i], " ", w[i])
	}

	first := [4]int{-1, -1, -1, -1}
	next := [5]int{-1, -1, -1, -1, -1}

	fmt.Println("fisrt : ", first)
	fmt.Println("next: ", next)
	fmt.Println("-----------------------")

	// 遍历5条边
	for i := 0; i < 5; i++ {
		next[i] = first[u[i]] // 将第i条边的下一条边设置为 u[i]号顶点的当前第一条边
		first[u[i]] = i       // 将u[i]号顶点的第一条边设置为当前边
	}

	fmt.Println("first: ", first)
	fmt.Println("next: ", next)

	return first, next
}

func TraverseEdge(point int, first [4]int, next [5]int) {
	k := first[point]
	for -1 != k {
		fmt.Println(u[k], " -> ", v[k], " ", w[k])
		k = next[k]
	}
}
