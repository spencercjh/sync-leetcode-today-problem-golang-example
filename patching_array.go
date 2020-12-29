package leetcode

// https://leetcode-cn.com/problems/patching-array
func minPatches(nums []int, n int) int {
	result := 0
	for max, index := 1, 0; max <= n; {
		if index < len(nums) && nums[index] <= max {
			max += nums[index]
			index++
		} else {
			max += max
			result++
		}
	}
	return result
}
