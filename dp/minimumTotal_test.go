package dp

import (
	"fmt"
	"testing"
)

//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
func Test_MinimumTotal(t *testing.T) {

	var triangle = [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal1(triangle))
	fmt.Println(minimumTotal(triangle))
}

//dp自底向上
func minimumTotal1(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	// 1、状态定义：f[i][j] 表示从i,j出发，到达最后一层的最短路径
	var l = len(triangle)
	var f = make([][]int, l)
	// 2、初始化
	for i := 0; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if f[i] == nil {
				f[i] = make([]int, len(triangle[i]))
			}
			f[i][j] = triangle[i][j]
		}
	}
	// 3、递推求解
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			fmt.Println("b", f[i][j])
			f[i][j] = min(f[i+1][j], f[i+1][j+1]) + triangle[i][j]
			fmt.Println("a", f[i][j])

		}
	}
	// 4、答案
	return f[0][0]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//dp自顶向下

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	// 1、状态定义：f[i][j] 表示从0,0出发，到达i,j的最短路径
	var l = len(triangle)
	var f = make([][]int, l)
	// 2、初始化
	for i := 0; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if f[i] == nil {
				f[i] = make([]int, len(triangle[i]))
			}
			f[i][j] = triangle[i][j]
		}
	}
	// 递推求解
	for i := 1; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			// 这里分为两种情况：
			// 1、上一层没有左边值
			// 2、上一层没有右边值
			if j-1 < 0 {
				f[i][j] = f[i-1][j] + triangle[i][j]
			} else if j >= len(f[i-1]) {
				f[i][j] = f[i-1][j-1] + triangle[i][j]
			} else {
				f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	result := f[l-1][0]
	for i := 1; i < len(f[l-1]); i++ {
		result = min(result, f[l-1][i])
	}
	return result
}

func minimumTotal3(triangle [][]int) int {
	h := len(triangle)
	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	for i := h - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == h-1 {
				dp[i][j] = triangle[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func minimumTotal4(triangle [][]int) int {
	bottom := triangle[len(triangle)-1]
	dp := make([]int, len(bottom))
	for i := range dp {
		dp[i] = bottom[i]
	}

	for i := len(dp) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
