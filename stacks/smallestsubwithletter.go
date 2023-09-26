package stacks

// This question reminds me of 316-Remove Duplicate Letters, so I am going with a similar approach here.

// I maintain a stack containing the built prefix of the output. Starting from left to right of the string s, for any character that can improve the lexicographical order, we will pop a character from the stack only if:
// 1. There are enough characters left to construct a k-size string.
// 2. If the character to be popped is the same as letter, then there should also be at least r character letter left.

// When adding new element to stack, we just need to check if the stack is already full (has size of k), and if the character is not the same as letter, then also check if we are leaving enough space for the matching characters
// aaabbbcccddd,3,b,2
func SmallestSubsequenceWithLetter(s string, k int, letter byte, r int) string { // leetcode ,k = 4,rep=2,letter=e
	n_letters := 0

	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			n_letters++
		}
	}

	stack := NewMyStack()

	for i := 0; i < len(s); i++ {
		for stack.Len() > 0 && stack.Peek() > s[i] && (len(s)-i+stack.Len() > k) &&
			(stack.Peek() != letter || n_letters > r) {
			if stack.Pop() == letter {
				r++
			}
		}

		if stack.Len() < k {
			if s[i] == letter {
				stack.Push(s[i])
				r--
			} else if k-stack.Len() > r {
				stack.Push(s[i])
			}
		}

		if s[i] == letter {
			n_letters--
		}
	}

	res := make([]byte, stack.Len())

	for i := stack.Len(); i > 0; i-- {
		res[i-1] = stack.Pop()
	}

	return string(res)
}
