package dynamic

import (
	"fmt"
	"testing"
)

// #70 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

func Test_Pa_lou_ti(t *testing.T) {
	fmt.Println(climbStairs(3))

}

//dp[i]：爬 i 级楼梯的方式数
func climbStairs(n int) int {
	dp := make([]int, n+1)
	if n < 2 {
		return 1
	}
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]

}
