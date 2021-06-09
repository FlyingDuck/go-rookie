package main

import "github.com/FlyingDuck/go-rookie/rsort"

func main() {
	arr := []int64{9, 10, 4, 5, 11, 13, 15}
	//rsort.BubbleSort(arr)
	rsort.InsertionSort(arr)
}
