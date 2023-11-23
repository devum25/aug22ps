package heaps

import "sort"

type kv struct {
	Key   byte
	Value int
}

func frequencySort(s string) string {

	hash := make(map[byte]int)
	var res string
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}

	var ss []kv
	for k, v := range hash {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss {
		for i := 0; i < kv.Value; i++ {
			res += string(kv.Key)
		}
	}

	return res
}

// ============================================================================================================
func FrequencySortOptimal(s string) string {

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

type HeapNode struct {
	Char byte
	Freq int
}

type CharHeap struct {
	Items []*HeapNode
	Hash  map[byte]int
	Size  int
}

func NewCharHeap() CharHeap {
	return CharHeap{Items: make([]*HeapNode, 0), Hash: make(map[byte]int)}
}

func (m *CharHeap) Insert(val byte) {
	var i int
	if len(m.Items) == 0 {
		i = m.Size
	} else {
		i = m.Size + 1
		m.Size = i
	}

	if f, ok := m.Hash[val]; !ok {
		m.Hash[val] = m.Size
		m.Items = append(m.Items, &HeapNode{Char: val, Freq: 1})
	} else {
		node := m.Items[f]
		node.Freq++
		m.Size--
		i = f
	}

	for i > 0 {

		parent := (i - 1) / 2
		if m.Items[parent].Freq < m.Items[i].Freq {
			p := m.Items[parent].Char
			c := m.Items[i].Char
			m.Hash[p] = i
			m.Hash[c] = parent
			m.Items[parent], m.Items[i] = m.Items[i], m.Items[parent]
			i = parent
		} else {
			break
		}
	}
}

func (m *CharHeap) DeleteMax() *HeapNode {
	val := m.Items[0]
	c := m.Items[0].Char
	p := m.Items[m.Size].Char
	m.Hash[c] = m.Size
	m.Hash[p] = 0
	m.Items[0], m.Items[m.Size] = m.Items[m.Size], m.Items[0]

	m.Items = m.Items[:m.Size]
	delete(m.Hash, c)
	m.Size--

	i := 0

	for i < m.Size {
		max := i

		l := 2*i + 1
		r := 2*i + 2

		if l < len(m.Items) && m.Items[l].Freq > m.Items[max].Freq {
			max = l
		}

		if r < len(m.Items) && m.Items[r].Freq > m.Items[max].Freq {
			max = r
		}

		if i == max {
			break
		}

		x := m.Items[i].Char
		y := m.Items[max].Char
		m.Hash[x] = max
		m.Hash[y] = i
		m.Items[max], m.Items[i] = m.Items[i], m.Items[max]

		i = max
	}

	return val
}

// ======================================================================================================

func frequencySortSuperOptimal(s string) string {
	freq := make([]int, 255)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		if freq[b[i]] == freq[b[j]] {
			return b[i] > b[j]
		}
		return freq[b[i]] > freq[b[j]]
	})
	return string(b)
}
