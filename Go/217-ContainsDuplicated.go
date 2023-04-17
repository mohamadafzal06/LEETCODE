package leetcode

func containsDuplicated(nums []int) bool {
	target := make(map[int]int, 0)

	for i := range nums {
		v, ok := target[nums[i]]
		if !ok {
			target[nums[i]]++
		} else if v > 1 {
			return true
		}
	}

	return false
}
