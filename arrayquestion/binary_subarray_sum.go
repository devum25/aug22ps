package arrayquestion

func Rotate(nums []int, k int) {
	if k%len(nums) == 0 {
		return
	}

	k = k % len(nums)

	// x := len(nums)-k
	reverse(nums)

	reverse(nums[0:k])

	reverse(nums[k:])

	return
}

func reverse(arr []int) {
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
