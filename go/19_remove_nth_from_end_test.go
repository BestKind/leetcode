package main

import "testing"

/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
给定的 n 保证是有效的。
进阶：
你能尝试使用一趟扫描实现吗？
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if nil == head {
		return nil
	}

	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	pre = result
	i := 1
	for head != nil {
		head = head.Next
		i++
		if i >= (n + 2) {
			pre = pre.Next
		}
	}

	pre.Next = pre.Next.Next

	return result.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	l := &ListNode{Val: 1}
	p := l
	for i := 2; i <= 10; i++ {
		p.Next = &ListNode{Val: i}
		p = p.Next
	}
	res := removeNthFromEnd(l, 3)
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
