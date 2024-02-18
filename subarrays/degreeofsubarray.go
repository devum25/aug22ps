package subarrays

func FindShortestSubArray(nums []int) int {
	hash := make(map[int][3]int)
	max := 0
	idx := -1
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; !ok {
			hash[nums[i]] = [3]int{1, i, i}
			temp := 1
			if temp > max {
				max = temp
				idx = i
			}
		} else {
			arr := hash[nums[i]]
			temp := arr[0] + 1

			arr[0] = temp
			arr[2] = i
			hash[nums[i]] = arr
			if temp > max {
				max = temp
				idx = i
			}
		}
	}

	i, j := hash[nums[idx]][1], hash[nums[idx]][2]

	return len(nums[i : j+1])
}
