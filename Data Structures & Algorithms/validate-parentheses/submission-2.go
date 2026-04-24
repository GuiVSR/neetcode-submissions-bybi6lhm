func isValid(s string) bool {

	if len(s) % 2 == 1 {
		return false
	}  

    openChars := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	closedChars := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []rune{}

	for _, v := range s {
	  _, isOpenChar := openChars[v]
	  if isOpenChar {
		stack = append(stack, v)
	  } 
	  expected, isClosedChar := closedChars[v]
	  if isClosedChar {
		if len(stack) == 0 {
			return false
		}
		if stack[len(stack) - 1] != expected {
			return false
		}
		stack = stack[:len(stack)  -1]
	  }

	}
	if len(stack) != 0 {
		return false
	}
	return true
}
