/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */
func singleNumber(nums []int) int {
    result := 0
    for i := 0; i < len(nums); i++ {
        result ^= nums[i]
    }
    return result
}
