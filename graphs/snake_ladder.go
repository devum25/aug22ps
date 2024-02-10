package graphs

// SNAKES AND LADDERS

// You are given an n x n integer matrix board where the cells are labeled from 1 to n2 in a Boustrophedon style starting from the bottom left of the board (i.e. board[n - 1][0]) and alternating direction each row.

// You start on square 1 of the board. In each move, starting from square curr, do the following:

// Choose a destination square next with a label in the range [curr + 1, min(curr + 6, n2)].
// This choice simulates the result of a standard 6-sided die roll: i.e., there are always at most 6 destinations, regardless of the size of the board.
// If next has a snake or ladder, you must move to the destination of that snake or ladder. Otherwise, you move to next.
// The game ends when you reach the square n2.
// A board square on row r and column c has a snake or ladder if board[r][c] != -1. The destination of that snake or ladder is board[r][c]. Squares 1 and n2 do not have a snake or ladder.

// Note that you only take a snake or ladder at most once per move. If the destination to a snake or ladder is the start of another snake or ladder, you do not follow the subsequent snake or ladder.

// For example, suppose the board is [[-1,4],[-1,3]], and on the first move, your destination square is 2. You follow the ladder to square 3, but do not follow the subsequent ladder to 4.
// Return the least number of moves required to reach the square n2. If it is not possible to reach the square, return -1.
// https://www.youtube.com/watch?v=6lH4nO3JfLk
func SnakesAndLadders(board [][]int) int {
	board = reverseBoard(board)

	queue := make([]Board, 0)
	n := len(board)
	queue = append(queue, Board{Val: 1, Dist: 0})
	visited := make(map[int]bool)

	visited[1] = true

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		val := 0
		for k := 1; k <= 6; k++ {
			val = top.Val + k
			r, c := getCoordinates(val, n)
			if board[r][c] != -1 {
				val = board[r][c]
			}
			if val == n*n {
				return top.Dist + 1
			}
			if !visited[val] {
				visited[val] = true
				queue = append(queue, Board{Dist: top.Dist + 1, Val: val})
			}
		}
	}

	return 0
}

func getCoordinates(val int, n int) (int, int) {
	r := (val - 1) / n
	c := (val - 1) % n
	if r%2 != 0 {
		c = n - 1 - c
	}

	return r, c
}

func reverseBoard(board [][]int) [][]int {
	newBoard := make([][]int, len(board))

	for i := 0; i < len(board); i++ {
		newBoard[i] = make([]int, len(board[i]))
	}
	row := 0
	for i := len(board) - 1; i >= 0; i-- {
		for j := 0; j < len(board[i]); j++ {
			newBoard[row][j] = board[i][j]
		}
		row++
	}

	return newBoard
}

type Board struct {
	Val  int
	Dist int
}
