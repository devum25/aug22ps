package heaps

import "math"

// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k,
// return the k closest points to the origin (0, 0).
// The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).
// You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).

func KClosest(points [][]int, k int) [][]int {
	heap := NewPriorityQueue()

	hash := make(map[float64][]int)
	for i := 0; i < len(points); i++ {
		val := math.Sqrt(math.Abs((float64((points[i][0] * points[i][0]) + (points[i][1] * points[i][1])))))
		valid := false
		if p, ok := hash[val]; ok {
			if p[0] == points[i][0] && p[1] == points[i][0] {
				valid = true
			}
		} else {
			x := make([]int, 0)
			x = append(x, points[i][0])
			x = append(x, points[i][1])

			hash[val] = x
		}

		if !valid {
			heap.Insert(val, points[i][0], points[i][1])
		}
	}

	ans := make([][]int, k)
	i := 0
	for i < k {
		x := make([]int, 2)

		temp := heap.DeleteMin()

		x[0] = temp.Left
		x[1] = temp.Right

		ans[i] = x
		i++
	}

	return ans

}

type Node struct {
	Val   float64
	Left  int
	Right int
}

type PriorityQueue struct {
	Items []Node
	size  int
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{Items: make([]Node, 0)}
}

func (m *PriorityQueue) Insert(val float64, left, Right int) {
	var i int
	if len(m.Items) == 0 {
		i = m.size
	} else {
		i = m.size + 1
		m.size = i
	}

	m.Items = append(m.Items, Node{Val: val, Left: left, Right: Right})
	for i > 0 {

		parent := (i - 1) / 2
		if m.Items[parent].Val > m.Items[i].Val {
			m.Items[parent], m.Items[i] = m.Items[i], m.Items[parent]
			i = parent
		} else {
			break
		}
	}
}

func (m *PriorityQueue) DeleteMin() Node {
	node := m.Items[0]
	m.Items[0], m.Items[m.size] = m.Items[m.size], m.Items[0]

	m.Items = m.Items[:m.size]
	m.size--

	i := 0

	for i < m.size {
		min := i

		l := 2*i + 1
		r := 2*i + 2

		if l < len(m.Items) && m.Items[l].Val < m.Items[min].Val {
			min = l
		}

		if r < len(m.Items) && m.Items[r].Val < m.Items[min].Val {
			min = r
		}

		if i == min {
			break
		}

		m.Items[min], m.Items[i] = m.Items[i], m.Items[min]

		i = min
	}

	return node
}
