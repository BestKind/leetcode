package main

import (
	"fmt"
	"testing"
)

/**
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。
示例:
给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
链接：https://leetcode-cn.com/problems/path-sum
*/

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func hasPathSum(root *TreeNode1, sum int) bool {
	if nil == root {
		return false
	}
	res := root.Val
	fmt.Println(res, sum)
	if res == sum && nil == root.Left && nil == root.Right {
		return true
	}
	if nil != root.Left {
		flag := hasPathSum(root.Left, sum-res)
		if flag {
			return flag
		}
	}
	if nil != root.Right {
		flag := hasPathSum(root.Right, sum-res)
		if flag {
			return flag
		}
	}

	return false
}

func TestHasPathSum(t *testing.T) {
	root := createTreeNode(5)
	root.Left = createTreeNode(4)
	root.Left.Left = createTreeNode(11)
	root.Left.Left.Left = createTreeNode(7)
	root.Left.Left.Right = createTreeNode(2)
	root.Right = createTreeNode(8)
	root.Right.Left = createTreeNode(13)
	root.Right.Right = createTreeNode(4)
	root.Right.Right.Right = createTreeNode(1)
	res := hasPathSum(root, 22)
	fmt.Println(res)

	root = createTreeNode(-2)
	root.Right = createTreeNode(-3)
	res = hasPathSum(root, -5)

	fmt.Println(res)
}

func createTreeNode(val int) *TreeNode1 {
	return &TreeNode1{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}
