package graphs

import "sort"

// There is a directed graph of n nodes with each node labeled from 0 to n - 1.
// The graph is represented by a 0-indexed 2D integer array graph where graph[i] is an integer array of nodes adjacent to node i,
//  meaning there is an edge from node i to each node in graph[i].

// A node is a terminal node if there are no outgoing edges.
//  A node is a safe node if every possible path starting from that node leads to a terminal node (or another safe node).

// Return an array containing all the safe nodes of the graph. The answer should be sorted in ascending order.

func EventualSafeNodes(graph [][]int) []int {
	incoming := make(map[int][]int)
	outgoing := make(map[int][]int)
	for i := 0; i < len(graph); i++ {
		incoming[i] = make([]int, 0)
		outgoing[i] = make([]int, 0)
	}

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			incoming[graph[i][j]] = append(incoming[graph[i][j]], i)
		}
	}

	for i, v := range graph {
		outgoing[i] = append(outgoing[i], v...)
	}

	return topological(incoming, outgoing, len(graph))
}

func topological(incoming, outgoing map[int][]int, n int) []int {
	outdegree := make([]int, n)
	queue := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		outdegree[i] = len(outgoing[i])
		if outdegree[i] == 0 {
			queue = append(queue, i)
			res = append(res, i)
			outdegree[i] = -1
		}
	}

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		for _, v := range incoming[temp] {
			outdegree[v]--
			if outdegree[v] == 0 {
				queue = append(queue, v)
				outdegree[v] = -1
				res = append(res, v)
			}
		}
	}

	sort.Ints(res)

	return res
}

// type Nodes struct{
// 	Incoming int
// 	Outgoing
// }
