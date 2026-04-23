type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for _, v := range strs {
		strlen := strconv.Itoa(len(v))
		encoded = encoded + strlen + "#" + v
	}
	fmt.Println(encoded)
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	decoded := []string{}
	i := 0

	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		wordLen, _ := strconv.Atoi(encoded[i:j])
		word := encoded[j+1 : j+1+wordLen]
		decoded = append(decoded, word)

		i = j + 1 + wordLen
	}

	return decoded
}
