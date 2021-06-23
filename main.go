package main

import (
	"fmt"
	"github.com/FlyingDuck/go-rookie/rsearch"
	"github.com/FlyingDuck/go-rookie/rsort"
)

func main() {
	arr := []int64{9, 10, 4, 5, 11, 13, 15, 1, 4, 4, 4}
	//rsort.BubbleSort(arr)
	//rsort.InsertionSort(arr)
	//rsort.SelectionSort(arr)
	//rsort.MergeSort(arr)
	rsort.QuickSort(arr)

	fmt.Println(arr)

	//searchIdx := rsearch.BinarySearch(arr, 4)
	//fmt.Println(searchIdx)
	//searchFirstIdx := rsearch.BinarySearchFirst(arr, 4)
	//fmt.Println(searchFirstIdx)
	//searchLastIdx := rsearch.BinarySearchLast(arr, 4)
	//fmt.Println(searchLastIdx)
	//searchFGEIdx := rsearch.BinarySearchFirstGE(arr, 8)
	//fmt.Println(searchFGEIdx)
	searchLLEIdx := rsearch.BinarySearchLastLE(arr, 9)
	fmt.Println(searchLLEIdx)
}
