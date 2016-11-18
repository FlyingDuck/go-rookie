## Learning Golang & Algorithms From [LeetCode](https://leetcode.com/)

NO. [1 Two Sum](https://leetcode.com/problems/two-sum/)
NO. [6 ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)


## Algorithm rookie

冒泡排序

```
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
```





