package dynamicprogramming

// The demons had captured the princess and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of m x n rooms laid out in a 2D grid. Our valiant knight was initially positioned in the top-left room and must fight his way through dungeon to rescue the princess.
// The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.
// Some of the rooms are guarded by demons (represented by negative integers), so the knight loses health upon entering these rooms; other rooms are either empty (represented as 0) or contain magic orbs that increase the knight's health (represented by positive integers).
// To reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.
// Return the knight's minimum initial health so that he can rescue the princess.
// Note that any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.

// Input: dungeon = [[-2,-3,3],[-5,-10,1],[10,30,-5]]
// Output: 7
// Explanation: The initial health of the knight must be at least 7 if he follows the optimal path: RIGHT-> RIGHT -> DOWN -> DOWN.

func CalculateMinimumHP(dungeon [][]int) int {
	dp := make([][]int, len(dungeon))
	n := len(dungeon[0])

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	x := 0

	if x+dungeon[len(dungeon)-1][len(dungeon[0])-1] < 0 {
		dp[len(dungeon)-1][len(dungeon[0])-1] = -(dungeon[len(dungeon)-1][len(dungeon[0])-1]) + 1
	} else {
		dp[len(dungeon)-1][len(dungeon[0])-1] = 1
	}

	r := len(dungeon) - 2
	c := len(dungeon[0]) - 1
	for r >= 0 {
		temp := dp[r+1][c] - dungeon[r][c]
		if temp > 0 {
			dp[r][c] = temp
		} else {
			dp[r][c] = 1
		}
		r--
	}

	r = len(dungeon) - 1
	c = c - 1
	for c >= 0 {
		temp := dp[r][c+1] - dungeon[r][c]
		if temp > 0 {
			dp[r][c] = temp
		} else {
			dp[r][c] = 1
		}
		c--
	}

	r = len(dungeon) - 2
	c = len(dungeon[0]) - 2

	for i := r; i >= 0; i-- {
		for j := c; j >= 0; j-- {
			temp := min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			if temp > 0 {
				dp[i][j] = temp
			} else {
				dp[i][j] = 1
			}
		}
	}

	return dp[0][0]
}
