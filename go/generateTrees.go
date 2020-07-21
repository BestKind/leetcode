package main

import "fmt"

/**
给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
示例：
输入：3
输出：
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：
   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
提示：
0 <= n <= 8
链接：https://leetcode-cn.com/problems/unique-binary-search-trees-ii
 */

type TreeNode2 struct {
	Val int
	Left *TreeNode2
	Right *TreeNode2
}

func generateTrees(n int) []*TreeNode2 {
	if 0 == n {
		return nil
	}
	return trees(1, n)
}

func trees(start, end int) []*TreeNode2 {
	if start > end {
		return []*TreeNode2{nil}
	}
	allTrees := []*TreeNode2{}
	for i := start; i <= end; i++ {
		leftTree := trees(start, i - 1)
		rightTree := trees(i + 1, end)
		for _, left := range leftTree {
			for _, right := range rightTree {
				curTree := &TreeNode2{i, nil, nil}
				curTree.Left = left
				curTree.Right = right
				allTrees = append(allTrees, curTree)
			}
		}
	}
	return allTrees
}

func main() {
	res := generateTrees(3)
	for _, node := range res {
		travel1(node)
		fmt.Println("===========")
	}
}

func travel1(node *TreeNode2) {
	if nil == node {
		return
	}
	travel1(node.Left)
	fmt.Println(node.Val)
	travel1(node.Right)
}
