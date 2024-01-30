package graphs

// A country consist of N cities connected by N - 1 roads. King of that country want to construct maximum number of roads such that the new country formed remains bipartite country.
// Bipartite country is a country, whose cities can be partitioned into 2 sets in such a way, that for each road (u, v) that belongs to the country, u and v belong to different sets. Also, there should be no multiple roads between two cities and no self loops.
// Return the maximum number of roads king can construct. Since the answer could be large return answer % 109 + 7.

func ConstructMaxRoads(A int, B [][]int) int {
	adjList := make(map[int][]int)

	for _, v := range B {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	return roads(adjList, A)
}

func roads(adjList map[int][]int, n int) int {
	lstColor := make([]int, n+1)
	queue := make([]Node, 0)
	visited := make(map[int]bool)
	for i := 1; i <= n; i++ {
		if !visited[i] {
			queue = append(queue, Node{Color: 1, Val: i})
			visited[i] = true
			lstColor[i] = 1
			for len(queue) > 0 {
				temp := queue[0]
				queue = queue[1:]

				for _, v := range adjList[temp.Val] {
					if !visited[v] {
						visited[v] = true
						lstColor[v] = 3 - temp.Color
						queue = append(queue, Node{Color: 3 - temp.Color, Val: v})
					} else if temp.Color == lstColor[v] {
						return 0
					}
				}
			}

		}
	}

	x := 0
	y := 0

	for i := 1; i < len(lstColor); i++ {
		if lstColor[i] == 1 {
			x++
		} else if lstColor[i] == 2 {
			y++
		}
	}

	return ((x * y) - (n - 1)) % 1000000007
}
