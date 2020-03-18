package main

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

参考图链接：https://leetcode-cn.com/problems/container-with-most-water
*/

func maxArea(height []int) int {
	// 双指针法，从两头往中间走，每次将两头中的短板位置指针往中间走一格，最终取面积最大值
	left, right, res := 0, len(height)-1, 0
	for left < right {
		min := height[left]
		if height[right] < min {
			min = height[right]
			// 右指针左移
			right -= 1
		} else {
			// 左指针右移
			left += 1
		}

		s := min * (right - left + 1)
		if s > res {
			res = s
		}
	}

	return res
}
