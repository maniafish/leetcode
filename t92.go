package main

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if n <= m || head == nil || head.Next == nil {
		return head
	}

	var h1 *ListNode
	p1, p2, p3 := head, head.Next, head.Next
	for i := 0; i < m-1 && p2 != nil; i++ {
		h1, p1, p2, p3 = p1, p1.Next, p2.Next, p3.Next
	}

	for i := m; i < n && p2 != nil; i++ {
		p3 = p2.Next
		p2.Next = p1
		p1, p2 = p2, p3
	}

	if h1 == nil { // 头是p1
		head.Next = p2
		return p1
	}

	h2 := h1.Next
	h1.Next, h2.Next = p1, p2
	return head
}
