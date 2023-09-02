package strings

import "fmt"

// S := a,e,i,o,u,A,E,I,O,U
func AmazingSubArrays() {
	s := "ABEC"

	//create map
	mapOfVowels := make(map[byte]int)

	x := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for i := 0; i < len(x); i++ {
		mapOfVowels[x[i]]++
	}

	count := 0

	for i := 0; i < len(s); i++ {
		if mapOfVowels[s[i]] > 0 {
			count = count + (len(s) - i)
		}

	}

	fmt.Print(count)
}
