package main

import "container/heap"

/*
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]
*/

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(a, b int) bool  { return h[a] > h[b] }
func (h IntHeap) Swap(a, b int)       { h[a], h[b] = h[b], h[a] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := h.Len()
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func getLeastNumbers(arr []int, k int) []int {
	l := len(arr)
	if k >= l {
		return arr
	}

	if k == 0 {
		return nil
	}

	// 前k个数构建最大堆
	h := IntHeap(arr[0:k])
	heap.Init(&h)
	for i := k; i < l; i++ {
		// 当前元素小于堆顶元素，说明当前元素属于前k小的数
		if arr[i] < h[0] {
			// 弹出堆顶元素，并压入当前元素，调整堆
			heap.Pop(&h)
			heap.Push(&h, arr[i])
		}
	}

	return h
}
