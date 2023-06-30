package main

import (
	"fmt"
	"testing"
)

/*
*
  - https://leetcode-cn.com/problems/unique-paths-ii/
  - 一个机器人位于一个 m x n 网格的左上角。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角。
说明：m 和 n 的值均不超过 100。

示例 1:

输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
*/
func uniquePathsWithObstacles(obstacle [][]int) int {
	m, n := len(obstacle), len(obstacle[0])
	if 1 > m {
		return 0
	}
	if 1 > n {
		return 0
	}
	if obstacle[0][0] == 1 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 && j != 0 {
				if obstacle[i][j] == 0 {
					dp[i][j] = dp[i][j-1]
				}
			} else if i != 0 && j == 0 {
				if obstacle[i][j] == 0 {
					dp[i][j] = dp[i-1][j]
				}
			} else {
				if obstacle[i][j] == 0 {
					dp[i][j] = dp[i-1][j] + dp[i][j-1]
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacle := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	res := uniquePathsWithObstacles(obstacle)
	fmt.Println(res)

	obstacle = [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}
	res = uniquePathsWithObstacles(obstacle)
	fmt.Println(res)
	obstacle = [][]int{
		{0},
		{0},
		{1},
	}
	res = uniquePathsWithObstacles(obstacle)
	fmt.Println(res)
}
