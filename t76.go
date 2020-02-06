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

func minWindow(s string, t string) string {
	lS, lT := len(s), len(t)
	if lT == 0 {
		return ""
	}

	// 当前左指针，当前右指针；最小窗口左指针，最小窗口右指针(默认超过长度上限)
	left, right, minLeft, minRight := 0, 0, 0, lS
	// 目标匹配次数map, 当前已完成匹配的字符数, 目标匹配字符数
	u, found, target := make(map[byte]int), 0, 0
	for i := 0; i < lT; i++ {
		if _, ok := u[t[i]]; ok {
			u[t[i]] += 1
		} else {
			u[t[i]] = 1
			target += 1
		}
	}

	for ; right < lS; right++ {
		if _, ok := u[s[right]]; ok {
			u[s[right]] -= 1
			// 刚好完成当前字符匹配
			if u[s[right]] == 0 {
				found += 1
			}

			// 完成匹配，做窗口滑动处理(右移left)
			if found >= target {
				for ; left <= right; left++ {
					if _, ok := u[s[left]]; ok {
						// 最小窗口变动
						if right-left <= minRight-minLeft {
							minLeft = left
							minRight = right
						}

						// 去掉当前字符
						u[s[left]] += 1
						// 当前字符数已不符合匹配要求
						if u[s[left]] > 0 {
							found -= 1
							left++
							break
						}
					}
				}
			}
		}
	}

	// 匹配上了
	if minRight != lS {
		return s[minLeft : minRight+1]
	}

	return ""
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("AA", "AA"))
	fmt.Println(minWindow("ABC", "AC"))
}
