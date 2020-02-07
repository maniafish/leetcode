package main

import "fmt"

/*
给定一个字符串s，找到其中最长的回文子序列。可以假设s的最大长度为1000。

示例 1:
输入:

"bbbab"
输出:

4
一个可能的最长回文子序列为 "bbbb"。

示例 2:
输入:

"cbbd"
输出:

2
一个可能的最长回文子序列为 "bb"。
*/

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestPalindromeSubseq(s string) int {
	// s长度，最大回文子序列长度
	l, maxLen := len(s), 1
	if l == 0 {
		return 0
	}

	// 动态规划矩阵，i(回文子序列起始), j(回文子序列终止), m[i][j](起止之间的回文子序列长度)
	m := make([][]int, l)
	for i := 0; i < l; i++ {
		m[i] = make([]int, l)
	}

	m[l-1][l-1] = 1
	// if s[i] == s[j]: m[i][j] = m[i+1][j-1] + 2
	// else: m[i][j] = max(m[i][j-1], m[i+1][j])
	// 由于i要用到i+1的值，所以反向遍历
	for i := l - 2; i >= 0; i-- {
		m[i][i] = 1
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				m[i][j] = m[i+1][j-1] + 2
			} else {
				m[i][j] = max(m[i+1][j], m[i][j-1])
			}

			if m[i][j] > maxLen {
				maxLen = m[i][j]
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
	fmt.Println(longestPalindromeSubseq("cbbc"))
	fmt.Println(longestPalindromeSubseq("cbc"))
}
