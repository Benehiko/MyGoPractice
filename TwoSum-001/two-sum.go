package TwoSum_001

func twoSum(nums []int, target int) []int {
	numsLen := len(nums) - 1
	for i := 0; i < numsLen; i++ {
		for y := i + 1; y < numsLen; y++ {
			if twoSumHelper(nums[i], nums[y]) == target {
				return []int{nums[i], nums[i+1]}
			}
		}
	}
	return nil
}

func twoSumHelper(i int, y int) (result int) {
	return i + y
}

func recursionTwoSum(nums []int, target int) []int {
	return recursion(nums, 0, target)
}

func recursion(nums []int, pos int, target int) []int {
	if pos > len(nums) {
		return nil
	}
	if nums[pos]+nums[pos+1] == target {
		return []int{nums[pos], nums[pos+1]}
	}
	return recursion(nums, pos+1, target)
}
