package leetcode

// https://leetcode-cn.com/problems/can-place-flowers
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; {
		if flowerbed[i] == 1 {
			i += 2
		} else if flowerbed[i] == 0 && (i+1 >= len(flowerbed) || flowerbed[i+1] == 0) {
			i += 2
			n--
		} else {
			i++
		}
	}
	return n == 0
}
