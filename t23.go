package main

import "container/heap"

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (h NodeHeap) Len() int            { return len(h) }
func (h NodeHeap) Less(a, b int) bool  { return h[a].Val < h[b].Val }
func (h NodeHeap) Swap(a, b int)       { h[a], h[b] = h[b], h[a] }
func (h *NodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *NodeHeap) Pop() interface{} {
	n := h.Len()
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func mergeKLists(l []*ListNode) *ListNode {
	ll := len(l)
	if ll < 1 {
		return nil
	}

	// 每一行第一个元素构建最小堆
	h := NodeHeap(make([]*ListNode, 0, ll))
	for i := 0; i < ll; i++ {
		if l[i] != nil {
			h.Push(l[i])
		}
	}

	if h.Len() < 1 {
		return nil
	}

	heap.Init(&h)
	res := h[0]
	p := res
	for {
		if h[0].Next != nil {
			h[0] = h[0].Next
			heap.Fix(&h, 0)
		} else {
			heap.Pop(&h)
		}

		if h.Len() < 1 {
			break
		}

		p.Next = h[0]
		p = p.Next
	}

	return res
}
