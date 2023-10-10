package queue

// Given an integer A, you have to find the Ath Perfect Number.

// A Perfect Number has the following properties:

// It comprises only 1 and 2.
// The number of digits in a Perfect number is even.
// It is a palindrome number.
// For example, 11, 22, 112211 are Perfect numbers, where 123, 121, 782, 1 are not.

var s [100005]string

func reverse(s string) string {
	var str string
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}
func PerfectNumber(A int) string {
	s[1] = "1"
	s[2] = "2"
	q := NewStringQueue()
	q.Enqueue(s[1])
	q.Enqueue(s[2])
	idx := 3
	for idx < A+3 {
		ss := q.Dequeue()
		s[idx] = ss + "1"
		idx++
		q.Enqueue(ss + "1")
		s[idx] = ss + "2"
		idx++
		q.Enqueue(ss + "2")

	}
	// s[A] has the first half of our final answer
	ans := s[A]
	anss := ans
	str := reverse(ans)
	return anss + str
}
