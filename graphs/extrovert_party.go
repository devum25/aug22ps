package graphs

// Only the people who have k friends can enter the party hall.
// each person should enter with his her k friends.

func ExtrovertParty(graph [][]int, k int, n int) int {

	adjList := make(map[int][]int)

	for _, v := range graph {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	return 0
}
