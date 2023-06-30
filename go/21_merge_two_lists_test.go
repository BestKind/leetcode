package main

import "testing"

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead
	for nil != l1 && nil != l2 {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}
	if nil != l1 {
		prehead.Next = l1
	}
	if nil != l2 {
		prehead.Next = l2
	}
	return result.Next
}

func TestMergeTwoLists(t *testing.T) {
	l := &ListNode{Val: 1}
	p := l
	p.Next = &ListNode{Val: 2}
	p = p.Next
	p.Next = &ListNode{Val: 4}
	r := &ListNode{Val: 1}
	p = r
	p.Next = &ListNode{Val: 3}
	p = p.Next
	p.Next = &ListNode{Val: 4}

	res := mergeTwoLists(l, r)
	
	for res != nil {
		t.Log(res.Val)
		res = res.Next
	}
}
