package graphs

func FindCenter(edges [][]int) int {
	adjList := make(map[int][]int)

	for _, v := range edges {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	return bfs4(adjList)
}

func bfs4(adjList map[int][]int) int {
	indegree := make(map[int]int)
	queue := make([]int, 0)
	for k, v := range adjList {
		indegree[k] = len(v)

		if indegree[k] == 1 {
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		delete(adjList, top)
	}

	for k := range adjList {
		return k
	}

	return -1
}
