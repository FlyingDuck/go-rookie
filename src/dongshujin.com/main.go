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

	array := []int{9, 1, 6, 2, 0, 3, 4}

	algorithm.BubbleSorter(&array)

	fmt.Println("**** End Algorithm ****")

}
