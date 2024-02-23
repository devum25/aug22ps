package arrayquestion

func Merge(nums1 []int, m int, nums2 []int, n int) {
	//arr := make([]int,len(nums1)+len(nums2))

	i := 0
	j := 0
	// c := 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			i++
		} else {
			for k := (m - 1); k >= i; k-- {
				nums1[k+1] = nums1[k]
			}
			nums1[i] = nums2[j]
			j++
			i++
			m++
		}
	}

	if i < m {
		for i < m {
			// nums1[i] = nums1[i]
			i++
		}
	}

	if j < n {
		for j < n {
			nums1[i] = nums2[j]
			j++
			i++
		}
	}

}
