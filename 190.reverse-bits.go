/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var result uint32

	for i := 0; i < 32; i++ {
		result *= 2
		if num != 0 {
			result += num % 2
			num /= 2
		}
	}

	return result
}
// @lc code=end
