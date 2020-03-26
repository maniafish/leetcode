package main

/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
示例 2:

输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/

func canJump(nums []int) bool {
	l := len(nums)
	if l < 1 {
		return false
	}

	// 自己跳自己
	if l == 1 {
		return true
	}

	// 一开始就跳不动了
	if nums[0] == 0 {
		return false
	}

	zero := -1
	for i := l - 2; i >= 0; i-- {
		// 已经有0值了
		if zero > 0 {
			// 可以跳过0值
			if i+nums[i] > zero {
				// 当前0值可以忽略
				zero = -1
			}

			continue
		}

		if nums[i] == 0 {
			zero = i
			continue
		}

	}

	if zero < 0 {
		return true
	}

	return false
}
