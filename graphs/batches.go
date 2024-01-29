package graphs

// A students applied for admission in IB Academy.
// An array of integers B is given representing the strengths of A people i.e. B[i] represents the strength of ith student.
// Among the A students some of them knew each other.
// A matrix C of size M x 2 is given which represents relations where ith relations depicts that C[i][0] and C[i][1] knew each other.
// All students who know each other are placed in one batch.
// Strength of a batch is equal to sum of the strength of all the students in it.
// Now the number of batches are formed are very much, it is impossible for IB to handle them.
// So IB set criteria for selection: All those batches having strength at least D are selected.
// Find the number of batches selected.

// NOTE: If student x and student y know each other, student y and z know each other then student x and student z will also know each other.
func SelectBatchs(A int, B []int, C [][]int, D int) int {
	adjList := make(map[int][]int)

	for _, v := range C {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	return bfs(adjList, D, B, A)
}

func bfs(adjList map[int][]int, cutoff int, Strength []int, n int) int {
	visited := make(map[int]bool)
	res := 0
	queue := make([]int, 0)
	for i := 1; i <= n; i++ {
		if !visited[i] {
			s := Strength[i-1]
			queue = append(queue, i)
			visited[i] = true
			for len(queue) > 0 {
				temp := queue[0]
				queue = queue[1:]

				for _, v := range adjList[temp] {
					if !visited[v] {
						queue = append(queue, v)
						s += Strength[v-1]
						visited[v] = true
					}
				}
			}
			if s >= cutoff {
				res++
			}
		}
	}
	return res
}
