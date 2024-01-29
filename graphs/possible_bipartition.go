package graphs

// We want to split a group of n people (labeled from 1 to n) into two groups of any size. Each person may dislike some other people, and they should not go into the same group.

// Given the integer n and the array dislikes where dislikes[i] = [ai, bi] indicates that the person labeled ai does not like the person labeled bi, return true if it is possible to split everyone into two groups in this way.

// Example 1:

// Input: n = 4, dislikes = [[1,2],[1,3],[2,4]]
// Output: true
// Explanation: The first group has [1,4], and the second group has [2,3].

func PossibleBipartition(n int, dislikes [][]int) bool {
	adjList := make(map[int][]int)

	for _, v := range dislikes {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	return bfsPartition(n, adjList)

}

func bfsPartition(n int, adjList map[int][]int) bool {
	queue := make([]Node, 0)

	visited := make(map[int]bool)
	colorList := make([]int, n+1)
	for k, _ := range adjList {
		if !visited[k] {
			queue = append(queue, Node{Color: 1, Val: k})
			visited[k] = true
			//p := k
			for len(queue) > 0 {
				temp := queue[0]
				queue = queue[1:]

				color := temp.Color

				for _, v := range adjList[temp.Val] {
					if !visited[v] {
						visited[v] = true
						queue = append(queue, Node{Val: v, Color: 3 - color})
						colorList[v] = 3 - color
					} else if color == colorList[v] {
						return false
					}
				}
			}
		}
	}

	return true
}

type Node struct {
	Color int
	Val   int
}
