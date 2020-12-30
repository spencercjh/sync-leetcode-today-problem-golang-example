package leetcode

import "sort"

// https://leetcode-cn.com/problems/last-stone-weight
func lastStoneWeight(stones []int) int {
	if stones == nil || len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	for stones[len(stones)-2] != 0 {
		max := stones[len(stones)-1]
		smaller := stones[len(stones)-2]
		stones[len(stones)-1] = 0
		stones[len(stones)-2] = max - smaller
		sort.Ints(stones)
	}
	return stones[len(stones)-1]
}
