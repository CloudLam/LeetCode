/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	result := 0
	prev := 0


	for i := 0; i < len(nums); i++ {
		temp := result

		if nums[i] + prev > result {
			result = nums[i] + prev
		}
		
		prev = temp
	}
	
	return result
}
// @lc code=end
