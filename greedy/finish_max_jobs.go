package greedy

import "sort"

// Problem Description
// There are N jobs to be done, but you can do only one job at a time.
// Given an array A denoting the start time of the jobs and an array B denoting the finish time of the jobs.
// Your aim is to select jobs in such a way so that you can finish the maximum number of jobs.
// Return the maximum number of jobs you can finish.

// Input Format
// The first argument is an integer array A of size N, denoting the start time of the jobs.
// The second argument is an integer array B of size N, denoting the finish time of the jobs.

// Output Format
// Return an integer denoting the maximum number of jobs you can finish.

// Example Input
// Input 1:

//  A = [1, 5, 7, 1]
//  B = [7, 8, 8, 8]
// Input 2:

//  A = [3, 2, 6]
//  B = [9, 8, 9]
// Example Output
// Output 1:

//  2
// Output 2:

//  1
func MaxNumberOfJobs(A []int, B []int) int {
	res := make([][]int, 0)

	for i := 0; i < len(A); i++ {
		res = append(res, []int{A[i], B[i]})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i][1] < res[j][1]
	})

	act := 1
	end := res[0][1]

	for i := 1; i < len(res); i++ {
		if res[i][0] < end {
			continue
		} else {
			act++
			end = res[i][1]
		}
	}

	return act
}
