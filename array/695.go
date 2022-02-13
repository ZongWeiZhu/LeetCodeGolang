package array


func dfs(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	directions := [][]int {
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}
	ans := 1
	grid[i][j] = 0
	for _, direction := range directions {
		ans += dfs(grid, i+direction[0], j+direction[1])
	}
	return ans
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			maxArea = max(maxArea, dfs(grid, i, j))
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
