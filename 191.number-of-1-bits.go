/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	var result uint32 = 0

	for num != 0 {
		result += num % 2
		num /= 2
	}
	
	return int(result)
}
// @lc code=end
