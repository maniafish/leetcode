package main

/*
将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	res := l1
	if l2.Val < l1.Val {
		res = l2
		l2 = l2.Next
	} else {
		l1 = l1.Next
	}

	p := res
	for {
		if l1 == nil {
			p.Next = l2
			return res
		}

		if l2 == nil {
			p.Next = l1
			return res
		}

		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}

		p = p.Next
	}

	return res
}
