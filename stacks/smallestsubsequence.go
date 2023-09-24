package stacks

func SmallestSubsequence(s string, k int, letter byte, repetition int) string { // leetcode ,k = 4,rep=2,letter=e
	// lastIdxs := getLastIndexArray(s)
	seen := make([]int, 26)
	// mapIdx := getMapIdx(s, letter)
	stack := NewMyStack()
	stack.Push(s[0])
	seen[s[0]-'a'] = 1
	pushed := false
	// isValid := false
	for i := 1; i < len(s); i++ {
		pushed = false
		// if s[i] == letter {
		// 	if stack.Peek() == letter {
		// 		isValid, mapIdx = checkValidity(mapIdx, i, repetition)
		// 		if isValid {
		// 			stack.Pop()
		// 			stack.Push(s[i])
		// 			continue
		// 		}
		// 	} else {
		// 		stack.Push(s[i])
		// 		continue
		// 	}
		// }
		for stack.Len() > 0 && s[i] < stack.Peek() && (stack.Len()+len(s)-1-i) >= k {
			stack.Pop()
		}
		if !pushed {
			stack.Push(s[i])
		}

	}

	res := make([]byte, stack.Len())

	for i := stack.Len(); i > 0; i-- {
		res[i-1] = stack.Pop()
	}

	return string(res)
}
