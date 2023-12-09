package backtracking

// Q3. Rat In a Maze
// Problem Description
// Given a grid A, a rat is at position (1, 1). He wants to go to the position (n, n) where n is size of the square matrix A.
// 1 represents a traversable cell and 0 represents a non-traversable cell. Rat can only move right or down.
// Return a path that the rat can take.

// Problem Constraints
// 1 <= |A| <= 4

// Input Format
// First and only argument is the vector of vectors A.

// Output Format
// Return a vector of vectors that denotes a path that can be taken.

// Example Input
// Input 1:
// A = [   [1, 0]
//         [1, 1]
//     ]
// Input 2:
// A = [    [1, 1, 1]
//          [1, 0, 1]
//          [0, 0, 1]
//     ]

// Example Output
// Output 1:
// [   [1, 0]
//     [1, 1]
// ]
// Output 2:

// [    [1, 1, 1]
//      [0, 0, 1]
//      [0, 0, 1]
// ]

func MazeSolver(A [][]int) [][]int {

	n := len(A)
	m := len(A[0])
	ans := make([][]int, n)
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, m)
	}
	maze(A, 0, 0, &ans, n, m)
	return ans
}

func maze(matrix [][]int, x, y int, ans *[][]int, N, M int) bool {
	if x >= N || x < 0 || y >= M || y < 0 {
		return false
	}
	if matrix[x][y] == 1 {
		(*ans)[x][y] = 1
	} else {
		return false
	}

	if x == N-1 && y == M-1 {
		return true
	}

	// Move right
	if maze(matrix, x, y+1, ans, N, M) {
		return true
	}

	// Move down
	if maze(matrix, x+1, y, ans, N, M) {
		return true
	}

	// If neither right nor down leads to the destination, backtrack
	(*ans)[x][y] = 0
	return false
}
