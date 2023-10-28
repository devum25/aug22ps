package subarrays

import "math"

func rotate(nums []int, k int) {

	k = int(math.Mod(float64(k), float64(len(nums))))

	reverse(nums, 0, len(nums)-1)

	reverse(nums, 0, k-1)

	reverse(nums, k, len(nums)-1)
}

func reverse(arr []int, s, e int) {

	i := s
	j := e

	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

}
