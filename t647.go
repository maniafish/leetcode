package main

import "fmt"

/*
给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。

示例 1:

输入: "abc"
输出: 3
解释: 三个回文子串: "a", "b", "c".
示例 2:

输入: "aaa"
输出: 6
说明: 6个回文子串: "a", "a", "a", "aa", "aa", "aaa".
*/

func countSubstrings(s string) int {
	// s长度
	l := len(s)
	if l == 0 {
		return 0
	}

	sum := 0
	// 动规矩阵：i(字符串起始), j(字符串末尾), m[i][j](该串是否为回文串)
	m := make([][]bool, l)
	for i := 0; i < l; i++ {
		m[i] = make([]bool, l)
	}

	// 只要符合s[i] == s[j] && m[i+1][j-1] == true，则当前m[i][j] = true；由于i要用到i+1的值，所以我们倒着遍历
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if s[i] == s[j] && (i+1 >= j-1 || m[i+1][j-1]) {
				m[i][j] = true
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}
