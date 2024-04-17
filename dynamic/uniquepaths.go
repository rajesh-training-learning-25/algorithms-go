// See https://leetcode.com/problems/unique-paths/
// author: Rares Mateizer (https://github.com/rares985)
package dynamic

// UniquePaths implements the solution to the "Unique Paths" problem
func UniquePaths(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	grid := make([][]int, m+1)
	for i := range grid {
		grid[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		grid[i][0] = 1
	}

	for j := 0; j < n; j++ {
		grid[0][j] = 1
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]
}
