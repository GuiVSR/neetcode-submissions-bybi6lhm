func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := map[rune]int{}

    for _, char := range s {
        count[char] += 1
    }

    for _, char := range t {
        count[char] -= 1
        if count[char] < 0 {
            return false
        }
    }

    return true
}
