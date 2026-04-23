func productExceptSelf(nums []int) []int {
	prefix := []int{}
	suffix := []int{}
	ans := []int{}
	prefixTotal := 1
	suffixTotal := 1

	for i := 0; i < len(nums); i++ {
		prefix = append(prefix, prefixTotal)
		prefixTotal = prefixTotal * nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		suffix = append(suffix, suffixTotal)
		suffixTotal = suffixTotal * nums[i]
	}

	for i := 0; i < len(nums); i++ {
		product := prefix[i] * suffix[len(nums)-1-i]
		ans = append(ans, product)
	}

	return ans
}