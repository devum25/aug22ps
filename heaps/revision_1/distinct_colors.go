package revision1

func QueryResults(limit int, queries [][]int) []int {
	hashColors := make(map[int]map[int]bool)
	q := make(map[int]int)
	for i := 0; i <= limit; i++ {
		q[i] = -1
	}
	ans := make([]int, 0)
	c := 0
	for i := 0; i < len(queries); i++ {
		x := queries[i][0]
		y := queries[i][1]

		if q[x] == -1 {
			q[x] = y
			if _, ok := hashColors[y]; !ok {
				temp := make(map[int]bool)
				temp[x] = true
				hashColors[y] = temp
				c++
			} else {
				temp := hashColors[y]
				temp[x] = true
				hashColors[y] = temp
			}
		} else {
			t := q[x]
			if _, ok := hashColors[t]; ok && len(hashColors[t]) == 1 {
				delete(hashColors, t)
				c--
			} else if _, ok := hashColors[t]; ok {
				temp := hashColors[t]
				delete(temp, x)
				hashColors[t] = temp
			}
			q[x] = y
			if _, ok := hashColors[y]; !ok {
				temp := make(map[int]bool)
				temp[x] = true
				hashColors[y] = temp
				c++
			} else {
				temp := hashColors[y]
				temp[x] = true
				hashColors[y] = temp
			}
		}

		ans = append(ans, c)
	}

	return ans
}

// ////// Optimal Solution
func queryResults(limit int, queries [][]int) []int {
	// Map to store the current color of each ball
	ballColors := make(map[int]int)
	// Map to store the count of each color
	colorCount := make(map[int]int)
	// Variable to store the number of distinct colors
	distinctColorCount := 0
	// Slice to store the result after each query
	result := make([]int, 0, len(queries))

	for _, query := range queries {
		x, y := query[0], query[1]

		if currentColor, exists := ballColors[x]; exists {
			// Ball x already has a color
			colorCount[currentColor]--
			if colorCount[currentColor] == 0 {
				distinctColorCount--
			}
		}

		// Update the color of ball x
		ballColors[x] = y
		colorCount[y]++
		if colorCount[y] == 1 {
			distinctColorCount++
		}

		// Append the current number of distinct colors to the result
		result = append(result, distinctColorCount)
	}

	return result
}
