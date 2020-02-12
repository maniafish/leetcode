package main

import "bytes"

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:

输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
示例 2:

输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
*/

func convert(s string, numRows int) string {
	/*
		第0行: 0, 2(n-1), 2*2(n-1), ...
		第1行: 1, 2(n-1)-1, 2*(n-1)+1, 4(n-1)-1, 4(n-1)+1, ...
		...
		第k行: k, 2(n-1)-k, 2*(n-1)+k, ... , 2*i*(n-1)-k, 2*i*(n-1)+k
		...
		第n-1行(最后一行): n-1, n-1 + 2(n-1), n-1 + 2*2*(n-1), ... n-1 + 2*i*(n-1)
	*/
	if numRows < 1 {
		return ""
	}

	if numRows == 1 {
		return s
	}

	l := len(s)
	var buffer bytes.Buffer
	for k := 0; k < numRows && k < l; k++ {
		buffer.WriteByte(s[k])
		index := 2 * (numRows - 1)
		if k == 0 || k == numRows-1 {
			for i := 1; k+index*i < l; i++ {
				buffer.WriteByte(s[k+index*i])
			}
		} else {
			for i := 1; index*i-k < l; i++ {
				buffer.WriteByte(s[index*i-k])
				if index*i+k < l {
					buffer.WriteByte(s[index*i+k])
				}
			}
		}
	}

	return buffer.String()
}
