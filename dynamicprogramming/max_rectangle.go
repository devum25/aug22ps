package dynamicprogramming

import "strconv"

// 85. Maximal Rectangle
// Hard
// 9.6K
// 159
// Companies
// Given a rows x cols binary matrix filled with 0's and 1's,
// find the largest rectangle containing only 1's and return its area.

// Input: matrix = [["1","0","1","0","0"]
// 				 ["1","0","1","1","1"]
// 				 ["1","1","1","1","1"]
// 				 ["1","0","0","1","0"]]
// Output: 6
// Explanation: The maximal rectangle is shown in the above picture.

func MaximalRectangle(matrix [][]byte) int {
	grid := getgrid(matrix)

	dp := make([][]int, len(grid))
	n := len(grid[0])

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			} else if j > 0 {
				dp[i][j] = 1 + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = 1
			}
		}
	}

	x := max_rectangle(dp)

	return x
}

func max_rectangle(dp [][]int) int {
	w, h, ans := 0, 0, 0

	for i := 0; i < len(dp); i++ {

		for j := 0; j < len(dp[0]); j++ {
			if dp[i][j] == 0 {
				continue
			}
			h = 1
			w = dp[i][j]
			ans = max(w*h, ans)
			if i != 0 {
				t := 1
				if (i-t) >= 0 && dp[i-t][j] != 0 {
					w = min(dp[i][j], dp[i-t][j])
				}
				for (i-t) >= 0 && dp[i-t][j] != 0 {
					if dp[i-t][j] < w {
						w = dp[i-t][j]
					}
					h++
					ans = max(w*h, ans)
					t++
				}
			}
		}
	}

	return ans
}

func getgrid(matrix [][]byte) [][]int {

	grid := make([][]int, len(matrix))

	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = getint(matrix[i][j])
		}
	}

	return grid
}

func getint(c byte) int {
	s := string(c)
	v, _ := strconv.Atoi(s)
	return v
}
