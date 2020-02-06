package main

import "fmt"

/*
对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。

给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。



示例 1：

输入：A = [1,2,0,0], K = 34
输出：[1,2,3,4]
解释：1200 + 34 = 1234
示例 2：

输入：A = [2,7,4], K = 181
输出：[4,5,5]
解释：274 + 181 = 455
示例 3：

输入：A = [2,1,5], K = 806
输出：[1,0,2,1]
解释：215 + 806 = 1021
示例 4：

输入：A = [9,9,9,9,9,9,9,9,9,9], K = 1
输出：[1,0,0,0,0,0,0,0,0,0,0]
解释：9999999999 + 1 = 10000000000


提示：

1 <= A.length <= 10000
0 <= A[i] <= 9
0 <= K <= 10000
如果 A.length > 1，那么 A[0] != 0
*/

func calVal(a, b int, carry bool) (int, bool) {
	sum := a + b
	if carry {
		sum += 1
	}

	if sum >= 10 {
		return sum - 10, true
	}

	return sum, false
}

func setPrefix(k []int, carry bool) []int {
	l := len(k)
	res := make([]int, l)
	for i := 0; i < l; i++ {
		if carry {
			// 继续进位
			if k[i] == 9 {
				res[l-i-1] = 0
				continue
			}

			res[l-i-1] = k[i] + 1
			carry = false
		} else {
			res[l-i-1] = k[i]
		}
	}

	// 进位超过最大位数
	if carry {
		return append([]int{1}, res...)
	}

	return res
}

func addToArrayForm(A []int, K int) []int {
	k := make([]int, 0, 5)
	for ; K != 0; K = K / 10 {
		// 反向数组存放K
		k = append(k, K%10)
	}

	l1, l2, carry := len(A), len(k), false
	for i := 0; i < l2; i++ {
		if i >= l1 {
			// 拼接剩下的k
			return append(setPrefix(k[i:], carry), A...)
		}

		A[l1-i-1], carry = calVal(A[l1-i-1], k[i], carry)
	}

	if carry {
		// 处理A剩下的进位
		for i := l2; i < l1; i++ {
			// 继续进位
			if A[l1-i-1] == 9 {
				A[l1-i-1] = 0
				continue
			}

			A[l1-i-1] += 1
			return A

		}

		// 进位超过最大位数
		return append([]int{1}, A...)
	}

	return A
}

func main() {
	fmt.Println(addToArrayForm([]int{2, 1, 5}, 806))
}
