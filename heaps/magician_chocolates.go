package heaps

import "math"

// Given N bags, each bag contains Bi chocolates. There is a kid and a magician.
// In a unit of time, the kid can choose any bag i, and eat Bi chocolates from it, then the magician will fill the ith bag with floor(Bi/2) chocolates.
// Find the maximum number of chocolates that the kid can eat in A units of time.
// NOTE:
// floor() function returns the largest integer less than or equal to a given number.
// Return your answer modulo 109+7

// Input Format
// The first argument is an integer A.
// The second argument is an integer array B of size N.

// Output Format
// Return an integer denoting the maximum number of chocolates the kid can eat in A units of time.
// Example Input
// Input 1:
//  A = 3
//  B = [6, 5]
// Input 2:
//  A = 5
//  b = [2, 4, 6, 8, 10]
// Example Output
// Output 1:
//  14
// Output 2:
//  33
// Example Explanation
// Explanation 1:

//  At t = 1 kid eats 6 chocolates from bag 0, and the bag gets filled by 3 chocolates.
//  At t = 2 kid eats 5 chocolates from bag 1, and the bag gets filled by 2 chocolates.
//  At t = 3 kid eats 3 chocolates from bag 0, and the bag gets filled by 1 chocolate.
//  so, total number of chocolates eaten are 6 + 5 + 3 = 14
// Explanation 2:
// Maximum number of chocolates that can be eaten is 33.

func Nchoc(A int, B []int) int {
	heap := NewChocolateHeap()
	heap.Items = toChocolateHeap(B)
	heap.Size = len(heap.Items) - 1
	ans := 0
	i := A
	for i > 0 {
		c := heap.ExtractMax()
		ans += c
		heap.Insert(c / 2)
		i--
	}

	return (ans % int(math.Pow(10, 9)+7))
}

type ChocolateHeap struct {
	Items []int
	Size  int
}

func NewChocolateHeap() ChocolateHeap {
	return ChocolateHeap{Items: make([]int, 0)}
}

func (c *ChocolateHeap) Insert(val int) {
	if c.Size == 0 {
		c.Items = append(c.Items, val)
	} else {
		c.Size++
		c.Items = append(c.Items, val)
	}

	i := c.Size

	for i > 0 {
		parent := (i - 1) / 2

		if c.Items[parent] < c.Items[i] {
			c.Items[parent], c.Items[i] = c.Items[i], c.Items[parent]
			i = parent
		} else {
			break
		}
	}
}

func (c *ChocolateHeap) ExtractMax() int {
	val := c.Items[0]
	c.Items[0], c.Items[len(c.Items)-1] = c.Items[len(c.Items)-1], c.Items[0]
	c.Items = c.Items[:len(c.Items)-1]
	c.Size--

	i := 0
	for i < len(c.Items) {
		max := i

		l := 2*i + 1
		r := 2*i + 2

		if l < len(c.Items) && c.Items[l] > c.Items[max] {
			max = l
		}
		if r < len(c.Items) && c.Items[r] > c.Items[max] {
			max = r
		}

		if max == i {
			break
		}

		c.Items[max], c.Items[i] = c.Items[i], c.Items[max]
		i = max
	}

	return val
}

func toChocolateHeap(arr []int) []int {
	leaf := (len(arr) + 1) / 2
	res := make([]int, len(arr))
	for i := leaf; i < len(arr); i++ {
		res[i] = arr[i]
	}
	s := leaf - 1
	for s >= 0 {
		i := s
		for i < len(arr) {
			max := i
			l := 2*i + 1
			r := 2*i + 2
			for l < len(arr) && res[l] > res[max] {
				max = l
			}
			for r < len(arr) && res[r] > res[max] {
				max = r
			}
			if max == i {
				break
			}
			res[max], res[i] = res[i], res[max]
			i = max
		}
		s--
	}
	return res
}
