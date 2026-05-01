func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	m := (l + r) / 2

	for l <= r {
		if nums[m] == target {
			return m
		}

		if target > nums[m] {
			l = m + 1
		}

		if target < nums[m] {
			r = m - 1
		}

		m = (l + r) / 2
	}

	return -1
}