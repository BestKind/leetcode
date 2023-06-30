package main

import (
	"math"
	"testing"
)

/**
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.

链接：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = minDe(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = minDe(minDepth(root.Right), minD)
	}
	return minD + 1
}

func minDe(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func TestMinDepth(t *testing.T) {
	r := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	t.Log(minDepth(r))
}
