package bfs

func Bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	// Write your code here
	adjList := make(map[int32][]int32)

	for _, v := range edges {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	ans := solve(adjList, s)
	res := make([]int32, n-1)

	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	for i := 1; i <= int(n); i++ {
		if _, ok := adjList[int32(i)]; !ok {
			res[i-1] = -1
			continue
		} else if int32(i) == s {
			continue
		} else {
			res[i-1] = int32(ans[int32(i)])
		}
	}

	return res
}

func solve(adjList map[int32][]int32, source int32) map[int32]int32 {
	queue := make([]Node, 0)
	queue = append(queue, Node{Val: source})

	visited := make(map[int32]bool)
	visited[source] = true
	ans := make(map[int32]int32)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		for _, v := range adjList[top.Val] {
			if !visited[v] {
				queue = append(queue, Node{Val: v, Dist: top.Dist + 6})
				ans[v] = top.Dist + 6
				visited[v] = true
			}
		}
	}

	return ans
}

type Node struct {
	Val  int32
	Dist int32
}
