package dfs

//https://www.hackerrank.com/challenges/even-tree/problem

func EvenForest(t_nodes int32, t_edges int32, t_from []int32, t_to []int32) int32 {
	adjList := make(map[int32][]int32)

	for i := 0; i < len(t_from); i++ {
		adjList[t_from[i]] = append(adjList[t_from[i]], t_to[i])
		adjList[t_to[i]] = append(adjList[t_to[i]], t_from[i])
	}
	subtree := make(map[int32]int32)
	dfs11(adjList, 1, -1, subtree)

	var ans int32

	for _, v := range subtree {
		if v%2 == 0 {
			ans++
		}
	}
	return ans - 1
}

func dfs11(adjList map[int32][]int32, node int32, parent int32, subtree map[int32]int32) int32 {

	var count int32
	count = 1
	for _, v := range adjList[node] {
		if v != parent {
			count += dfs11(adjList, v, node, subtree)
		}
	}

	subtree[node] = count
	return count
}
