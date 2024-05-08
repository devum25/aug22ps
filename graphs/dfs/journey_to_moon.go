package dfs

//https://www.hackerrank.com/challenges/journey-to-the-moon/problem?isFullScreen=true

func journeyToMoon(n int32, astronaut [][]int32) int64 {
	// Write your code here
	adjList := make(map[int32][]int32)

	for _, v := range astronaut {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	visited := make(map[int32]bool)
	c := make([]int32, 0)
	var ans int64
	var i int32
	for ; i < n; i++ {
		if !visited[int32(i)] {
			var count int32
			dfs12(adjList, visited, &count, int32(i))
			c = append(c, count)
		}
	}
	if len(c) == 0 || len(c) == 1 {
		return 0
	}
	for i := 0; i < len(c); i++ {
		for j := i + 1; j < len(c); j++ {
			ans += int64(c[i] * c[j])
		}
	}

	return int64(ans)
}

func dfs12(adjList map[int32][]int32, visited map[int32]bool, count *int32, node int32) {
	visited[node] = true
	x := *count
	x++
	*count = x

	lst := adjList[node]

	for i := 0; i < len(lst); i++ {
		if !visited[lst[i]] {
			dfs12(adjList, visited, count, lst[i])
		}
	}
}
