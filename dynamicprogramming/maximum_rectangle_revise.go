package dynamicprogramming

import "strconv"

//Given a rows x cols binary matrix filled with 0's and 1's,
// find the largest rectangle containing only 1's and return its area

// Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]

func MaximalRectangleRevision(matrix [][]byte) int {

	grid := getIntGrid(matrix)

	dp := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
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

	return getMaxRectangle(grid, dp)

}

func getMaxRectangle(grid, dp [][]int) int {
	w, h, ans := 0, 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			h = 1
			w = dp[i][j]
			ans = max(w*h, ans)

			if i != 0 {
				t := 1

				for (i-t) >= 0 && dp[i-t][j] != 0 {
					w = min(dp[i-t-1][j], dp[i-t][j])
					h++
					ans = max(w*h, ans)
					t++
				}
			}
		}
	}

	return ans
}

func getIntGrid(matrix [][]byte) [][]int {

	res := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		res[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			res[i][j] = getinteger(matrix[i][j])
		}
	}

	return res
}

func getinteger(c byte) int {
	s := string(c)
	v, _ := strconv.Atoi(s)
	return v
}
