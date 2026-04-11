func checkAnagram(a string, b string) bool {
    if len(a) != len(b) {
		return false
	}
    aMap := buildAnagramMap(a)
    bMap := buildAnagramMap(b)
    
    for aRune, aRuneCount := range aMap {
        bRuneCount, ok := bMap[aRune]
        if !ok {
            return false
        }
        if aRuneCount != bRuneCount {
            return false
        }
    }
    return true
}

func buildAnagramMap(word string) map[rune]int {
    newMap := map[rune]int{}

    for _, currRune := range word {
        newMap[currRune]++
    }

    return newMap
}

func groupAnagrams(strs []string) [][]string {
    processed := map[string]bool{}
    output := [][]string{}

    for i := 0; i < len(strs); i++ {
        _, ok := processed[strs[i]]
        if ok == true{
            continue
        }
        processed[strs[i]] = true

        group := []string{strs[i]}
        for j := i + 1; j < len(strs); j++ {
            if checkAnagram(strs[i], strs[j]) {
                group = append(group, strs[j])
                processed[strs[j]] = true
            }
        }
        output = append(output, group)
    }
    return output
}
