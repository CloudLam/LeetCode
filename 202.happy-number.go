/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	temp := make(map[int]bool)
	
    for n != 1 {
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		if temp[sum] {
			return false
		} else {
			temp[sum] = true
		}
		n = sum
	}
	return true
}
// @lc code=end
