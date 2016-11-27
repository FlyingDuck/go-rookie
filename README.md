## Learning Golang & Algorithms From [LeetCode](https://leetcode.com/)

NO. [1 Two Sum](https://leetcode.com/problems/two-sum/)  
NO. [6 ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)


## Algorithm rookie

### 冒泡排序

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

---

### 快速排序

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

---

# 深度优先搜索

+ 问题描述：一个整数N，找到1~N的全排列
+ 样例：
    - N = 3
    - (1,2,3) (1,3,2) (2,1,3) (2,3,1) (3,1,2) (3,2,1)

---

这样思考问题：N＝3，代表了有1、2、3个空位置和写有A、B、C的卡片，我们需要将卡片放到空位置上，并且每个位置只能放一张卡片，现在我们需要找出这3张卡片的所有不同摆放方法。

#### 1 
`约定顺序`

首先在位置1的时候，我们手里有A、B、C三张卡片，需要考虑应该先放哪张卡片。但是既然是要找到所有的可能，所以三种可能都需要尝试,我们可以约定一个顺序 A -> B -> C。

### 2
现在可以放置卡片了:
+ A-> 1
+ B-> 2
+ C-> 3

这时我们已经得到了一种全排 (A,B,C)

#### 3
+ 现在我们实际上已经走到了一个并不存在的位置4（结束位置，但是我们排列并没有结束）;
+ 现在我们需要立即回到位置3，取回卡片C，然后尝试是否还能尝试放入其它卡片（当然是按照第一步我们的约定A->B->C顺序），结果显然是我们手里并没有其它可以放入的卡片。
+ 我们需要继续回退一步来到位置2，取出卡片B，然后继续尝试是否能尝试放入其它卡片（当然是按照第一步我们的约定A->B->C顺序），这时我们发现可以放入卡片C。
+ 放入卡片后，我们向前一步来到位置3，尝试放入卡片（当然是按照第一步我们的约定A->B->C顺序），这时我们发现可以放入卡片B
+ 放完卡片，我们继续向前一步来到了并不存在的位置4
+ 我们得到了一种全排 (A, C, B)

#### 4
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

---

# 广度优先搜索

- 问题描述：

 *    |`[x1]`|`[x2]`|`[x3]`|`[x4]`
------|-----|-----|-----|-----
`[y1]`|  0  |  0  |  1  |  0 
`[y2]`|  0  |  0  |  0  |  0 
`[y3]`|  0  |  0  |  1  |  0
`[y4]`|  0  |  1  |__A__|  0
`[y5]`|  0  |  0  |  0  |  1
上面的表格是一张藏宝图的抽象，坐标中`0`表示空地，`1`表示有障碍物， `A`表示藏宝点，现在我们从坐标(1, 1)出发开始寻宝之旅。

---

这里我们用广度优先的思想来思考这个问题。


```
type node struct {
        x int // x坐标
        y int // y坐标
        f int // 父节点在队列中的编号
        s int // 步数
}


func BFS() {
        fmt.Println("BFS")

        var queue []node = make([]node, 50) // 队列

        var head, tail int // 分别标记队列头和尾

        var treasureMap = [5][4]int{ // 地图
                {0, 0, 1, 0},
                {0, 0, 0, 0},
                {0, 0, 1, 0},
                {0, 1, 0, 0},
                {0, 0, 0, 1},
        }

        var locationMarks [5][4]int // 标记地图上哪些点已经走过

        var derections = [4][2]int{ // 4个移动方向
                {1, 0},  // 向右
                {0, 1},  // 向下
                {-1, 0}, // 向左
                {0, -1}, // 向上

        }

        var targetX, targetY int = 2, 3
        var flag bool = false
        /*--------------------- 开始寻宝 ---------------------------*/

        /*------------ 从起点开始 ---------------*/
        // 初始化队列
        head = 0
        tail = 0

        // 标示起点
        // 起点入队 这里我们默认从(X0, Y0)开始
        queue[tail].x = 0
        queue[tail].y = 0
        queue[tail].f = head
        queue[tail].s = 0

        tail++                  // 尾指针后移一位
        locationMarks[0][0] = 1 // 标记已经走过的点

        /*------------ 走起 ---------------*/

        for head < tail { // 队列不为空

                // 枚举4个方向
                for i := 0; i < len(derections); i++ {
                        // 计算下一个点坐标
                        tx := queue[head].x + derections[i][0]
                        ty := queue[head].y + derections[i][1]

                        if tx > 3 || tx < 0 || ty > 4 || ty < 0 {
                                continue
                        }

                        // 坐标不是障碍物 并且 没有标记过
                        if treasureMap[ty][tx] == 0 && locationMarks[ty][tx] == 0 {
                                locationMarks[ty][tx] = 1 // 标记已经走过

                                // 将新的坐标入队
                                queue[tail].x = tx
                                queue[tail].y = ty
                                queue[tail].s = queue[head].s + 1
                                tail++
                        }

                        // 如果已经到达藏宝点
                        if tx == targetX && ty == targetY {

                                flag = true
                                break
                        }

                }

                if flag {
                        break
                } else {
                        head++ // 移动头指针，遍历下一层
                }

        }

        for i := 0; i < tail; i++ {
                fmt.Println(i, " ", queue[i])
        }
        
        for i := tail - 1; (queue[i].x != 0 || queue[i].y != 0) && i >= 0; {
                fmt.Println("Step ", queue[i].s, " (", queue[i].x, queue[i].y, ")")
                i = queue[i].f
        }

}

```


![结果](http://upload-images.jianshu.io/upload_images/1366868-d7e67de4427134aa.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


![步骤](http://upload-images.jianshu.io/upload_images/1366868-2816fa7f6283fa20.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)



---

# Floyd-Marshall最短路径


![图](http://upload-images.jianshu.io/upload_images/1366868-cfd11dfb2c3259fe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

在寻找任意两点间的最短路径时(A->B)，需要引入第三个点(C)，通过第三个点中转，看是否能够使 A->C->B 的距离小于 A->B 的距离。

```
// 999 表示两点之间不连通
var theMap = [4][4]int{
        {0, 2, 6, 4},
        {999, 0, 3, 999},
        {7, 999, 0, 1},
        {5, 999, 12, 0},
}

func Floyd() {
        fmt.Println("Floyd")

        for point := 0; point < 4; point++ {
                for i := 0; i < 4; i++ {
                        for j := 0; j < 4; j++ {
                                if theMap[i][j] > (theMap[i][point] + theMap[point][j]) {
                                        theMap[i][j] = theMap[i][point] + theMap[point][j]
                                }
                        }
                }
                fmt.Println(" Cover point()", point, ") ", theMap)
        }
}

```

![结果](http://upload-images.jianshu.io/upload_images/1366868-6553fdf1fffc2066.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


---

# Dijkstra单源最短路径


![图](http://upload-images.jianshu.io/upload_images/1366868-3484a0a3cff7a43a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


这里求定点A到各顶点的最短距离？

---

### 0

我们需要有一个数组记录当前已知的从顶点A到各顶点的最小距离：

![1.png](http://upload-images.jianshu.io/upload_images/1366868-59f6444dd9d4bf9d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 1 (第一轮)

从当前数组中找到一个离A顶点最近的顶点，即B （A->B = 2）。当选择了B顶点之后，A->B 也就是Dis[B]的值就从“估计值”变成了“确定值”。为什么呢？因为目前离A顶点最近的顶点已经是B了，图中并不存在负值的路径，就不可能有第三个点X使得 A->X->B 的距离小于当前的 A->B 的距离；如果存在这样一个点X的话，那么当前距离顶点A最近的点就不是B了，而是X。

### 2
既然选定了顶点B，那么我们可以看到B订单有两条出边：

- B -> C : 9
- B -> D : 3

这时我们想，既然B有到C、D的出边，那么从A到C、D是否会通过B顶点而变短呢（毕竟当前A->B的距离是已经是确信最短的了），所以我们比较：

- Dis[C]=`12` 和 Dis[B]+Map[B][C]=`10`
- Dis[D]=`&` 和 Dis[B]+Map[B][D]=`4`

结果我们发现A->C 和 A->D 的距离因为加入了B顶点中转使得距离变短，因此，我们可以更新顶点A的最小距离数组：

![2.png](http://upload-images.jianshu.io/upload_images/1366868-6c76be17f4c69938.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)



> 这个过程叫做 ***“松弛”***

### 4  (第二轮)

这时我们可以重复 **1** 的操作，cong从当前数组中的“估计值”（也就是 C、D、E、F）中找到离A最近的顶点，即D。同样Dis[D]的值从“估计值”变成了“确定值”。

### 5

D顶点的出边：

- D -> C : 4
- D -> E : 13
- D -> F : 15

通过D顶点来对qi其出边上的顶点进行松弛

- Dis[C]=`8` 和 Dis[D]+Map[D][C]=`13`
- Dis[E]=`&` 和 Dis[D]+Map[D][E]=`17`
- Dis[F]=`&` 和 Dis[D]+Map[D][F]=`19`

我们来更新顶点A的最小距离数组：

![3.png](http://upload-images.jianshu.io/upload_images/1366868-bfc4486fe03ddc64.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 6 (第三轮)

![4.png](http://upload-images.jianshu.io/upload_images/1366868-fbcec5037cafc5c3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 7 (第四轮)

![5.png](http://upload-images.jianshu.io/upload_images/1366868-474506a2b9287586.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 8 (第五轮)

![6.png](http://upload-images.jianshu.io/upload_images/1366868-ea690b81aee1c099.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

---

```
func Dijkstra() {
        // 999表示顶点之间没有连通
        var theMap = [6][6]int{
                {0, 1, 12, 999, 999, 999},
                {999, 0, 9, 3, 999, 999},
                {999, 999, 0, 999, 5, 999},
                {999, 999, 4, 0, 13, 15},
                {999, 999, 999, 999, 0, 4},
                {999, 999, 999, 999, 999, 0},
        }

        var marks = [6]int{1, 0, 0, 0, 0, 0} // 1，表示该顶点最短路径为确定值；0，表示该顶点的最短路径为估计值

        var dis [6]int

        // 初始化A顶点的最小路径数组
        for i := 0; i < 6; i++ {
                dis[i] = theMap[0][i]
        }

        fmt.Println("Dijkstra")
        // Dijkstra
        for i := 0; i < 5; i++ { // 这里为6个顶点，所以总共要进行5次 “松弛”
                minDistance := 1000 // 记录一次松弛中“估计值”中的最小距离
                currentPoint := 0   // 记录一次松弛中“估计值”中的顶点
                // 遍历最短距离数组，找到“估计值”中距离A顶点最近的顶点
                for j := 0; j < len(dis); j++ { //
                        if marks[j] == 0 && minDistance > dis[j] {
                                minDistance = dis[j]
                                currentPoint = j
                        }
                }

                marks[currentPoint] = 1 // 标记最小“估计值”为“确认值”

                // 遍历该顶点的出边并进行松弛
                for k := 0; k < 6; k++ {
                        if theMap[currentPoint][k] < 999 && dis[k] > (dis[currentPoint]+theMap[currentPoint][k]) {
                                dis[k] = dis[currentPoint] + theMap[currentPoint][k]
                        }
                }
        }

        fmt.Println("Dijkstra A -> ... ", dis)

}

```

![结果](http://upload-images.jianshu.io/upload_images/1366868-42404bb5e6485645.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


---


这个算法的时间复杂度是 O(N2)，其中每次寻找离A顶点最近的顶点的时间复杂度是O(N)，我们可以用“堆”来优化这部分，将这部分复杂度优化到O(logN)；

另外，我们考虑到在图中，边数M 通常是远小于N2的（这种图叫稀疏图，M相对较大的叫稠密图），我们可以考虑用另外一种表示方式来代替我们一直在用的 **邻接矩阵** —— **邻接表**


---

# 图的邻接表

![图](http://upload-images.jianshu.io/upload_images/1366868-41b334d840b19bfa.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

上图中有4个顶点5条边：


起始顶点 | 目标顶点 | 权值 |  边编号
---------|----------|------|---------
   A     |     B    |  9   |   1
   D     |     C    |  8   |   2
   A     |     B    |  5   |   3
   B     |     D    |  6   |   4
   A     |     C    |  7   |   5


---

## 邻接表

这里用数组来实现邻接表：

![邻接表](http://upload-images.jianshu.io/upload_images/1366868-b70c2a7c02768dfb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- `U` `V` `W` : U[i]->V[i] 权值为 W[i], 边编号为 i;
- `first`: first[i] 表示 i（A, B, C, D）号顶点的第一条边；
- `next` : next[i] 表示 i（1，2，3，4，5）号边的下一条边；


### 构建邻接表

看图不说话

![1.png](http://upload-images.jianshu.io/upload_images/1366868-e030d44df58bddb5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![2.png](http://upload-images.jianshu.io/upload_images/1366868-9ad8067d5335f904.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![3.png](http://upload-images.jianshu.io/upload_images/1366868-62e9222b79c2c75c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


![4.png](http://upload-images.jianshu.io/upload_images/1366868-90f80a9359509130.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![5.png](http://upload-images.jianshu.io/upload_images/1366868-a0ad8761db4e894f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

---

```
// 顶点编号：0，1，2，3 表示顶点 A，B，C，D
// 边编号：0，1，2，3，4
var u = [5]int{0, 3, 0, 1, 0}
var v = [5]int{3, 2, 1, 3, 2}
var w = [5]int{9, 8, 5, 6, 7}

func AdjacencyList() ([4]int, [5]int) {

        for i := 0; i < 5; i++ {
                fmt.Println(u[i], " -> ", v[i], " ", w[i])
        }

        first := [4]int{-1, -1, -1, -1}
        next := [5]int{-1, -1, -1, -1, -1}

        fmt.Println("fisrt : ", first)
        fmt.Println("next: ", next)
        fmt.Println("-----------------------")

        // 遍历5条边
        for i := 0; i < 5; i++ {
                next[i] = first[u[i]] // 将第i条边的下一条边设置为 u[i]号顶点的当前第一条边
                first[u[i]] = i       // 将u[i]号顶点的第一条边设置为当前边
        }

        fmt.Println("first: ", first)
        fmt.Println("next: ", next)

        return first, next
}


```

![结果](http://upload-images.jianshu.io/upload_images/1366868-52aa3ef75c6f3c7d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


--- 

> 用**邻接表**存储图的空间复杂度是O(M) （M边数），查找的时间复杂度也 为O(M)


遍历顶点的出边：

```
func TraverseEdge(point int, first [4]int, next [5]int) {
        k := first[point]
        for -1 != k {
                fmt.Println(u[k], " -> ", v[k], " ", w[k])
                k = next[k]
        }
}

```





































