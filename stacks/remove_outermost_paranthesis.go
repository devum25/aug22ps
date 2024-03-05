package stacks


func RemoveOuterParentheses(s string) string {
	var result string
    count := 0

    for _, char := range s {
        if char == '(' {
            if count > 0 {
                result += string(char)
            }
            count++
        } else if char == ')' {
            count--
            if count > 0 {
                result += string(char)
            }
        }
    }

    return result
  }