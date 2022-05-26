package alg

var Prices = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

func CutRod(prices []int, n int) int {
	if n == 0 {
		return 0
	}
	q := -1
	for i:=1; i<=n; i++ {
		q_i := prices[i] + CutRod(prices, n-i)
		q = MaxInt(q, q_i)
	}
	return q
}


func MemoCutRod(prices []int, n int) int {
	r := make([]int, n+1, n+1)
	return MemoCutRodAux(prices, n, r)
}
func MemoCutRodAux(prices []int, n int, r []int) int {
	if n == 0 {
		r[0] = 0
		return 0
	}
	if r[n] > 0 {
		return r[n]
	}
	q := -1
	for i:=1; i<=n; i++ {
		q_i := prices[i] + MemoCutRodAux(prices, n-i, r)
		q = MaxInt(q, q_i)
	}
	r[n] = q
	return q
}

func BottomUpCutRod(prices []int, n int) int {
	r := make([]int, n+1, n+1)
	for j:=1; j<=n; j++ {
		q := -1
		for i:=1; i<=j; i++ {
			q = MaxInt(q, prices[i]+r[j-i])
		}
		r[j] = q
	}
	return r[n]
}


// r = [0 1 5 8 10 13 17 18 22 25 30]
// sols =[0 1 2 3 2 2 6 1 2 3 10]
func ExtendedBottomUpCutRod(prices []int, n int) (r, sols [] int){
	r = make([]int, n+1, n+1)
	sols = make([]int, n+1, n+1) // 保存的是求解规模为 j 的子问题时，第一段钢条的最优切割长度 i；
	for j:=1; j<=n; j++ {
		q := -1
		for i:=1; i<=j; i++ {
			new_q := prices[i] + r[j-i]
			if new_q > q {
				q = new_q
				sols[j] = i
			}
		}
		r[j] = q
	}
	return r, sols
}


func MaxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

