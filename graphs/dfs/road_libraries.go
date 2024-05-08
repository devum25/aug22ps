package dfs

func RoadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	// Write your code here
	visited := make(map[int32]bool)

	adjList := make(map[int32][]int32)

	for _, v := range cities {
		adjList[v[0]] = append(adjList[v[0]], v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}
	connected := 0
	for i := 1; i <= int(n); i++ {
		if !visited[int32(i)] {
			count := 0
			dfs7(adjList, visited, int32(i), &count)
			if count == 1 {
				connected += int(c_lib)
			} else if c_road > c_lib {
				connected += count * int(c_lib)
			} else {
				connected += (count-1)*int(c_road) + int(c_lib)
			}
		}
	}
	return int64(connected)
}

func dfs7(adjList map[int32][]int32, visited map[int32]bool, node int32, count *int) {
	visited[node] = true
	x := *count
	x++
	*count = x
	lst := adjList[node]

	for i := 0; i < len(lst); i++ {
		if !visited[lst[i]] {
			dfs7(adjList, visited, lst[i], count)
		}
	}
}
