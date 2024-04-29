package dfs

func IsBipartite1(graph [][]int) bool {
	adjList := make(map[int][]int)

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			adjList[i] = append(adjList[i], graph[i][j])
			// adjList[graph[i][j]] = append(adjList[graph[i][j]], i)
		}
	}

	visited := make(map[int]bool)
	color := make(map[int]int)
	for k := range adjList {
		if !visited[k] {
			color[k] = 1
			if !dfsBipartite(adjList, visited, color, -1, k) {
				return false
			}
		}
	}

	return true
}

func dfsBipartite(adjList map[int][]int, visited map[int]bool, color map[int]int, parent int, node int) bool {
	visited[node] = true

	lst := adjList[node]

	for i := 0; i < len(lst); i++ {
		if !visited[lst[i]] {
			color[lst[i]] = 3 - color[node]
			if !dfsBipartite(adjList, visited, color, node, lst[i]) {
				return false
			}
		} else if parent != lst[i] && color[lst[i]] == color[node] {
			return false
		}
	}

	return true
}
