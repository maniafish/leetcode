package main

import "fmt"

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。

示例:

输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。

*/

func minSubArrayLen(s int, nums []int) int {
	l := len(nums)
	// 左右指针，总和，最小长度
	left, right, sum, min := 0, 0, 0, l+1
	for ; right < l; right++ {
		sum += nums[right]
		// 左指针处理
		for ; left <= right && sum >= s; left++ {
			if right-left+1 < min {
				min = right - left + 1
			}

			sum -= nums[left]
		}
	}

	if min != l+1 {
		return min
	}

	return 0
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
