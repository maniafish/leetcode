package main

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
import (
	"fmt"
	"sort"
)

func fourSum(n []int, target int) [][]int {
	sort.Ints(n)
	l := len(n)
	res := make([][]int, 0, l)
	fmt.Println(n)
	for i := 0; i < l-3; i++ {
		if i > 0 && n[i] == n[i-1] {
			continue
		}

		// 遍历接下来的数组元素，每次拿当前元素和后两个元素取sum
		for j0 := i + 1; j0 < l-2; j0++ {
			// 去重
			if j0 > i+1 && n[j0] == n[j0-1] {
				continue
			}

			for j1, j2 := j0+1, l-1; j1 < j2; {
				sum := n[j0] + n[j1] + n[j2]
				switch {
				case sum == target-n[i]:
					res = append(res, []int{n[i], n[j0], n[j1], n[j2]})
					j1, j2 = j1+1, j2-1
					// j1继续右移并去重
					for j1 < j2 && n[j1] == n[j1-1] {
						j1++
					}
					// j2继续左移并去重
					for j1 < j2 && n[j2] == n[j2+1] {
						j2--
					}
				case sum > target-n[i]:
					// 需要缩小sum
					j2--
					// j2继续左移并去重
					for j1 < j2 && n[j2] == n[j2+1] {
						j2--
					}
				default:
					// 需要增大sum
					j1++
					// j1继续右移并去重
					for j1 < j2 && n[j1] == n[j1-1] {
						j1++
					}
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
