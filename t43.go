package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
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

func multiply(num1 string, num2 string) string {
	// 把num2 * (1 ~ 9)的情况都算出来
	if len(num1) == 0 || len(num2) == 0 || num1 == "0" || num2 == "0" {
		return "0"
	}

	temp := make([]string, 10)
	temp[0] = "0"
	for i := 1; i < 10; i++ {
		temp[i] = addStrings(temp[i-1], num2)
	}

	l, res := len(num1), "0"
	for n := 0; n < l; n++ {
		index, _ := strconv.Atoi(num1[l-n-1 : l-n])
		res = addStrings(res, temp[index]+strings.Repeat("0", n))
	}

	return res
}

func main() {
	fmt.Println(multiply("123", "456"))
}
