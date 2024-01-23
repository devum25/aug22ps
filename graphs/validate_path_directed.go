package graphs

// Given an directed graph having A nodes labelled from 1 to A containing M edges given by matrix B of size M x 2such that there is a edge directed from nod
// B[i][0] to node B[i][1].
// Find whether a path exists from node 1 to node A.
// Return 1 if path exists else return 0.
// NOTE:
// There are no self-loops in the graph.
// There are no multiple edges between two nodes.
// The graph may or may not be connected.
// Nodes are numbered from 1 to A.
// Your solution will run on multiple test cases. If you are using global variables make sure to clear them.
func ValidatePathDirect(A int, B [][]int) int {
	queue := make([]int, 0)
	hash := make(map[int]bool)
	queue = append(queue, 1)
	hash[1] = true
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]

		for i := 0; i < len(B); i++ {
			if B[i][0] == x {
				if B[i][1] == A {
					return 1
				}
				if !hash[B[i][1]] {
					hash[B[i][1]] = true
					queue = append(queue, B[i][1])
				}
			}
		}
	}

	return 0
}
