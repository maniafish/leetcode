package main

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
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

func threeSum(n Num, target, addition int) [][]int {
	res := make([][]int, 0, 1)
	done := make(map[string]bool)
	for i := 1; i < len(n)-1; i++ {
		j1, j2 := 0, len(n)-1
		for j1 < i && j2 > i {
			sum := n[j1] + n[i] + n[j2]
			if sum == target {
				doneStr := fmt.Sprintf("%d%d%d%d", addition, n[j1], n[i], n[j2])
				if _, ok := done[doneStr]; !ok {
					res = append(res, []int{addition, n[j1], n[i], n[j2]})
					done[doneStr] = true
				}
				j1 += 1
				j2 -= 1
				continue
			}

			if sum < target {
				j1 += 1
				continue
			}

			if sum > target {
				j2 -= 1
				continue
			}
		}
	}

	return res
}

func fourSum(nums []int, target int) [][]int {
	n := Num(nums)
	sort.Sort(n)
	res := make([][]int, 0, 1)
	for i := 0; i < n.Len()-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		resTemp := threeSum(nums[i+1:], target-nums[i], nums[i])
		if len(resTemp) > 0 {
			res = append(res, resTemp...)
		}
	}

	return res
}
