package main

import (
	"fmt"
	"testing"
)

/**
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

链接：https://leetcode-cn.com/problems/unique-binary-search-trees
*/

/*
*
题解：符合 卡塔兰数 ，即可用卡塔兰数解题
*/
func numTrees(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2*i + 1) / (i + 2)
	}
	return C
}

func TestNumTrees(t *testing.T) {
	n := 3
	res := numTrees(n)
	fmt.Println(res)
}
