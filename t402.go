package main

import (
	"bytes"
	"fmt"
)

/*
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

注意:

num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。
示例 1 :

输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
示例 2 :

输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
示例 3 :

输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
*/

func removeKdigits(num string, k int) string {
	l := len(num)
	// 边界条件
	if k <= 0 {
		return num
	}

	if k >= l || l == 0 {
		return "0"
	}

	buffer := bytes.Buffer{}
	buffer.WriteByte(num[0])
	// 从左往右删，一旦出现右边的数小于左邻居的情况，就删掉左边的邻居；直到删除k个元素为止
	for i, left := 1, num[0]; i < l; i++ {
		for num[i] < left && k > 0 && buffer.Len() >= 1 {
			buffer.Truncate(buffer.Len() - 1)
			k--

			// left变成当前最后一位
			if buffer.Len() >= 1 {
				left = buffer.Bytes()[buffer.Len()-1]
			} else {
				left = '0'
			}
		}

		left = num[i]
		// 前导0不写入
		if left == '0' && buffer.Len() < 1 {
			continue
		}

		buffer.WriteByte(num[i])
	}

	lb := buffer.Len()
	if lb < 1 || lb <= k {
		// 删完以后没有数了，或者剩下的数不够删了
		return "0"
	}

	buffer.Truncate(lb - k)
	return buffer.String()
}

func main() {
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
}