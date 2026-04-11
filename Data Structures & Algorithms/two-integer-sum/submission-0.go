func twoSum(nums []int, target int) []int {
	// i + j = target
	// i = target - j
	diffMap := map[int]int{}

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if j, ok := diffMap[nums[i]]; ok {
			return []int{j, i}
		}
		diffMap[diff] = i
	}
	return []int{0, 0}
}

