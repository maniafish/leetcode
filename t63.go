package main

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

图示: https://leetcode-cn.com/problems/unique-paths-ii

网格中的障碍物和空位置分别用 1 和 0 来表示。

说明：m 和 n 的值均不超过 100。

示例 1:

输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	l1 := len(obstacleGrid)
	if l1 < 1 {
		return 0
	}

	l2 := len(obstacleGrid[0])
	if l2 < 1 {
		return 0
	}

	if obstacleGrid[0][0] > 0 || obstacleGrid[l1-1][l2-1] > 0 {
		// 起点或终点有障碍物
		return 0
	}

	// 动态规划矩阵，dp[i][j] = n 表示走到(i, j)共有n条路径
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
		for j := 0; j < l2; j++ {
			// 当前点有障碍物
			if obstacleGrid[i][j] > 0 {
				dp[i][j] = 0
				continue
			}

			switch {
			case i == 0 && j == 0:
				// 起点
				dp[i][j] = 1
			case i == 0 && j > 0:
				// 上边框，只能从左边来
				if dp[i][j-1] == 0 {
					// 左边没路了
					dp[i][j] = 0
				} else {
					dp[i][j] = 1
				}
			case i > 0 && j == 0:
				// 左边框，只能从上面来
				if dp[i-1][j] == 0 {
					// 上边没路了
					dp[i][j] = 0
				} else {
					dp[i][j] = 1
				}
			default:
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[l1-1][l2-1]
}
