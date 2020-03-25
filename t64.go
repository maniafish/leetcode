package main

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/

/*
题解：
动态规划，dp[i][j]为终点为(i, j)的最小路径
*/

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minPathSum(grid [][]int) int {
	l1 := len(grid)
	if l1 < 1 {
		return 0
	}

	l2 := len(grid[0])
	if l2 < 1 {
		return 0
	}

	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		if len(grid[i]) != l2 {
			return 0
		}

		dp[i] = make([]int, l2)
		for j := 0; j < l2; j++ {
			switch {
			case i == 0 && j == 0:
				// 起点
				dp[i][j] = grid[i][j]
			case i == 0 && j > 0:
				// 上边框，只能从左边走到当前点
				dp[i][j] = dp[i][j-1] + grid[i][j]
			case i > 0 && j == 0:
				// 左边框，只能从上边走到当前点
				dp[i][j] = dp[i-1][j] + grid[i][j]
			default:
				// 上边和左边都能到当前节点
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[l1-1][l2-1]
}
