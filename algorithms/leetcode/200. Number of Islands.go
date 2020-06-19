package main

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				count++
				dfs(grid, i, j, visited)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int, visited [][]bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' || visited[i][j] {
		return
	}

	visited[i][j] = true
	dfs(grid, i+1, j, visited)
	dfs(grid, i-1, j, visited)
	dfs(grid, i, j+1, visited)
	dfs(grid, i, j-1, visited)
}
