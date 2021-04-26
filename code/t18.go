package main

import (
	"sort"
)

/*
类似三数求和的思路，只是中间index指针从一个变成两个，时间复杂度上升一个量级
*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	// 最终输出结果，头指针，头指针中间值，尾指针，数组末尾，中间计算结果
	result, start, startTemp, end, last, temp1, temp2 := make([][]int, 0), 0, 0, 0, len(nums)-1, 0, 0
	for index1, index2 := 1, 2; index1 < last-1; index1++ {
		startTemp = 0
		if index1 > 1 && nums[index1] == nums[index1-1] {
			// 重复了两次，不用算了
			if nums[index1] == nums[index1-2] {
				continue
			}

			startTemp = index1 - 1
		}

		for index2 = index1 + 1; index2 < last; index2++ {
			// 中间指针和重复
			if index2 > index1+1 && nums[index1]+nums[index2] == temp1 {
				continue
			}

			start, end, temp1 = startTemp, last, nums[index1]+nums[index2]
			for start < index1 && index2 < end {
				if start > 0 && nums[start] == nums[start-1] {
					start++
					continue
				}

				if end < last && nums[end] == nums[end+1] {
					end--
					continue
				}

				temp2 = nums[start] + nums[end] + temp1
				if temp2 == target {
					result = append(result, []int{nums[start], nums[index1], nums[index2], nums[end]})
					start++
					end--
				} else if temp2 < target {
					start++
				} else {
					end--
				}
			}
		}
	}

	return result
}
