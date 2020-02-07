package main

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

func longestPalindrome(s string) string {
	// s长度，最大回文串长度，最大回文串
	l, max, subString := len(s), -1, ""
	if l == 0 {
		return ""
	}

	// 动规矩阵：i(字符串起始), j(字符串末尾), m[i][j](该串是否为回文串)
	m := make([][]bool, l)
	for i := 0; i < l; i++ {
		m[i] = make([]bool, l)
	}

	// 只要符合s[i] == s[j] && m[i+1][j-1] == true，则当前m[i][j] = true；由于i要用到i+1的值，所以我们倒着遍历
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if s[i] == s[j] {
				if i+1 >= j-1 || m[i+1][j-1] {
					m[i][j] = true
					if j-i > max {
						max = j - i
						subString = s[i : j+1]
					}
					continue
				}
			}
		}
	}

	return subString
}
