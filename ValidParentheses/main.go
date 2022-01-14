package main

func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}
	sl := make([]string, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			sl = append(sl, ")")
		case '[':
			sl = append(sl, "]")
		case '{':
			sl = append(sl, "}")
		case ')':
			if len(sl) == 0 || sl[len(sl)-1] != ")" {
				return false
			}
			sl = sl[:len(sl)-1]
		case ']':
			if len(sl) == 0 || sl[len(sl)-1] != "]" {
				return false
			}
			sl = sl[:len(sl)-1]
		case '}':
			if len(sl) == 0 || sl[len(sl)-1] != "}" {
				return false
			}
			sl = sl[:len(sl)-1]
		}
	}
	return len(sl) == 0
}

func isValid2(s string) bool {
	var stack []byte
	for _, c := range s {
		if c == '(' {
			stack = append(stack, ')')
		} else if c == '{' {
			stack = append(stack, '}')
		} else if c == '[' {
			stack = append(stack, ']')
		} else if len(stack) == 0 || stack[len(stack)-1] != byte(c) {
			return false
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}
func main() {
	strs := []string{"()", "()[]{}", "(]", "([)]", "{[]}", ""}
	for _, s := range strs {
		println(s, isValid2(s))
	}

}
