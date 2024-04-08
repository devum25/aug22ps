package arrayquestion

import "math"

func MinimumDistance(points [][]int) int {
	n := len(points)

	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		maxDist := math.MinInt
		for j := i + 1; j < n; j++ {
			// Calculate Manhattan distance between points[i] and points[j]
			//dist := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])

			// Remove one point and recalculate distances
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				maxDist = int(math.Max(float64(maxDist), float64(abs(points[k][0]-points[j][0])+abs(points[k][1]-points[j][1]))))
			}
		}
		arr = append(arr, maxDist)
	}
	min := math.MaxInt
	for i := 0; i < len(arr); i++ {
		if arr[i] < min && arr[i] != math.MinInt {
			min = arr[i]
		}
	}

	return min
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
