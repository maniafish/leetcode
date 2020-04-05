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

func threeSum(n []int) [][]int {
	l := len(n)
	if l < 3 {
		return nil
	}

	sort.Ints(n)
	res := make([][]int, 0, l)
	// 遍历数组元素，每次拿当前元素和后两个元素取sum
	for i := 0; i < l-2; i++ {
		// 排序过后剩下的都比i大
		if n[i] > 0 {
			return res
		}

		// 去重
		if i > 0 && n[i] == n[i-1] {
			continue
		}

		for j1, j2 := i+1, l-1; j1 < j2; {
			sum := n[i] + n[j1] + n[j2]
			switch {
			case sum == 0:
				res = append(res, []int{n[i], n[j1], n[j2]})
				j1, j2 = j1+1, j2-1
				// j1继续右移并去重
				for j1 < j2 && n[j1] == n[j1-1] {
					j1++
				}
				// j2继续左移并去重
				for j1 < j2 && n[j2] == n[j2+1] {
					j2--
				}
			case sum > 0:
				// 需要缩小sum
				j2--
				// j2继续左移并去重
				for j1 < j2 && n[j2] == n[j2+1] {
					j2--
				}
			default:
				// 需要增大sum
				j1++
				// j1继续右移并去重
				for j1 < j2 && n[j1] == n[j1-1] {
					j1++
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, 0, 1, 0}))
}
