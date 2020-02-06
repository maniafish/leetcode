package main

import "fmt"

/*
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
说明：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。
*/

// Index 记录匹配字符出现的位置
type Index struct {
	Idx  int
	Data byte
}

func minWindow(s string, t string) string {
	l1, l2 := len(s), len(t)
	if l1 == 0 || l2 == 0 {
		return ""
	}

	// 目标map匹配次数，当前窗口已匹配的字符次数，当前窗口字符是否已满足匹配要求
	u, found, match := make(map[byte]int, l2), make(map[byte]int), make(map[byte]bool)
	for i := 0; i < l2; i++ {
		count := u[t[i]]
		u[t[i]] = count + 1
	}

	idl, j := make([]*Index, 0, l2), 0
	// 当前窗口左指针, 长度；最小窗口左指针，长度
	left, currentLen, begin, minLen := 0, 0, 0, 0
	for i := 0; i < l1; i++ {
		currentLen++
		if _, ok := u[s[i]]; ok {
			count := found[s[i]]
			found[s[i]] = count + 1
			if found[s[i]] >= u[s[i]] {
				match[s[i]] = true
			}

			idl = append(idl, &Index{i, s[i]})
		}

		// 全找到了
		for len(match) == len(u) {
			idx, data := idl[j].Idx, idl[j].Data
			if minLen == 0 || currentLen <= minLen {
				// 最小窗口左指针指向第一个匹配的
				begin = idx
				minLen = currentLen - (idx - left)
			}

			// 窗口左指针右移
			j++
			// 只要求匹配一次
			if j >= len(idl) {
				return t
			}

			currentLen -= idl[j].Idx - left
			left = idl[j].Idx

			// 匹配次数自减
			found[data]--
			if found[data] < u[data] {
				delete(match, data)
			}
		}

	}

	// 没有找全
	if minLen == 0 {
		return ""
	}

	return s[begin : begin+minLen]
}

func minWindow1(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	hash := make([]int, 256)
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}

	l, count, max, results := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}

		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}

		if count == 0 && max > r-l+1 {
			max = r - l + 1
			results = s[l : r+1]
		}
	}

	return results
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("AA", "AA"))
}
