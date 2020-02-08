package main

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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

func isPalindrome(head *ListNode) bool {
	// 快慢指针
	p1, p2 := head, head
	n := make([]int, 0, 1)
	for p2 != nil {
		if p2.Next == nil {
			p1 = p1.Next
			break
		}

		n = append(n, p1.Val)
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	// 反转后半段链表, 对比前后半段的值
	for p1, p2 = reverseList(p1), head; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Val != p2.Val {
			return false
		}
	}

	return true
}
