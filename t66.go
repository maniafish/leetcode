package main

import "fmt"

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/

func plusOne(digits []int) []int {
	l := len(digits)
	for i := 0; i < l; i++ {
		// 进位
		if digits[l-i-1] == 9 {
			digits[l-i-1] = 0
			continue
		}

		digits[l-i-1] += 1
		return digits
	}

	// 进位超过最高位数
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}
