package queue

// Given a string A denoting a stream of lowercase alphabets, you have to make a new string B.
// B is formed such that we have to find the first non-repeating character each time a character is inserted to the stream and append it at the end to B.
// If no non-repeating character is found, append '#' at the end of B.

// abadbc ==> a,ab,aba,abad,abadb,abadbc
// a,a,b,b,d,d

func NonRepeatingStream(A string) string {

	mapchar := make(map[byte]int)
	q1 := NewQueue()

	res := make([]byte, 0)

	for i := 0; i < len(A); i++ {

		mapchar[A[i]]++

		v := mapchar[A[i]]
		if v == 1 {
			q1.Enqueue(int(A[i]))
		} else if q1.Top() == int(A[i]) {
			q1.Dequeue()

			for mapchar[byte(q1.Top())] > 1 {
				q1.Dequeue()
			}
		}

		if q1.Top() != -1 {
			res = append(res, byte(q1.Top()))
		} else {
			res = append(res, '#')
		}

	}

	return string(res)
}
