func isPalindrome(s string) bool {
	lwrcase := strings.ToLower(s)
	runes := []rune{}

	for _, r := range lwrcase {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			runes = append(runes, r)
		}
	}

	i := 0
	j := len(runes) - 1

	for i <= j {
		temp := runes[i]
		temp2 := runes[j]
		if temp != temp2 {
			return false
		}
		if j == i+1 || j == i {
			return true
		} else {
			i++
			j--
		}
	}
	return true
}