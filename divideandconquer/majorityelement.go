package divideandconquer

func MajorityElement(nums []int) int {
	majority := nums[0]
	vote := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == majority {
			vote++
		} else if nums[i] != majority {
			vote--
		}

		if vote == 0 {
			majority = nums[i]
			vote = 1
		}
	}

	return majority
}

///////////////////////////////////////////////////////////////

func occurence(a []int, x int) int {
	if len(a) == 0 {
		return 0
	}
	res := 0
	if a[0] == x {
		res = 1 + occurence(a[1:], x)
	} else {
		res = occurence(a[1:], x)
	}

	return res
}

func majority_element(arr []int) (bool, int, int) {
	n := len(arr)
	if n == 0 {
		return false, 0, 0
	} else if n == 1 {
		return true, 1, arr[0]
	} else {
		b := arr[0 : n/2]
		c := arr[n/2:]

		exist, times, element := majority_element(b)
		if exist {
			occurences := occurence(c, element)
			if times+occurences > n/2 {
				return true, times + occurences, element
			}
		}

		e, t, ele := majority_element(c)
		if e {
			occurences := occurence(b, ele)
			if occurences+t > n/2 {
				return true, t + occurences, ele
			}
		}
	}

	return false, 0, 0
}
