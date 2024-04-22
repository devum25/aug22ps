package graphs

import "strconv"

func OpenLock(deadends []string, target string) int {
	visited := make(map[string]bool)

	for i := 0; i < len(deadends); i++ {
		visited[deadends[i]] = true
	}

	queue := make([]Node4, 0)
	queue = append(queue, Node4{Val: "0000"})

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		childs := getChildren(top.Val)
		dist := top.Dist

		for i := 0; i < len(childs); i++ {
			if childs[i] == target {
				return dist + 1
			}
			if !visited[childs[i]] {
				queue = append(queue, Node4{Val: childs[i], Dist: dist + 1})
			}
		}
	}

	return -1
}

func getChildren(str string) []string {
	arr := make([]string, 0)

	for i := 0; i < 4; i++ {
		val, _ := strconv.Atoi(string(str[i]))
		val = (val + 1) % 10
		s := str[:i] + strconv.Itoa(val) + str[i+1:]
		arr = append(arr, s)
	}

	for i := 0; i < 4; i++ {
		val, _ := strconv.Atoi(string(str[i]))
		val = (val - 1 + 10) % 10
		s := str[:i] + strconv.Itoa(val) + str[i+1:]
		arr = append(arr, s)
	}

	return arr
}

type Node4 struct {
	Val  string
	Dist int
}
