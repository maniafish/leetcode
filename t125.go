package main

import "fmt"

/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
*/

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		a, b := s[i], s[j]
		switch {
		case a >= 'A' && a <= 'Z':
			// 大写转换成小写
			a += 32
		case (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9'):
		default:
			i++
			continue
		}

		switch {
		case b >= 'A' && b <= 'Z':
			// 大写转换成小写
			b += 32
		case (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9'):
		default:
			j--
			continue
		}

		if a != b {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("9 A man, a plan, a canal: Panama9"))
	fmt.Println(isPalindrome("rase a car"))
}
