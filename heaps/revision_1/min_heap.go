package revision1

type MinHeap struct {
	Items []int
	Size  int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{Items: make([]int, 0)}
}

func (m *MinHeap) InsertHeap(x int) {
	var i int
	if len(m.Items) == 0 {
		i = m.Size
	} else {
		m.Size++
		i = m.Size
	}
	m.Items = append(m.Items, x)
	p := (i - 1) / 2

	for i > 0 {
		if m.Items[p] > m.Items[i] {
			m.Items[p], m.Items[i] = m.Items[i], m.Items[p]
			i = p
		} else {
			break
		}
	}
}

func (m *MinHeap) DeleteMin() int {
	val := m.Items[0]

	m.Items[0], m.Items[len(m.Items)-1] = m.Items[len(m.Items)-1], m.Items[0]
	m.Size--
	i := 0

	for i < len(m.Items) {
		min := i

		l := 2*i + 1
		r := 2*i + 2
		if l < len(m.Items) && m.Items[l] < m.Items[min] {
			min = l
		}
		if r < len(m.Items) && m.Items[r] < m.Items[min] {
			min = r
		}

		if min == i {
			break
		}

		m.Items[i], m.Items[min] = m.Items[min], m.Items[i]
		i = min
	}

	return val
}
