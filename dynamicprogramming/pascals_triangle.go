package dynamicprogramming

// Given an integer numRows, return the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

// Example 1:
// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
// Example 2:
// Input: numRows = 1
// Output: [[1]]

func PascalTriangle(numRows int) [][]int {
	if numRows == 1 {
		x := make([][]int, 1)
		x = append(x, []int{1})
		return x
	}
	dp := make([][]int, numRows)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = 1
	dp[1][0] = 1
	dp[1][1] = 1
	for i := 2; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j == 0 || j == len(dp[i])-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}

	return dp
}
