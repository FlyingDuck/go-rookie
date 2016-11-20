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

快速排序



深度优先搜索

+ 问题描述：一个整数N，找到1~N的全排列
+ 样例：
    - N = 3
        - (1,2,3) (1,3,2) (2,1,3) (2,3,1) (3,1,2) (3,2,1)

        ---

        这样思考问题：N＝3，代表了有1、2、3个空位置和写有A、B、C的卡片，我们需要将卡片放到空位置上，并且每个位置只能放一张卡片，现在我们需要找出这3张卡片的所有不同摆放方法。

### 1 
`约定顺序`

首先在位置1的时候，我们手里有A、B、C三张卡片，需要考虑应该先放哪张卡片。但是既然是要找到所有的可能，所以三种可能都需要尝试,我们可以约定一个顺序 A -> B -> C。

### 2
现在可以放置卡片了:
+ A-> 1
+ B-> 2
+ C-> 3

这时我们已经得到了一种全排 (A,B,C)

### 3
+ 现在我们实际上已经走到了一个并不存在的位置4（结束位置，但是我们排列并没有结束）;
+ 现在我们需要立即回到位置3，取回卡片C，然后尝试是否还能尝试放入其它卡片（当然是按照第一步我们的约定A->B->C顺序），结果显然是我们手里并没有其它可以放入的卡片。
+ 我们需要继续回退一步来到位置2，取出卡片B，然后继续尝试是否能尝试放入其它卡片（当然是按照第一步我们的约定A->B->C顺序），这时我们发现可以放入卡片C。
+ 放入卡片后，我们向前一步来到位置3，尝试放入卡片（当然是按照第一步我们的约定A->B->C顺序），这时我们发现可以放入卡片B
+ 放完卡片，我们继续向前一步来到了并不存在的位置4
+ 我们得到了一种全排 (A, C, B)

### 4
按照上面的步骤进行下去，便会得到N=3的全排 (1,2,3) (1,3,2) (2,1,3) (2,3,1) (3,1,2) (3,2,1)

---


```
var n int = 3                        // N个卡片
var marks []int = make([]int, n)     // 用于标记N个卡片是否已经使用
var locations []int = make([]int, n) // N个位置

// step 用于记录当前所在位置,这里从0开始
    func DeepFirstSearch(step int) {
        
                if step == n { // 如果走到了位置N＋1，则表示得到了一次全排
                                    fmt.Println("N=", n, " ", locations)
                                                    return
                                                            }

                for card := 0; card < n; card++ {
                                    if 0 == marks[card] { // 如果卡片i还没有被使用
                                                                locations[step] = card    // 将卡片i放入位置 locations[step]
                                                                                        marks[card] = 1           // 标记卡片为已使用
                                                                                                                DeepFirstSearch(step + 1) // 进入下一步
                                                                                                                                        marks[card] = 0           // ** 将刚才尝试过的卡片i收回，才能进行下一次尝试 **
                                                                                                                                                        }
                                                                                                                                                                
                }
                        return

    }

```

![结果](http://upload-images.jianshu.io/upload_images/1366868-5fc2c57a5e673a3c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)



快速排序

```
func QuickSorter(array *[]int, start int, end int) {
        if start > end {
                return
        }

        pivot, head, tail := start, start, end

        for head < tail {
                // 移动尾指针 寻找小于基准值的数 （这里基准值选取的是头部数，所以先移动尾指针）
                for (*array)[tail] >= (*array)[pivot] && head < tail {
                        tail--
                }
                // 移动头指针 寻找大于基准值的数
                for (*array)[head] <= (*array)[pivot] && head < tail {
                        head++
                }

                if head < tail {
                        (*array)[head], (*array)[tail] = (*array)[tail], (*array)[head]
                } else {
                        // 将基准值归位
                        (*array)[pivot], (*array)[head] = (*array)[head], (*array)[pivot]
                }
        }

        QuickSorter(array, start, head-1)
        QuickSorter(array, head+1, end)
}
```





