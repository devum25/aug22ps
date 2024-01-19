package dynamicprogramming

import (
	"sort"
)

// You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi] represents the width and the height of an envelope.
// One envelope can fit into another if and only if both the width and height of one envelope are greater than the other envelope's width and height.
// Return the maximum number of envelopes you can Russian doll (i.e., put one inside the other).
// Note: You cannot rotate an envelope.

// Input: envelopes = [[5,4],[6,4],[6,7],[2,3]]
// Output: 3
// Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).

func MaxEnvelopes(envelopes [][]int) int {
	m := Matrix(envelopes)
	sort.Sort(m)

	// fmt.Print(m)

	dp := make([]int, len(m))

	dp[0] = 1
	res := 0
	for i := 1; i < len(m); i++ {
		ans := 0
		for j := 0; j < i; j++ {
			if m[j][1] < m[i][1] && m[j][0] != m[i][0] {
				ans = max(dp[j], ans)
			}
		}
		dp[i] = ans + 1
		res = max(res, dp[i])
	}

	return res
}

type Matrix [][]int

func (m Matrix) Less(i, j int) bool {
	for x := range m[i] {
		if m[i][x] == m[j][x] {
			continue
		}
		return m[i][x] < m[j][x]
	}
	return false
}

func (m Matrix) Len() int {
	return len(m)
}

func (m Matrix) Swap(i, j int) {

	m[i], m[j] = m[j], m[i]
	// m = &x
}
