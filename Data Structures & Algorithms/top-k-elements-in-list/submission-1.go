func topK(freq map[int]int, k int) []int {
    res := []int{}

    for i := 0; i < k; i++ {
        maxKey := 0
        maxVal := -1

        for key, val := range freq {
            if val > maxVal {
                maxVal = val
                maxKey = key
            }
        }
        res = append(res, maxKey)
        delete(freq, maxKey)
    }
    return res
}


func topKFrequent(nums []int, k int) []int {
    numMap := map[int]int{}

    for _, v := range nums {
        numMap[v]++
    }



    return topK(numMap, k)
}
