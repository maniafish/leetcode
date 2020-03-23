package main

import "fmt"

/*
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true
示例 2:

输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出: false
*/

/*
题解：
动态规划，dp[i][j]中的i表示s1的下标，j表示s2的下标，值表示s1[:i]和s2[:j]能否组成s3[:i+j]
if dp[i-1][j-1] && s1[i-1] = s3[i+j-1]{dp[i][j] = true}
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	dp, l1, l2, l3 := make(map[int]map[int]bool), len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}

	// 初始化dp，dp[i][0]表示纯s1是否能组成s3前缀，dp[0][j]表示纯s2是否能组成s3前缀
	for i := 0; i < l1+1; i++ {
		dp[i] = make(map[int]bool, l2+1)
		// 首位空出
		if i == 0 {
			dp[0][0] = true
		}
		for j := 0; j < l2+1; j++ {
			if (i > 0 && dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (j > 0 && dp[i][j-1] && s2[j-1] == s3[i+j-1]) {
				dp[i][j] = true
			}
		}
	}

	return dp[l1][l2]
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	fmt.Println(isInterleave(s1, s2, s3))
}
