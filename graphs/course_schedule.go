package graphs

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.
func CanFinish(numCourses int, prerequisites [][]int) bool {

	adjList := make(map[int][]int)

	for _, v := range prerequisites {
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	return bfsCourse(adjList, numCourses)

}

func bfsCourse(adjList map[int][]int, n int) bool {
	indegree := make(map[int]int)

	for i := 0; i < n; i++ {
		indegree[i] = 0
	}

	for k := range adjList {
		for _, v := range adjList[k] {
			indegree[v]++
		}
	}
	queue := make([]int, 0)
	for k, v := range indegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		for _, v := range adjList[temp] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	for _, v := range indegree {
		if v > 0 {
			return false
		}
	}

	return true
}
