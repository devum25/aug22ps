package arrayquestion

//Given an array of int, count the numbr of triplets i,j & k such that
//i<j<k
//A[i]<A[j]<A[k]
// [2,6,9,4,10] --> [2,6,9],[2,6,10],[2,9,10],[2,4,10],[6,9,10] --> 5

func CountTriplet(arr []int) int {
	ans := 0

	//for every A[i] check how many numbers are less that A[i] on left side
	// and how many numbers on right side is greater that A[i]
	for i := 0; i < len(arr); i++ {
		left := 0
		right := 0
		temp := arr[i]
		j := i - 1

		for j >= 0 {
			if arr[j] < temp {
				left++
			}
			j--
		}
		k := i + 1
		for k < len(arr) {
			if arr[k] > temp {
				right++
			}
			k++
		}

		ans += left * right
	}

	return ans
}
