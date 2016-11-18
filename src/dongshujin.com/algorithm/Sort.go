package algorithm

import "fmt"

func BubbleSorter(array *[]int) {
	fmt.Println("BubbleSorter Start ", *array)
	len := len(*array)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if (*array)[j] > (*array)[j+1] {
				(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
			}
		}
	}
	fmt.Println("BubbleSorter End ", *array)
}

func Traversal(array *[]int) {
	fmt.Println(&array)
}
