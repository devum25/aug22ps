package stacks

// This question reminds me of 316-Remove Duplicate Letters, so I am going with a similar approach here.

// I maintain a stack containing the built prefix of the output. Starting from left to right of the string s, for any character that can improve the lexicographical order, we will pop a character from the stack only if:
// 1. There are enough characters left to construct a k-size string.
// 2. If the character to be popped is the same as letter, then there should also be at least r character letter left.

// When adding new element to stack, we just need to check if the stack is already full (has size of k), and if the character is not the same as letter, then also check if we are leaving enough space for the matching characters
// aaabbbcccddd,3,b,2
func SmallestSubsequenceWithLetter(s string, k int, letter byte, repetition int) string { // leetcode ,k = 4,rep=2,letter=e
	// lastIdxs := getLastIndexArray(s)
	// seen := make([]int, 26)
	mapIdx := getMapIdx(s, letter)
	stack := NewMyStack()
	// stack.Push(s[0])
	// seen[s[0]-'a'] = 1
	pushed := false
	isValid := false
	for i := 0; i < len(s); i++ {
		pushed = false
		if stack.Len() > 0 && s[i] == letter && stack.Peek() == letter {
			isValid, mapIdx = checkValidity(mapIdx, i, repetition, letter)
			if isValid {
				stack.Pop()
				repetition++
			}
		}

		for stack.Len() > 0 && s[i] < stack.Peek() && (stack.Len()+len(s)-1-i) >= k {
			if stack.Peek() == letter {
				isValid, mapIdx = checkValidity(mapIdx, i, repetition, letter)
				if isValid {
					stack.Pop()
					repetition++
					continue
				} else if stack.Len() == 1 {
					break
				}
			} else {
				stack.Pop()
			}
		}
		if !pushed && (stack.Len() == 0 || stack.Len() < k) {
			if s[i] == letter {
				repetition--
				stack.Push(s[i])
			} else if k-stack.Len() > repetition {
				stack.Push(s[i])
			}
		}

	}

	res := make([]byte, stack.Len())

	for i := stack.Len(); i > 0; i-- {
		res[i-1] = stack.Pop()
	}

	return string(res)
}

// remove if there are letter available at later index with count = rep
func checkValidity(mapIdx map[byte][]int, curr int, rep int, letter byte) (bool, map[byte][]int) {
	x := mapIdx[letter]
	res := []int{}
	count := 0
	for i := 0; i < len(x); i++ {
		if x[i] >= curr {
			res = append(res, x[i])
		} else {
			count++
		}

	}

	if len(res) >= curr && count > 0 {
		mapIdx[letter] = res
		return true, mapIdx
	}

	return false, mapIdx
}

func getMapIdx(s string, letter byte) map[byte][]int {
	mapIdx := make(map[byte][]int)
	x := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			x = append(x, i)
		}
	}

	mapIdx[letter] = x
	return mapIdx
}
