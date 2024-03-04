package stacks

func IsValid(s string) bool {
    stack := make([]string,0)
    if s[0] != '(' && s[0] != '[' && s[0] != '{'{
            return false
     }

    stack = append(stack,string(s[0]))

    for i:=1;i<len(s);i++{   
        if string(s[i]) == ")"{
            if stack[len(stack)-1] == "("{
                stack = stack[:len(stack)-1]
				continue
            }else{
                return false
            }
        }
         if string(s[i]) == "}"{
            if stack[len(stack)-1] == "{"{
                stack = stack[:len(stack)-1]
				continue
            }else{
                return false
            }
        }
        if string(s[i]) == "]"{
            if stack[len(stack)-1] == "["{
                stack = stack[:len(stack)-1]
				continue
            }else{
                return false
            }
        }
        stack = append(stack,string(s[i]))
    }

    return true
}