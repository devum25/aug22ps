package dfs

import "fmt"

func CountCompleteComponents(n int, edges [][]int) int {
	adjList := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
		adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])
	}
	count := 0
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		if _, ok := adjList[i]; !ok {
			adjList[i] = make([]int, 0)
			count++
			visited[i] = true
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			edge := 0
			nodes := 1
			x := dfs6(adjList, visited, -1, i, &edge, &nodes)
			if x || (nodes == 2 && edge == 1) {
				fmt.Print(nodes)
				count++
			}
		}
	}

	return count
}

func dfs6(adjList map[int][]int, visited map[int]bool, parent, node int, edge *int, nodes *int) bool {
	visited[node] = true

	lst := adjList[node]

	for i := 0; i < len(lst); i++ {
		if !visited[lst[i]] {
			x := *edge
			x++
			*edge = x
			y := *nodes
			y++
			*nodes = y
			if dfs6(adjList, visited, node, lst[i], edge, nodes) {
				return true
			}
		} else if parent != -1 && parent != lst[i] {
			return true
		}
	}

	return false
}
