package graphs

// Given an undirected graph having A nodes labelled from 1 to A with M edges given in a form of matrix B of size M x 2 where (B[i][0], B[i][1]) represents two nodes B[i][0] and B[i][1] connected by an edge.
// Find whether the graph contains a cycle or not, return 1 if cycle is present else return 0.
// NOTE:
// The cycle must contain atleast three nodes.
// There are no self-loops in the graph.
// There are no multiple edges between two nodes.
// The graph may or may not be connected.
// Nodes are numbered from 1 to A.
// Your solution will run on multiple test cases. If you are using global variables make sure to clear them.

func solve(A int, B [][]int) int {
	adjList := make(map[int][]int)
	visited := make(map[int]bool)
	for _, v := range B {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	for i := 1; i <= A; i++ {
		if !visited[i] {
			if dfs(i, visited, adjList) {
				return 1
			}
		}
	}

	return 0
}

func dfs(node int, visted map[int]bool, adjlist map[int][]int) bool {
	visted[node] = true

	for _, neighbours := range adjlist[node] {
		if !visted[neighbours] {
			if dfs(neighbours, visted, adjlist) {
				return true
			}
			// }else if
		}
	}

	return false
}
