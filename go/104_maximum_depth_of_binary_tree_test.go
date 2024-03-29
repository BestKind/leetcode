package main

import (
	"fmt"
	"testing"
)

/**
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
示例：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
链接：https://leetcode-cn.com/problems/maximum-depth-of-binary-tree
*/

type myTreeNode struct {
	Val   int
	Left  *myTreeNode
	Right *myTreeNode
}

func maxDepth(root *myTreeNode) int {
	if nil == root {
		return 0
	}
	leftLength := maxDepth(root.Left)
	rightLength := maxDepth(root.Right)

	depth := leftLength
	if leftLength < rightLength {
		depth = rightLength
	}
	return 1 + depth
}

func TestMaxDepth(t *testing.T) {
	tree := createNode(3)
	tree.Left = createNode(9)
	tree.Right = createNode(20)
	tree.Right.Left = createNode(15)
	tree.Right.Right = createNode(7)

	fmt.Println(maxDepth(tree))
}

func createNode(node int) *myTreeNode {
	tree := myTreeNode{
		Val:   node,
		Left:  nil,
		Right: nil,
	}
	return &tree
}
