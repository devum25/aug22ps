package backtracking

// Given a array of integers A of size N and an integer B.
// Return number of non-empty subsequences of A of size B having sum <= 1000.

// Problem Constraints
// 1 <= N <= 20

// 1 <= A[i] <= 1000

// 1 <= B <= N
// Input Format
// The first argument given is the integer array A.
// The second argument given is the integer B.
// Output Format
// Return number of subsequences of A of size B having sum <= 1000.
// Example Input
// Input 1:

//     A = [1, 2, 8]
//     B = 2
// Input 2:

//     A = [5, 17, 1000, 11]
//     B = 4

// Example Output
// Output 1:
// 3
// Output 2:

// 0

// Example Explanation
// Explanation 1:

//  {1, 2}, {1, 8}, {2, 8}
// Explanation 1:

//  No valid subsequence

func Sixlet(A []int, B int) int {
	lst := make([]int, 0)
	ans := make([][]int, 0)
	sum := 0
	currSum := 0
	solve(A, 0, &lst, &ans, B, &sum, &currSum)
	return len(ans)
}

func solve(arr []int, currIdx int, lst *[]int, ans *[][]int, B int, sum *int, currSum *int) {
	if len(*lst) == B && *currSum <= 1000 {
		x := make([]int, len(*lst))
		for i := 0; i < len(x); i++ {
			y := *lst
			x[i] = y[i]
		}
		*ans = append(*ans, x)
		*sum++
		return
	}
	if currIdx == len(arr) {
		return
	}

	*lst = append(*lst, arr[currIdx])
	*currSum = *currSum + arr[currIdx]
	solve(arr, currIdx+1, lst, ans, B, sum, currSum)

	y := *lst
	t := y[len(y)-1]
	y = y[:len(y)-1]
	*lst = y
	*currSum = *currSum - t
	solve(arr, currIdx+1, lst, ans, B, sum, currSum)
}
