package dfs

func CycleInDirectedGraph(A int, B [][]int) int {
	adjList := make(map[int][]int)

	for _, v := range B {
		adjList[v[0]] = append(adjList[v[0]], v[1])
	}

	visited := make(map[int]bool)
	rec := make([]bool, A+1)

	for i := 1; i <= A; i++ {
		if !visited[i] {
			if dfsDirectedCycle(i, visited, rec, adjList) {
				return 1
			}
		}
	}

	return 0

}

func dfsDirectedCycle(node int, visited map[int]bool, rec []bool, adjList map[int][]int) bool {
	visited[node] = true
	rec[node] = true
	for _, v := range adjList[node] {
		if !visited[v] {
			if dfsDirectedCycle(v, visited, rec, adjList) {
				return true
			}
		} else if rec[v] {
			return true
		}
	}
	rec[node] = false
	return false
}
