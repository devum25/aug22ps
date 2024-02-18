package arrayquestion

func InsertInterval(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	} else if len(newInterval) == 0 {
		return intervals
	}

	res := make([][]int, 0)

	x := make([]int, 0)
	x = append(x, intervals[0][0])
	x = append(x, intervals[0][1])
	merged := false
	if evaluate1(x[0], x[1], newInterval[0], newInterval[1]) {
		start := min(x[0], newInterval[0])
		end := max(x[1], newInterval[1])
		temp := make([]int, 0)
		temp = append(temp, start)
		temp = append(temp, end)
		res = append(res, temp)
		merged = true
	} else {
		res = append(res, x)
	}

	for i := 1; i < len(intervals); i++ {
		if merged {
			if evaluate1(res[len(res)-1][0], res[len(res)-1][1], intervals[i][0], intervals[i][1]) {
				start := min(res[len(res)-1][0], intervals[i][0])
				end := max(res[len(res)-1][1], intervals[i][1])
				res = res[:len(res)-1]
				temp := make([]int, 0)
				temp = append(temp, start)
				temp = append(temp, end)
				res = append(res, temp)
			} else {
				temp := make([]int, 0)
				temp = append(temp, intervals[i][0])
				temp = append(temp, intervals[i][1])
				res = append(res, temp)
			}
		} else {
			if evaluate1(intervals[i][0], intervals[i][1], newInterval[0], newInterval[1]) {
				start := min(intervals[i][0], newInterval[0])
				end := max(intervals[i][1], newInterval[1])
				temp := make([]int, 0)
				temp = append(temp, start)
				temp = append(temp, end)
				res = append(res, temp)
				merged = true
			} else {
				temp := make([]int, 0)
				temp = append(temp, intervals[i][0])
				temp = append(temp, intervals[i][1])
				res = append(res, temp)
			}
		}
	}

	if !merged {
		res = append(res, newInterval)
	}

	return res
}

func evaluate1(si, ei, sj, ej int) bool {
	if sj <= si && ej >= ei {
		return true
	} else if si <= sj && ej <= ei {
		return true
	} else if sj >= si && sj <= ei {
		return true
	} else if sj <= si && ej <= ei {
		return true
	}

	return false
}
