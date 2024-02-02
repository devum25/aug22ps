package graphs

func ParallelCourse(n int, relations [][]int, time []int) int {

	adjList := make(map[int][]int)

	for i := 1; i <= n; i++ {
		adjList[i] = make([]int, 0)
	}

	for _, v := range relations {
		adjList[v[0]] = append(adjList[v[0]], v[1])
	}

	prereq := make(map[int][]int)

	for _, v := range relations {
		prereq[v[1]] = append(prereq[v[1]], v[0])
	}

	return topologicalSort(adjList, prereq, time, n)
}

func topologicalSort(adjList, prereq map[int][]int, time []int, n int) int {
	indegree := make(map[int]int)

	for i := 1; i <= n; i++ {
		indegree[i] = 0
	}

	for i := 1; i <= n; i++ {
		for _, v := range adjList[i] {
			indegree[v]++
		}
	}
	queue := make([]int, 0)
	topoorder := make([]int, 0)
	for k, _ := range adjList {
		if indegree[k] == 0 {
			queue = append(queue, k)
			topoorder = append(topoorder, k)
		}
	}

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		for _, v := range adjList[temp] {
			indegree[v]--
			if indegree[v] == 0 {
				topoorder = append(topoorder, v)
				queue = append(queue, v)
			}
		}

	}

	dp := make([]int, n+1)

	for _, v := range topoorder {
		maxPrereqTime := 0
		for _, prevCourse := range prereq[v] {
			if dp[prevCourse] > maxPrereqTime {
				maxPrereqTime = dp[prevCourse]
			}
		}
		dp[v] = maxPrereqTime + time[v-1]
	}
	ans := 0
	for i := 1; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}

	return ans
}
