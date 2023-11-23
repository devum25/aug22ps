package heaps

type MinHeap struct {
	Items []int
	size  int
}

func NewMinHeap() MinHeap {
	return MinHeap{Items: make([]int, 0)}
}

func (m *MinHeap) Insert(val int) {
	var i int
	if len(m.Items) == 0 {
		i = m.size
	} else {
		i = m.size + 1
		m.size = i
	}
	m.Items = append(m.Items, val)
	for i > 0 {

		parent := (i - 1) / 2
		if m.Items[parent] > m.Items[i] {
			m.Items[parent], m.Items[i] = m.Items[i], m.Items[parent]
			i = parent
		} else {
			break
		}
	}
}

func (m *MinHeap) DeleteMin() int {
	val := m.Items[0]
	m.Items[0], m.Items[m.size] = m.Items[m.size], m.Items[0]

	m.Items = m.Items[:m.size]
	m.size--

	i := 0

	for i < m.size {
		min := i

		l := 2*i + 1
		r := 2*i + 2

		if l < len(m.Items) && m.Items[l] < m.Items[min] {
			min = l
		}

		if r < len(m.Items) && m.Items[r] < m.Items[min] {
			min = r
		}

		if i == min {
			break
		}

		m.Items[min], m.Items[i] = m.Items[i], m.Items[min]

		i = min
	}

	return val
}

func InsertOptimalMinHeap(arr []int) []int { // O(N)
	res := make([]int, len(arr))

	n := len(arr)

	leaf := (n + 1) / 2

	for i := leaf; i < len(arr); i++ {
		res[i] = arr[i]
	}

	s := leaf - 1
	c := 2
	for s >= 0 {
		i := s
		l := c*i + 1
		r := c*i + 2
		min := i
		res[i] = arr[i]

		if l < len(arr) && res[l] < res[i] {
			min = l
		}

		if r < len(arr) && res[r] < res[i] {
			min = r
		}

		if i == min {
			s--
			continue
		}

		res[min], res[i] = res[i], res[min]

		s--
	}

	return res
}
