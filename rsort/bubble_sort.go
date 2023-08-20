package rsort

import "fmt"

func BubbleSort(arr []int64) {
	loopN := 0
	fmt.Println(arr)
	if len(arr) <= 1 {
		return
	}

	for i := range arr {
		fmt.Println("i=", i)
		swap := false
		for j := i+1; j<len(arr); j++ {
			loopN++
			if arr[i] > arr[j] {
				fmt.Printf("%d>%d, swap %d <-> %d\n", arr[i], arr[j], i, j)
				arr[i], arr[j] = arr[j], arr[i]
				swap = true
			}
		}
		// 某次冒泡操作如果一次元素都没有发生过交换，说明 i 之后的元素已经是有序的了，就不需要再进行冒泡操作
		if !swap {
			fmt.Println("break")
			break
		}
	}
	fmt.Println(arr)
	fmt.Println("loopN=", loopN)
}
