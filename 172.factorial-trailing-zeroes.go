/*
 * @lc app=leetcode id=172 lang=golang
 *
 * [172] Factorial Trailing Zeroes
 */
func trailingZeroes(n int) int {
    result := 0
    for n > 0 {
        result += n / 5
        n /= 5
    }
    return result
}
