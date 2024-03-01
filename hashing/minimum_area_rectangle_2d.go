package hashset

import "math"

// You are given an array of points in the X-Y plane points where points[i] = [xi, yi].

// Return the minimum area of a rectangle formed from these points, with sides parallel to the X and Y axes. If there is not any such rectangle, return 0.
func MinAreaRect(points [][]int) int {
	hash := make(map[string]bool)

	for i := 0; i < len(points); i++ {
		s := getString(points[i][0]) + ":" + getString(points[i][1])
		hash[s] = true
	}
	min := math.MaxInt
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if points[i][0] != points[j][0] && points[i][1] != points[j][1] {
				s1 := getString(points[i][0]) + ":" + getString(points[j][1])
				s2 := getString(points[j][0]) + ":" + getString(points[j][0])
				if hash[s1] && hash[s2] {
					l := int(math.Sqrt(float64(points[j][1]-points[i][1]) * float64(points[j][1]-points[i][1])))
					r := int(math.Sqrt(float64(points[j][0]-points[i][0]) * float64(points[j][0]-points[i][0])))
					min = minimum(min, l*r)
				}
			}
		}
	}

	return min
}

func minimum(a, b int) int {
	if a < b {
		return a
	}

	return b
}
