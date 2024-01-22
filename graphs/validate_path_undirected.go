package graphs

// There is a bi-directional graph with n vertices, where each vertex is labeled from 0 to n - 1 (inclusive).
// The edges in the graph are represented as a 2D integer array edges, where each edges[i] = [ui, vi]
// denotes a bi-directional edge between vertex ui and vertex vi. Every vertex pair is connected by at most one edge, and no vertex has an edge to itself.
// You want to determine if there is a valid path that exists from vertex source to vertex destination.
// Given edges and the integers n, source, and destination, return true if there is a valid path from source to destination, or false otherwise.

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	if len(edges) == 0 && source == 0 && destination == 0 {
		return true
	}
	queue := make([]int, 0)
	queue = append(queue, source)
	hash := make(map[int]bool)
	hash[source] = true
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]

		for i := 0; i < len(edges); i++ {
			if edges[i][0] == x {
				if edges[i][1] == destination {
					return true
				}
				if !hash[edges[i][1]] {
					queue = append(queue, edges[i][1])
					hash[edges[i][1]] = true
				}
			} else if edges[i][1] == x {
				if edges[i][0] == destination {
					return true
				}
				if !hash[edges[i][0]] {
					queue = append(queue, edges[i][0])
					hash[edges[i][0]] = true
				}
			}
		}
	}

	return false
}
