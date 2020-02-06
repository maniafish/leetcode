package main

import "fmt"

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	// 长度，当前滑动窗口(最长子串)的起始位置，当前滑动窗口长度，最大滑动窗口的长度
	l, begin, currentLen, res := len(s), 0, 0, 0
	// 存放已经出现过的字符(ascii)和最后一次出现的位置
	u := make(map[byte]int, l)
	for i := 0; i < l; i++ {
		// 该字符在当前滑动窗口出现了
		if index, ok := u[s[i]]; ok && index >= begin {
			// 设置下一滑动窗口起始为止为出现处的下一位
			begin = index + 1
			if currentLen > res {
				res = currentLen
			}

			currentLen = i - index - 1
		}

		currentLen++
		// 当前滑动窗口长度 > 最大滑动窗口长度
		if currentLen > res {
			res = currentLen
		}
		u[s[i]] = i
	}

	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
