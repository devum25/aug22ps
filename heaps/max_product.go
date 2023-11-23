package heaps

func MaxProduct(nums []int) int {

	heap := NewMaxHeapP()

	heap.Items = toMaxHeap(nums)
	heap.Size = len(heap.Items) - 1

	first := heap.ExtractMax() - 1
	second := heap.ExtractMax() - 1

	return first * second

}

type MaxHeapP struct {
	Items []int
	Size  int
}

func NewMaxHeapP() *MaxHeapP {
	return &MaxHeapP{Items: make([]int, 0)}
}

func (m *MaxHeapP) Insert(val int) {
	if len(m.Items) == 0 {
		m.Size = 0
	} else {
		m.Size++
	}

	m.Items = append(m.Items, val)

	i := m.Size

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

func (m *MaxHeapP) ExtractMax() int {
	val := m.Items[0]
	m.Items[0], m.Items[len(m.Items)-1] = m.Items[len(m.Items)-1], m.Items[0]
	m.Items = m.Items[:m.Size]
	m.Size--

	i := 0

	for i < len(m.Items) {
		max := i

		l := 2*i + 1
		r := 2*i + 2

		for l < len(m.Items) && m.Items[l] > m.Items[max] {
			max = l
		}

		for r < len(m.Items) && m.Items[r] > m.Items[max] {
			max = r
		}

		if i == max {
			break
		}
		m.Items[max], m.Items[i] = m.Items[i], m.Items[max]
		i = max
	}

	return val
}

func toMaxHeap(arr []int) []int {

	leaf := (len(arr) + 1) / 2

	res := make([]int, len(arr))

	for i := leaf; i < len(arr); i++ {
		res[i] = arr[i]
	}

	s := leaf - 1

	for s >= 0 {
		res[s] = arr[s]
		i := s
		for i < len(arr) {
			max := i
			l := 2*i + 1
			r := 2*i + 2

			if l < len(arr) && res[l] > res[max] {
				max = l
			}
			if r < len(arr) && res[r] > res[max] {
				max = r
			}
			if i == max {
				break
			}
			res[max], res[i] = res[i], res[max]
			i = max
		}
		s--
	}

	return res
}
