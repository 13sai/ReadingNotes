package code

import "math"

// 斐波那契数列
// 时间复杂度：O(2^n)
func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// 重叠子问题
func Fib2(n int) int {
	sli := make(map[int]int)
	return fib2_helper(sli, n)
}

func fib2_helper(arr map[int]int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	if v, ok := arr[n]; ok {
		return v
	}

	arr[n] = arr[n-1] + arr[n-2]
	return arr[n]
}

// dp 数组的迭代解法
func DPFib(n int) int {
	if n == 0 ||n == 1 {
		return n
	}
	res := make([]int, n + 1)
	res[0] = 0
	res[1] = 1
	for i := 2 ; i <= n ; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}

// 优化空间复杂度为O(1)
func DPFib2(n int) int {
	if n < 2 {
		return n
	}
	m1, m2 := 0, 1
	for i := 0; i < n; i++ {
		m1, m2 = m2, m1+m2
	}

	return m1
}


// 凑零钱问题
// 给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。

// dp[i] = min(dp[i], dp[i - coins[i]] + 1)
func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	arr := make([]int, amount+1)
	arr[0] = 0
	for i:=1; i <= amount; i++ {
		arr[i] = amount + 1
		for _, v := range coins {
			if i >= v {
				arr[i] = min(arr[i], arr[i-v]+1)
			}
		}
	}

	if arr[amount] == amount + 1 {
		return -1
	}

	return arr[amount]
}


func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
