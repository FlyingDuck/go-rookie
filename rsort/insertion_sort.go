package rsort

import "fmt"

func InsertionSort(arr []int64) {
	fmt.Println(arr)
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		fmt.Println("i=", i)
		value := arr[i]
		j := i-1
		for ; j >= 0; j-- {
			if value < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
	fmt.Println(arr)
}
