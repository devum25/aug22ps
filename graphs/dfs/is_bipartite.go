package dfs

func IsBipartite(graph [][]int) bool {
	adjList := make(map[int][]Node)

	for i := 0; i < len(graph); i++ {
		adjList[i] = make([]Node, 0)
		for j := 0; j < len(graph[i]); j++ {
			adjList[i] = append(adjList[i], Node{Val: graph[i][j]})
			adjList[graph[i][j]] = append(adjList[graph[i][j]], Node{Val: i})
		}
	}

	return bfs(adjList)
}

func bfs(adjList map[int][]Node) bool {
	queue := make([]Node, 0)

	queue = append(queue, Node{Color: 1, Val: 0})

	visited := make(map[int]bool)

	visited[0] = true

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		color := temp.Color

		for _, v := range adjList[temp.Val] {
			if !visited[v.Val] {
				visited[v.Val] = true
				if color == 1 {
					queue = append(queue, Node{Val: v.Val, Color: 2})
				} else {
					queue = append(queue, Node{Val: v.Val, Color: 1})
				}
			} else if v.Color == color {
				return false
			}
		}

	}

	return true

}

type Node struct {
	Color int
	Val   int
}
