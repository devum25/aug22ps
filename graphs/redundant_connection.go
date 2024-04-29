package graphs

func FindRedundantConnection(edges [][]int) []int {

	adjList := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		if _, ok := adjList[edges[i][0]]; !ok {
			adjList[edges[i][0]] = make([]int, 0)
		}
		adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])

		if _, ok := adjList[edges[i][1]]; !ok {
			adjList[edges[i][1]] = make([]int, 0)
		}
		adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])
	}

	visited := make(map[int]bool)
	ans := make([][]int, 0)

	for k, _ := range adjList {
		if !visited[k] {
			dfs(adjList, visited, -1, k, &ans)
		}
	}

	return ans[len(ans)-1]
}

func dfs(adjList map[int][]int, visited map[int]bool, parent int, node int, ans *[][]int) {
	visited[node] = true

	lst := adjList[node]

	for i := 0; i < len(lst); i++ {
		if !visited[lst[i]] {
			dfs(adjList, visited, node, lst[i], ans)
		} else if visited[lst[i]] && parent != -1 && parent != lst[i] {
			temp := make([]int, 0)
			temp = append(temp, node)
			temp = append(temp, parent)
			*ans = append(*ans, temp)
			break
		}
	}
}
