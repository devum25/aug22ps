package dynamicprogramming

import "math"

func CalculateMinimumHPRecusive(dungeon [][]int) int {
	dp := make([][]int, len(dungeon))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(dungeon[i]))
	}

	for i := 0; i < len(dungeon); i++ {
		for j := 0; j < len(dungeon[0]); j++ {
			dp[i][j] = -1
		}
	}

	x := calculate(dp, dungeon, 0, 0)
	return x
}

func calculate(dp, dungeon [][]int, i, j int) int {
	if i >= len(dungeon) || j >= len(dungeon[0]) {
		return math.MaxInt
	}

	if i == len(dungeon)-1 && j == len(dungeon[0])-1 {
		x := dungeon[i][j]
		if x < 0 {
			return -(x) + 1
		} else {
			return 1
		}
	}

	if dp[i][j] == -1 {

		temp := min(calculate(dp, dungeon, i, j+1), calculate(dp, dungeon, i+1, j))
		c := temp - dungeon[i][j]
		if c < 0 {
			dp[i][j] = 1
		} else {
			dp[i][j] = c
		}

	}

	return dp[i][j]
}
