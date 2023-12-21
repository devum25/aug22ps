package dynamicprogramming

// Problem Description
// You are given an array A of N integers and three integers B, C, and D.
// You have to find the maximum value of A[i]*B + A[j]*C + A[k]*D, where 1 <= i <= j <= k <= N.
// Input 1:
//  A = [1, 5, -3, 4, -2]
//  B = 2
//  C = 1
//  D = -1
// Input 2:
//  A = [3, 2, 1]
//  B = 1
//  C = -10
//  D = 3

// Example Output
// Output 1:
//  18
// Output 2:
//  -4

// Example Explanation
// Explanation 1:

//  If you choose i = 2, j = 2, and k = 3 then we will get
//  A[2]*B + A[2]*C + A[3]*D = 5*2 + 5*1 + (-3)*(-1) = 10 + 5 + 3 = 18
// func solve(A []int, B int, C int, D int) int {
// }
