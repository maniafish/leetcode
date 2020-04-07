package main

/*
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5
*/

/*
归并排序，拆分成两个有序链表的排序(最小拆分至两个长度为1的链表排序)
*/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 快慢指针找中点
	p1, p2 := head, head, head.Next
	for ; p2 != nil && p2.Next != nil; p1, p2 = p1.Next, p2.Next.Next {
	}

	// 切断链表
	p0 := p1
	p1 = p1.Next
	p0.Next = nil
	return mergeList(sortList(head), sortList(p1))
}

func mergeList(l1, l2 *ListNode) *ListNode {
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
			p, l1 = p.Next, l1.Next
		} else {
			p.Next = l2
			p, l2 = p.Next, l2.Next
		}
	}

	return res
}
