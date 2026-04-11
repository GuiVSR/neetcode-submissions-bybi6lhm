func hasDuplicate(nums []int) bool {
    table := map[int]int{}

    for _, v := range nums {
        table[v] += 1
        if table[v] > 1 {
            return true
        }
    }
    return false
}
