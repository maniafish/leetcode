package main

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

import (
	"fmt"
	"sort"
)

type Num []int

func (a Num) Len() int           { return len(a) }
func (a Num) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Num) Less(i, j int) bool { return a[i] < a[j] }

func threeSum(nums []int) [][]int {
	n := Num(nums)
	sort.Sort(n)
	res := make([][]int, 0, len(nums))
	done := make(map[string]bool)
	for i := 1; i < n.Len()-1; i++ {
		j1, j2 := 0, n.Len()-1
		for j1 < i && j2 > i {
			sum := n[j1] + n[i] + n[j2]
			if sum == 0 {
				doneStr := fmt.Sprintf("%d%d%d", n[j1], n[i], n[j2])
				if _, ok := done[doneStr]; !ok {
					res = append(res, []int{n[j1], n[i], n[j2]})
					done[doneStr] = true
				}
				j1 += 1
				j2 -= 1
				continue
			}

			if sum < 0 {
				j1 += 1
				continue
			}

			if sum > 0 {
				j2 -= 1
				continue
			}
		}
	}

	return res
}
