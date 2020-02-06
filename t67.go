package main

import "fmt"

/*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/

func doCarry(prefix, res string) string {
	l := len(prefix)
	for i := 0; i < l; i++ {
		if prefix[l-i-1:l-i] == "1" {
			res = "0" + res
			continue
		}

		return prefix[:l-i-1] + "1" + res
	}

	// 超过最大进位
	return "1" + res
}

func calVal(a, b string, carry bool) (string, bool) {
	if carry {
		if a == "0" && b == "0" {
			return "1", false
		}

		if (a == "0" && b == "1") || (a == "1" && b == "0") {
			return "0", true
		}

		if a == "1" && b == "1" {
			return "1", true
		}
	} else {
		if a == "0" && b == "0" {
			return "0", false
		}

		if (a == "0" && b == "1") || (a == "1" && b == "0") {
			return "1", false
		}

		if a == "1" && b == "1" {
			return "0", true
		}
	}

	return "0", false
}

func addBinary(a string, b string) string {
	l1, l2, n, res, carry, val := len(a), len(b), 0, "", false, "0"
	for ; ; n++ {
		if n >= l1 {
			// 拼接剩下的b
			if carry {
				// 处理进位
				return doCarry(b[:l2-n], res)
			}

			return b[:l2-n] + res
		}

		if n >= l2 {
			// 拼接剩下的a
			if carry {
				// 处理进位
				return doCarry(a[:l1-n], res)
			}

			return a[:l1-n] + res
		}

		val, carry = calVal(a[l1-n-1:l1-n], b[l2-n-1:l2-n], carry)
		res = val + res
	}

	return res
}

func main() {
	fmt.Println(addBinary("1010", "1011"))
}
