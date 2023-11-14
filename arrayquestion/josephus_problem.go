package arrayquestion

import "math"

// Observation : If N is a power of 2 then position 1 is the winner.
// For N, find the closest power of 2
// kill N-(closest power of 2 soldiers)
// Find position of sword after killing N-(closest power of 2 soldiers)
// that will be the answer

func JosephusProblem(n int) int {
	ans := 0

	x := findClosePowerOf2(n)
	ans = 2*(n-x) + 1 // N-X is number of persons to be killed,
	//since pointer is moving by 2 ==> 2*kills +1 is the index where sword will be when exactly 2powN persons are remaining
	return ans
}

func findClosePowerOf2(n int) int {
	i := 0
	for {
		if int(math.Pow(float64(2), float64(i))) > n {
			return int(math.Pow(float64(2), float64(i-1)))
		}
		i++
	}
}
