package divideandconquer

// 229. Majority Element II
// Medium
// 9.1K
// 388
// Companies
// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.
func MajorityElement2(nums []int) []int {
	hshmap := make(map[int]int)
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(hshmap) < 4 {
			hshmap[nums[i]]++ // add only 3 elements
		}
		rem := make([]int, 0)
		if len(hshmap) > 2 { // start reducing count as element reaches 3
			for k, v := range hshmap {
				hshmap[k] = v - 1
				if hshmap[k] == 0 {
					rem = append(rem, k)
				}
			}
		}

		for i := 0; i < len(rem); i++ { // remove elements whose count is 0
			delete(hshmap, rem[i])
		}
	}

	check := make([]int, 2)
	i := 0
	for k, _ := range hshmap {
		check[i] = k
		i++
	}
	c1 := 0
	c2 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == check[0] {
			c1++
		} else if nums[i] == check[1] {
			c2++
		}
	}

	if c1 > len(nums)/3 {
		result = append(result, check[0])
	}
	if c2 > len(nums)/3 {
		result = append(result, check[1])
	}

	return result
}
