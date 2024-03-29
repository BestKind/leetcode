package main

import (
	"fmt"
	"math"
	"testing"
)

/**
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。

例如，给定三角形：
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：
如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
链接：https://leetcode-cn.com/problems/triangle
*/

// 方法一
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i][0] = f[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[i][j] = myMin1(f[i-1][j-1], f[i-1][j]) + triangle[i][j]
		}
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = myMin1(ans, f[n-1][i])
	}
	return ans
}

// 方法二
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	f := [2][]int{}
	for i := 0; i < 2; i++ {
		f[i] = make([]int, n)
	}
	f[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		curr := i % 2
		prev := 1 - curr
		f[curr][0] = f[prev][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[curr][j] = myMin1(f[prev][j-1], f[prev][j]) + triangle[i][j]
		}
		f[curr][i] = f[prev][i-1] + triangle[i][i]
		fmt.Println(f)
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = myMin1(ans, f[(n-1)%2][i])
	}
	return ans
}

func myMin1(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{
		{-1},
		{2, 3},
		{1, -1, -3},
	}
	t.Log(minimumTotal2(triangle))
	t.Log(minimumTotal(triangle))
}
