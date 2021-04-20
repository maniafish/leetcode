package main

/*
搞个map m存数字和下标，如果对当前数字nums[i]，存在m[target - nums[i]] = j，则返回j, i
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}

		m[v] = i
	}

	return nil
}
