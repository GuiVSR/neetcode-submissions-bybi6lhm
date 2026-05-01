func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	lwr := strings.ToLower(s)
	for i = 0; i < len(lwr)/2; i++ {

		if i == j {
			return true
		}

		for !unicode.IsLetter(rune(lwr[i])) && !unicode.IsNumber(rune(lwr[i])) {
			i++
			if i == j {
				return true
			}
		}
		for !unicode.IsLetter(rune(lwr[j])) && !unicode.IsNumber(rune(lwr[j])) {
			j--
			if i == j {
				return true
			}
		}

		if i > j {
			return false
		}

		if lwr[i] != lwr[j] {
			return false
		}

		j--
	}
	return true
}