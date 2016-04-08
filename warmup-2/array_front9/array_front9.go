package array_front9

func array_front9(nums []int) bool {
	var max = 4
	if len(nums) < 4 {
		max = len(nums)
	} else {
		max = max - 1
	}

	for i := 0; i < max; i++ {
		if nums[i] == 9 {
			return true
		}
	}
	return false
}
