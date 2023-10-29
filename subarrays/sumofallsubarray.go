package subarrays

func SumSubArray(arr []int) int {

	sum := 0

	for i := 0; i < len(arr); i++ {
		x := 0
		for j := i; j < len(arr); j++ {
			x = x + arr[j]
			sum = sum + x
		}
	}

	return sum

}

// Contribution Approach // O(n) O(1)
func SumOfSubArrayONApproach(arr []int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		countof_Ai_insubarray := (i + 1) * (len(arr) - i)
		sum = sum + countof_Ai_insubarray*arr[i]

	}

	return sum
}
