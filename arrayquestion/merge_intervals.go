package arrayquestion

import (
	"sort"
)

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

// Example 1:

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6]

func MergeIntervals(intervals [][]int) [][]int {
	lstInterval := make([]Interval, 0)
	for i := 0; i < len(intervals); i++ {
		lstInterval = append(lstInterval, Interval{Start: intervals[i][0], End: intervals[i][1]})
	}

	res := make([][]int, 0)

	sort.Sort(ByStartTime(lstInterval))

	x := make([]int, 0)
	x = append(x, lstInterval[0].Start)
	x = append(x, lstInterval[0].End)

	res = append(res, x)
	for i := 1; i < len(lstInterval); i++ {
		if evaluate(res[len(res)-1][0], res[len(res)-1][1], lstInterval[i].Start, lstInterval[i].End) {
			//res = res[:(i-1)]
			start := min(res[len(res)-1][0], lstInterval[i].Start)
			end := max(res[len(res)-1][1], lstInterval[i].End)

			res = res[:(len(res) - 1)]
			temp := make([]int, 0)
			temp = append(temp, start)
			temp = append(temp, end)
			res = append(res, temp)
		} else {
			temp := make([]int, 0)
			temp = append(temp, lstInterval[i].Start)
			temp = append(temp, lstInterval[i].End)
			res = append(res, temp)
		}
	}

	return res
}

func evaluate(si, ei, sj, ej int) bool {
	if si == sj {
		return true
	} else if ei == ej {
		return true
	} else if sj <= ei {
		return true
	}

	return false
}

type Interval struct {
	Start int
	End   int
}

type ByStartTime []Interval

func (a ByStartTime) Len() int { return len(a) }
func (a ByStartTime) Less(i, j int) bool {
	return a[i].Start < a[j].Start
}

func (a ByStartTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
