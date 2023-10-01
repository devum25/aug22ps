package main

import (
	"fmt"

	"github.com/devum25/techbench/binarysearch"
)

// // [4,6,7,4,6,9,7,2,3,6] //9
// func main() {

// 	lru := linkedlist.NewLRUCache(2)
// 	// lru.Get(2)
// 	lru.Put(2, 6)
// 	// fmt.Print(lru.Get(1))
// 	lru.Put(1, 1)
// 	lru.Put(2, 3)
// 	lru.Put(4, 1)
// 	fmt.Print(lru.Get(1))
// 	fmt.Print(lru.Get(2))
// 	lru.Get(4)
// 	lru.Get(2)
// 	lru.Put(1, 1)
// 	// for i := 2; i <= 9; i++ {
// 	// 	linkedlist.InsertAtEnd(head, i) // 0-->1-->2-->3-->4-->5-->null   output 5-->4-->3-->2-->1-->0-->null
// 	// }
// 	//67 -> 27 -> 64 -> 10 -> 4 -> 85
// 	// fmt.Print(head)
// 	head1 := &linkedlist.Node{Val: 2}

// 	linkedlist.InsertAtEnd(head1, 1)
// 	linkedlist.InsertAtEnd(head1, 2)
// 	linkedlist.InsertAtEnd(head1, 1)
// 	linkedlist.InsertAtEnd(head1, 2)
// 	linkedlist.InsertAtEnd(head1, 2)
// 	linkedlist.InsertAtEnd(head1, 1)
// 	linkedlist.InsertAtEnd(head1, 3)
// 	linkedlist.InsertAtEnd(head1, 2)
// 	linkedlist.InsertAtEnd(head1, 2)
// 	// linkedlist.InsertAtEnd(head2, 4)
// 	// linkedlist.InsertAtEnd(head2, 85)
// 	// linkedlist.InsertAtEnd(head1, 33)
// 	// linkedlist.InsertAtEnd(head1, 1)
// 	// linkedlist.InsertAtEnd(head1, 7)
// 	// head = linkedlist.DeleteAtEnd(head)

// 	// head2 := &linkedlist.Node{Val: 2}

// 	// linkedlist.InsertAtEnd(head2, 6)
// 	// linkedlist.InsertAtEnd(head2, 7)
// 	// linkedlist.InsertAtEnd(head2, 11)
// 	// linkedlist.InsertAtEnd(head2, 12)
// 	// linkedlist.InsertAtEnd(head2, 15)

// 	fmt.Print(linkedlist.LongestPalindrome(head1))

// 	// head = linkedlist.FindMidNodeFirstMid(head)

// 	fmt.Print(head1)

// 	// arr := []int{12, 9, 2, 10, 23, 19, 33, 1, 7}
// 	// sorting.Mergesort(arr)
// 	// fmt.Print(arr)

// }
// 8,222,222,-1,222,1111111,-1,1
// 59,59,97,58,58,97,-1,-23,84,97
func main() {
	// s := stacks.ConvertToPostFix("3 + 10 * ( 3 - 4 / 2 ) + 3")
	// s := stacks.ConvertToPostFix("13+10*(3-4/2)+3")
	// x := []int{0, 2, 0}
	// fmt.Print(stacks.SmallestSubsequenceWithLetter("wuynymkihfdcbabefiiymnoyyytywzy", 16, 'y', 4))
	// fmt.Print(s)
	// matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	// fmt.Print(binarysearch.SearchMatrix(matrix, 3)) // 230412
	// matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	// matrix := [][]int{{1, 1}}
	fmt.Print(binarysearch.AggresiveCows([]int{3538, 1919, 533, 7521, 6, 7758, 5455, 9358, 4031, 6172, 7610, 5122, 6179, 2257, 7978, 5864, 9268, 2267, 7720, 1086, 3070, 5974, 8425, 1635, 4104, 9321, 2377, 8663, 7580, 107, 8741, 7233, 5639, 3719, 8697, 3875, 5831, 5538, 7641, 8384, 6501, 1612, 5324, 1540, 1171, 9406, 5482, 2057, 2350, 5988, 4933, 7996, 6266, 1922, 8387, 231, 3956, 5265, 9281, 8078, 869, 1596, 8842, 104, 6687, 2811, 5276, 6404, 9385, 9516, 2490, 2064, 2790, 9291, 3065, 5472, 4805, 5869, 3366, 9191, 4282, 2269, 9581, 5644, 8158, 4364, 2269, 4597, 6827, 9985, 3562, 4528, 5021, 2701, 5535, 5369, 7434, 8694, 494, 5816, 3248, 9774, 3898, 4432, 1344, 5518, 8470, 109, 8474, 2758, 3706, 4095, 5946, 1190, 8416, 2766, 6522, 9755, 5180, 7589, 2067, 9629, 4710, 6321, 4309, 6142, 4972, 7375, 5518, 3748, 4609, 7507, 3348, 5318, 2737, 1510, 4199, 9414, 7508, 1417, 2442, 6984, 3936, 1880, 8670, 1966, 7486, 7863, 8696, 8632, 3619, 547, 6489, 2223, 5123, 7435, 8412, 8408, 8394, 3051, 3206, 1370, 7056, 7345, 2008, 17, 6317, 9917, 5344, 6661, 4260, 9553, 9533, 2822, 6936, 166, 3222, 2335, 2531, 6245, 5881, 2808, 6198, 3141, 3978, 2676, 1133, 180, 7336, 2581, 3464, 7394, 7643, 6610, 8107, 9966, 3510, 1053, 350, 7963, 1598, 5770, 3196, 2773, 1299, 5989, 3618, 5815, 1306, 441, 1890, 9570, 9952, 5073, 6847, 5730, 7627, 289, 6537, 368, 8858, 3657, 8117, 1433, 3589, 9302, 3015, 9562, 2148, 9869, 3416, 9265, 5975, 6836, 3942, 3858, 6717, 214, 8018, 4151, 9557, 9218, 1895, 7888, 2761, 1382, 533, 4208, 4340, 8767, 8615, 4977, 8445, 6367, 6886, 2724, 8671, 9700, 8668, 3173, 2784, 5611, 3920, 8840, 3553, 2799, 8564, 9026, 3358, 6311, 4878, 698, 6026, 3275, 4775, 4179, 7947, 5564, 1467, 7369, 1960, 8724, 6658, 2315, 1695, 5499, 9746, 8316, 1121, 6617, 2735, 6483, 3908, 1084, 9559, 1714, 9845, 3206, 1337, 9344, 1077, 3751, 9741, 5402, 9186, 8284, 9055, 1395, 5239, 6510, 3821, 6748, 8888, 2700, 2996, 4690, 9265, 8725, 3450, 231, 24, 8054, 5322, 2098, 3193, 1195, 9282, 4025, 1666, 464, 1026, 4649, 8192, 4733, 1531, 9363, 9420, 2860, 2455, 2003, 7267, 7877, 4940, 8985, 6197, 2668, 6524, 3574, 4301, 1916, 5369, 3454, 8061, 4337, 3118, 351, 5549, 3897, 7757, 6626, 506}, 339))

	// arr := []int{5, 2, 4, 7, 1, 1}
	// for i := 0; i < len(arr); i++ {
	// 	s.Push(arr[i])
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Print(s.GetMin())
	// 	s.Pop()
	// }
}

//abefiimnoyytywzy
//uynyabefiiymnoyy
