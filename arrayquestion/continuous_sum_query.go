package arrayquestion

// There are A beggars sitting in a row outside a temple. Each beggar initially has an empty pot. When the devotees come to the temple, they donate some amount of coins to these beggars. Each devotee gives a fixed amount of coin(according to their faith and ability) to some K beggars sitting next to each other.

// Given the amount P donated by each devotee to the beggars ranging from L to R index, where 1 <= L <= R <= A, find out the final amount of money in each beggar's pot at the end of the day, provided they don't fill their pots by any other means.
// For ith devotee B[i][0] = L, B[i][1] = R, B[i][2] = P, given by the 2D array B

// Problem Constraints
// 1 <= A <= 2 * 105
// 1 <= L <= R <= A
// 1 <= P <= 103
// 0 <= len(B) <= 105

// Input Format
// The first argument is a single integer A.
// The second argument is a 2D integer array B.

// Output Format
// Return an array(0 based indexing) that stores the total number of coins in each beggars pot.

// Example Input
// Input 1:-
// A = 5
// B = [[1, 2, 10], [2, 3, 20], [2, 5, 25]]

func ContinuousSumSubQuery(A int, B [][]int) []int {
	res := make([]int, A)

	for i := 0; i < len(B); i++ {
		s := B[i][0]
		e := B[i][1]
		val := B[i][2]
		s = s - 1
		e = e - 1

		res[s] += val
		if (e + 1) < len(res) {
			res[e+1] += -val
		}
	}

	pf := make([]int, len(res))
	for i := 0; i < len(res); i++ {
		if i == 0 {
			pf[i] = res[i]
		} else {
			pf[i] = pf[i-1] + res[i]
		}
	}

	return pf
}
