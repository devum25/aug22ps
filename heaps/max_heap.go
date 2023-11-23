package heaps

type MaxHeap struct {
	Items []int
	size  int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{Items: make([]int, 0)}
}

func (m *MaxHeap) Insert(val int) {
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
		if m.Items[parent] < m.Items[i] {
			m.Items[parent], m.Items[i] = m.Items[i], m.Items[parent]
			i = parent
		} else {
			break
		}
	}
}

func (m *MaxHeap) DeleteMax() int {
	val := m.Items[0]
	m.Items[0], m.Items[m.size] = m.Items[m.size], m.Items[0]

	m.Items = m.Items[:m.size]
	m.size--

	i := 0

	for i < m.size {
		min := i

		l := 2*i + 1
		r := 2*i + 2

		if l < len(m.Items) && m.Items[l] > m.Items[min] {
			min = l
		}

		if r < len(m.Items) && m.Items[r] > m.Items[min] {
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
