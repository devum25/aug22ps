package binarysearch

import "math"

func NthMagicalNumber(n int, a int, b int) int {
	l := a
	if b < a {
		l = b
	}

	r := l * n

	for l <= r {

		mid := (l + r) / 2
		rank := ((mid / a) + (mid / b) - (mid / lcm(a, b)))
		if rank == n {
			if mid%a == 0 || mid%b == 0 {
				return int(math.Mod(float64(mid), math.Pow(10, 9)+7))
			} else {
				r = mid - 1
			}
		} else if rank > n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func lcm(a, b int) int {
	greater := b
	smaller := a

	if a > b {
		greater = a
		smaller = b
	}

	for i := greater; ; i += greater {
		if i%smaller == 0 {
			return i
		}
	}
}
