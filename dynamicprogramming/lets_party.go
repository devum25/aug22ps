package dynamicprogramming

// In Danceland, one person can party either alone or can pair up with another person.
// Can you find in how many ways they can party if there are A people in Danceland?
// Note: Return your answer modulo 10003, as the answer can be large.

//Return an integer denoting the number of ways people of Danceland can party.

// Example Input
// Input 1:A = 3
// Input 2: A = 5
// Example Output
// Output 1:4
// Output 2: 26
// Example Explanation

// Explanation 1:

//  Let suppose three people are A, B, and C. There are only 4 ways to party
//  (A, B, C) All party alone
//  (AB, C) A and B party together and C party alone
//  (AC, B) A and C party together and B party alone
//  (BC, A) B and C party together and A
//  here 4 % 10003 = 4, so answer is 4.

// Explanation 2:

//  Number of ways they can party are: 26.
func LetsParty(A int) int {
	dp := make([]int, A+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = 0
	}

	return party(A, &dp)
}

func party(n int, dp *[]int) int {
	(*dp)[1] = 1

	for i := 1; i <= n; i++ {
		(*dp)[i] = i
		for j := i + 1; j <= n; j++ {
			(*dp)[i] = (*dp)[j-1] + 1
		}
	}

	return (*dp)[n]
}
