package slidingwindow

func MinimumSwapsRequired(arr []int, B int) int {
	k := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] <= B {
			k++
		}
	}
	ans := 0
	for i := 0; i < k; i++ {
		if arr[i] <= B {
			ans++
		}
	}
	s := 1
	e := k + 1
	for s < e {
		count := 0
		if arr[s] <= B {
			count = ans - 1
		}
		if arr[e] <= B {
			count = ans + 1
		}

		if count < ans {
			ans = count
		}
		s++
		e++
	}

	return ans
}
