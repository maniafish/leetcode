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
	if n <= m {
		return head
	}

	// 前段最后一个节点，中段三个反转指针，中段头一个节点
	p, p1, p2, p3, p4 := head, head, head, head, head
	if m > 1 { // 前面至少有一个节点保持原样
		for i := 1; i < m-1 && p != nil; i, p = i+1, p.Next {
		}

		if p == nil || p.Next == nil || p.Next.Next == nil {
			return head
		}

		p1, p2, p3, p4 = p.Next, p.Next.Next, p.Next, p.Next
	} else { // 从第一位开始反转
		if p == nil || p.Next == nil {
			return head
		}

		p1, p2, p3, p4 = p, p.Next, p, p
		p = nil
	}

	for i := m; i < n && p2 != nil; i++ {
		p3 = p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}

	// 中段头(即反转后的尾)指向后段头
	p4.Next = p3
	if p != nil { // 有前段
		p.Next = p1
		return head
	} else {
		return p1
	}
}
