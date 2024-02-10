package graphs

// 1615. Maximal Network Rank
// There is an infrastructure of n cities with some number of roads connecting these cities. Each roads[i] = [ai, bi] indicates that there is a bidirectional road between cities ai and bi.
// The network rank of two different cities is defined as the total number of directly connected roads to either city. If a road is directly connected to both cities, it is only counted once.
// The maximal network rank of the infrastructure is the maximum network rank of all pairs of different cities.
// Given the integer n and the array roads, return the maximal network rank of the entire infrastructure.

func maximalNetworkRank(n int, roads [][]int) int {
	connected := make([]int, n)
	directConnections := make(map[int]map[int]bool)

	for _, road := range roads {
		a, b := road[0], road[1]
		connected[a]++
		connected[b]++

		if directConnections[a] == nil {
			directConnections[a] = make(map[int]bool)
		}
		directConnections[a][b] = true

		if directConnections[b] == nil {
			directConnections[b] = make(map[int]bool)
		}
		directConnections[b][a] = true
	}

	maxRank := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := connected[i] + connected[j]

			if directConnections[i][j] {
				rank--
			}

			maxRank = max(maxRank, rank)
		}
	}

	return maxRank
}
