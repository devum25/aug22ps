package dynamicprogramming

func MaximalSquare(matrix [][]byte) int {
	grid := getIntGrid(matrix)

	dp := make([][]int, len(matrix))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + dp[i][j-1]
			}
		}
	}

	return maxSquare(grid, dp)
}

func maxSquare(grid, dp [][]int) int {
	w, h, ans := 0, 0, 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			h = 1
			w = dp[i][j]
			if h == w {
				ans = max(w*h, ans)
			}
			if i != 0 {
				t := 1

				if (i-t) >= 0 && dp[i-t][j] != 0 {
					w = min(dp[i][j], dp[i-t][j])
				}

				for (i-t) >= 0 && dp[i-t][j] != 0 {
					w = min(w, dp[i-t][j])
					h++
					if h == w {
						ans = max(w*h, ans)
					}
					t++
				}
			}
		}
	}

	return ans
}
