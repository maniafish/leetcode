package main

import "container/heap"

/*
编写一个程序，找出第 n 个丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。
*/

/*
题解：堆
基本思路：如何完整地遍历丑数，并且按序弹出
我们从堆中包含一个数字开始：1，去计算下一个丑数。
1. 将1从堆中弹出然后将三个数字添加到堆中：1*2, 1*3，1*5。
2. 现在最小数字是2，说明它是第二个丑数，将2弹出，并压入2*2, 2*3, 2*5
3. 现在最小的数字是3，将它弹出并压入3*2, 3*3, 3*5
4. 依次类推，压入k*2, k*3, k*5

如何去重：每次弹出堆顶元素后，若当前堆顶元素和之前的元素重复，就继续弹出堆顶元素
*/

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(a, b int) bool  { return h[a] < h[b] }
func (h IntHeap) Swap(a, b int)       { h[a], h[b] = h[b], h[a] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := h.Len()
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func nthUglyNumber(n int) int {
	if n < 1 {
		return 0
	}

	h := IntHeap([]int{1})
	heap.Init(&h)
	for i := 0; i < n-1; i++ {
		min := heap.Pop(&h).(int)
		for h.Len() > 0 && min == h[0] {
			heap.Pop(&h)
		}

		heap.Push(&h, 2*min)
		heap.Push(&h, 3*min)
		heap.Push(&h, 5*min)
	}

	return h[0]
}
