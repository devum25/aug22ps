package heaps

func Frequency(s string) string {

	var res string
	heap := NewCharHeap()
	for i := 0; i < len(s); i++ {
		heap.Insert(s[i])
	}

	for len(heap.Items) > 0 {
		x := heap.DeleteMax()
		for i := 0; i < x.Freq; i++ {
			res += string(x.Char)
		}

	}

	return res
}

type HNode struct {
	Char int
	Freq int
}

type Heap struct {
	Items []*HNode
	Hash  map[int]int
	Size  int
}

func NewHeap() Heap {
	return Heap{Items: make([]*HNode, 0), Hash: make(map[int]int)}
}

// func (m *CharHeap) Insert(val byte) {
// 	var i int
// 	if len(m.Items) == 0 {
// 		i = m.SizeInt
// 	} else {
// 		i = m.Size + 1
// 		m.Size = i
// 	}

// 	if f, ok := m.Hash[val]; !ok {
// 		m.Hash[val] = m.Size
// 		m.Items = append(m.Items, &HeapNode{Char: val, Freq: 1})
// 	} else {
// 		node := m.Items[f]
// 		node.Freq++
// 		m.Size--
// 		i = f
// 	}

// 	for i > 0 {

// 		parent := (i - 1) / 2
// 		if m.Items[parent].Freq < m.Items[i].Freq {
// 			p := m.Items[parent].Char
// 			c := m.Items[i].Char
// 			m.Hash[p] = i
// 			m.Hash[c] = parent
// 			m.Items[parent], m.Items[i] = m.Items[i], m.Items[parent]
// 			i = parent
// 		} else {
// 			break
// 		}
// 	}
// }

// func (m *CharHeap) DeleteMax() *HeapNode {
// 	val := m.Items[0]
// 	c := m.Items[0].Char
// 	p := m.Items[m.Size].Char
// 	m.Hash[c] = m.Size
// 	m.Hash[p] = 0
// 	m.Items[0], m.Items[m.Size] = m.Items[m.Size], m.Items[0]

// 	m.Items = m.Items[:m.Size]
// 	delete(m.Hash, c)
// 	m.Size--

// 	i := 0

// 	for i < m.Size {
// 		max := i

// 		l := 2*i + 1
// 		r := 2*i + 2

// 		if l < len(m.Items) && m.Items[l].Freq > m.Items[max].Freq {
// 			max = l
// 		}

// 		if r < len(m.Items) && m.Items[r].Freq > m.Items[max].Freq {
// 			max = r
// 		}

// 		if i == max {
// 			break
// 		}

// 		x := m.Items[i].Char
// 		y := m.Items[max].Char
// 		m.Hash[x] = max
// 		m.Hash[y] = i
// 		m.Items[max], m.Items[i] = m.Items[i], m.Items[max]

// 		i = max
// 	}

// 	return val
// }
