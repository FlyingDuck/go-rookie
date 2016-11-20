package main

import (
	"fmt"
	//"dongshujin.com/leetcode"
	//"dongshujin.com/rookie"
	"dongshujin.com/algorithm"
	//"dongshujin.com/web"
)

func main() {

	//fmt.Println(leetcode.Convert2("ABCDEFGHIJL", 6))

	//fmt.Println("**Rookie Start rokcing...**")
	//fmt.Println(">> goto statement")
	//rookie.GotoFunc()
	//fmt.Println("**Rookie Completed**")

	//fmt.Println("**** Web Server ****")
	//web.RegisterHandler()
	//fmt.Println("**** End Web Server ****")

	fmt.Println("**** Algorithm ****")

	array1 := []int{9, 1, 6, 2, 0, 3, 4}

	fmt.Println("-*- BubbleSorter -*-")
	algorithm.BubbleSorter(&array1)

	fmt.Println("-*- QuickSorter -*-")
	array2 := []int{10, 8, 3, 7, 9, 2, 0}
	fmt.Println("QuickSorter Start ", array2)
	algorithm.QuickSorter(&array2, 0, len(array2)-1)
	fmt.Println("QuickSorter End ", array2)

	fmt.Println("-*- DeepFirstSearch -*-")
	fmt.Println("N Start")
	algorithm.DeepFirstSearch(0)
	fmt.Println("N End")
	fmt.Println("**** End Algorithm ****")

}
