package main

import (
	"fmt"
	"strconv"
)

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
*/

func doCarry(prefix, res string) string {
	l := len(prefix)
	for n := 0; n < l; n++ {
		num := prefix[l-n-1 : l-n]
		if num == "9" {
			// 进位
			res = "0" + res
			continue
		}

		numInt, _ := strconv.Atoi(num)
		numInt += 1
		res = prefix[:l-n-1] + strconv.Itoa(numInt) + res
		return res
	}

	// 超过最高进位
	return "1" + res
}

func calVal(num1, num2 string, carry bool) (int, bool) {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	sum := n1 + n2
	if carry {
		sum += 1
	}

	if sum/10 > 0 {
		return sum % 10, true
	}

	return sum % 10, false
}

func addStrings(num1 string, num2 string) string {
	l1, l2, n, carry, res, val := len(num1), len(num2), 0, false, "", 0
	for {
		if n >= l1 {
			if carry {
				// 处理剩下的进位
				return doCarry(num2[:l2-n], res)
			}

			// 拼上剩下的num2
			return num2[:l2-n] + res
		}

		if n >= l2 {
			if carry {
				// 处理剩下的进位
				return doCarry(num1[:l1-n], res)
			}

			// 拼上剩下的num1
			return num1[:l1-n] + res
		}

		val, carry = calVal(num1[l1-n-1:l1-n], num2[l2-n-1:l2-n], carry)
		res = strconv.Itoa(val) + res
		n += 1
	}

	return res
}

func main() {
	fmt.Println(addStrings("123", "8888"))
}
