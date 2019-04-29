/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */
func rotate(nums []int, k int) {
    if len(nums) != 0 {
        copy(nums, append(nums[len(nums) - k % len(nums):], nums[:len(nums) - k % len(nums)]...))
    }
}
