package main

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	/*
		链表中的点已经相连，一次旋转操作意味着：
			先将链表闭合成环
			找到相应的位置断开这个环，确定新的链表头和链表尾
	*/
	if head == nil || head.Next == nil || k < 1 {
		return head
	}

	// 链表长度
	l, p := 1, head
	for ; p.Next != nil; p, l = p.Next, l+1 {
	}

	k = k % l
	if k == 0 {
		return head
	}

	p.Next = head
	p = head
	// l-k的位置断开环并返回
	for i := 1; i < l-k; i, p = i+1, p.Next {
	}

	head = p.Next
	p.Next = nil
	return head
}
