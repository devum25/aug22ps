package graphs

// There is a rectangle with left bottom as (0, 0) and right up as (x, y).
// There are N circles such that their centers are inside the rectangle.
// Radius of each circle is R. Now we need to find out if it is possible that we can move from (0, 0) to (x, y) without touching any circle.
// Note : We can move from any cell to any of its 8 adjecent neighbours and we cannot move outside the boundary of the rectangle at any point of time.
func ValidPath1(n int, edges [][]int, source int, destination int) bool {
	if len(edges) == 0 && source == destination {
		return true
	}
	visited := make(map[int]bool)

	adjList := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		if _, ok := adjList[edges[i][0]]; !ok {
			x := make([]int, 0)
			x = append(x, edges[i][1])
			adjList[edges[i][0]] = x
			continue
		}
		t := adjList[edges[i][0]]
		t = append(t, edges[i][1])
		adjList[edges[i][0]] = t
	}

	queue := make([]int, 0)

	queue = append(queue, source)
	visited[source] = true

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		lst := adjList[temp]

		for i := 0; i < len(lst); i++ {
			if lst[i] == destination {
				return true
			}
			if !visited[lst[i]] {
				visited[lst[i]] = true
				queue = append(queue, lst[i])
			}
		}
	}

	return false
}
