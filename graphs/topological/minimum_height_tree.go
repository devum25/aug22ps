package topological

func MinHeightTree(n int, edges [][]int) []int {
	adjList := make(map[int][]int)

	for _, v := range edges {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	if len(adjList) == 0 {
		return []int{0}
	}

	return topo(adjList, n)
}

func topo(adjList map[int][]int, n int) []int {
	indegree := make([]int, n)
	queue := make([]int, 0)
	ans := make([]int, 0)
	for k, v := range adjList {
		indegree[k] = len(v)
		if indegree[k] == 1 {
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		x := len(queue)
		ans = make([]int, 0)
		for i := 0; i < x; i++ {
			ans = append(ans, queue[i])
			for _, v := range adjList[queue[i]] {
				indegree[v]--
				if indegree[v] == 1 {
					queue = append(queue, v)
				}
			}
		}

		queue = queue[x:]
	}

	return ans
}
