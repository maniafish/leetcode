package main

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func doCarry(p *ListNode) {
	for {
		// 进位已经超过最高位
		if p.Next == nil {
			p.Next = &ListNode{Val: 1}
			return
		}

		if p.Next.Val == 9 {
			// 进位
			p.Next.Val = 0
			p = p.Next
			continue
		}

		p.Next.Val += 1
		return
	}
}

func calVal(v1, v2 int, carry bool) (int, bool) {
	c := 0
	if carry {
		c = 1
	}

	sum := v1 + v2 + c
	// 有进位
	if sum/10 > 0 {
		return sum % 10, true
	}

	return sum % 10, false
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 用一个空的头指针
	l := new(ListNode)
	p, carry := l, false
	for {
		if l1 == nil {
			p.Next = l2
			if carry {
				// 单独处理剩余进位
				doCarry(p)
			}

			break
		}

		if l2 == nil {
			p.Next = l1
			if carry {
				// 单独处理剩余进位
				doCarry(p)
			}

			break
		}

		p.Next = &ListNode{}
		// 计算当前位数和进位
		p.Next.Val, carry = calVal(l1.Val, l2.Val, carry)
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	return l.Next
}
