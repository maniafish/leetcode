package main

/*
在二维网格 grid 上，有 4 种类型的方格：

1 表示起始方格。且只有一个起始方格。
2 表示结束方格，且只有一个结束方格。
0 表示我们可以走过的空方格。
-1 表示我们无法跨越的障碍。
返回在四个方向（上、下、左、右）上行走时，从起始方格到结束方格的不同路径的数目，每一个无障碍方格都要通过一次。

示例 1：

输入：
[
	[1,0,0,0],
	[0,0,0,0],
	[0,0,2,-1]
]

输出：2
解释：我们有以下两条路径：
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
示例 2：

输入：
[
	[1,0,0,0],
	[0,0,0,0],
	[0,0,0,2]
]

输出：4
解释：我们有以下四条路径：
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
示例 3：

输入：
[
	[0,1],
	[2,0]
]
输出：0
解释：
没有一条路能完全穿过每一个空的方格一次。
请注意，起始和结束方格可以位于网格中的任意位置。


提示：

1 <= grid.length * grid[0].length <= 20
*/

func uniquePathsIII(grid [][]int) int {
	l1 := len(grid)
	if l1 < 1 {
		return 0
	}

	l2 := len(grid[0])
	if l2 < 1 {
		return 0
	}

	// 起点坐标；终点坐标；需要遍历的格子数
	si, sj, ei, ej, todo := -1, -1, -1, -1, 0
	// 找到起点和终点
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if grid[i][j] == 1 {
				// 多个起点
				if si > 0 || sj > 0 {
					return 0
				}

				si, sj = i, j
				continue
			}

			if grid[i][j] == 2 {
				// 多个终点
				if ei > 0 || ej > 0 {
					return 0
				}

				ei, ej = i, j
				continue
			}

			if grid[i][j] == 0 {
				todo++
			}
		}
	}

	// 没有起点或终点
	if si < 0 || sj < 0 || ei < 0 || ej < 0 {
		return 0
	}

	// 从si, sj开始进行深度优先遍历
	return dfs(grid, si, sj, l1-1, l2-1, todo)
}

func dfs(grid [][]int, i, j, l1, l2, todo int) int {
	if grid[i][j] < 0 {
		return 0
	}

	if grid[i][j] == 2 {
		if todo == 0 {
			return 1
		}

		return 0
	}

	if grid[i][j] == 0 {
		todo--
	}

	// 标记为已经走过的
	grid[i][j] = 3
	res := 0
	if i > 0 && (grid[i-1][j] == 0 || grid[i-1][j] == 2) {
		res += dfs(grid, i-1, j, l1, l2, todo)
	}

	if j > 0 && (grid[i][j-1] == 0 || grid[i][j-1] == 2) {
		res += dfs(grid, i, j-1, l1, l2, todo)
	}

	if i < l1 && (grid[i+1][j] == 0 || grid[i+1][j] == 2) {
		res += dfs(grid, i+1, j, l1, l2, todo)
	}

	if j < l2 && (grid[i][j+1] == 0 || grid[i][j+1] == 2) {
		res += dfs(grid, i, j+1, l1, l2, todo)
	}

	// 回溯
	grid[i][j] = 0
	todo++
	return res
}
