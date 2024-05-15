package revision1


type MinHeap struct{
	Items []int
	Size int
}

func NewMinHeap() *MinHeap{
	return &MinHeap{Items: make([]int, 0)}
}

func(m *MinHeap) InsertHeap(x int){
	var i  int
	if len(m.Items) == 0{
		i = m.Size
	}else{
	   m.Size++
	   i = m.Size
	}
	m.Items = append(m.Items, x)
	p := (i-1)/2

	for i > 0{
		if m.Items[p] > m.Items[i]{
			m.Items[p],m.Items[i] = m.Items[i],m.Items[p]
			i = p
		}else{
			break
		}
	}
} 

func (m *MinHeap) DeleteMin() int{
	
}