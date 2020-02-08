package main

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p1, p2, p3 := head, head.Next, head
	for p1.Next = nil; p2 != nil; p2 = p3 {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
	}

	return p1
}
