package leetcode

import "fmt"

func twoSum(nums []int, target int) []int {
	for former, formerV := range nums {
		for latter, count := former+1, len(nums); latter < count; latter++ {
			latterV := nums[latter]
			if formerV+latterV == target {
				return []int{former, latter}
			}
		}
	}
	return []int{0, 0}
}

func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, count := 0, len(nums); index < count; index++ {
		firstV := nums[index]
		secondV := target - firstV
		second, ok := numMap[secondV]
		if ok {
			return []int{second, index}
		}
		numMap[firstV] = index
	}
	return []int{0, 0}
}

func RunTwoSum() {
	nums := []int{3, 2, 4}
	target := 6

	result := twoSum(nums, target)
	fmt.Println(result)

	result = twoSum2(nums, target)
	fmt.Println(result)
}
