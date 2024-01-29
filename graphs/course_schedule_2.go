package graphs

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

func FindOrder(numCourses int, prerequisites [][]int) []int {
	adjList := make(map[int][]int)

	for _, v := range prerequisites {
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	return bfsOrderCourse(adjList, numCourses)
}

func bfsOrderCourse(adjList map[int][]int, n int) []int {
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
	order := make([]int, 0)
	for k, v := range indegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		order = append(order, temp)

		for _, v := range adjList[temp] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	for _, v := range indegree {
		if v > 0 {
			return []int{}
		}
	}

	return order
}
