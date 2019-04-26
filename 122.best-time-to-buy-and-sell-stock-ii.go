/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */
func maxProfit(prices []int) int {
    max := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            max += prices[i] - prices[i-1]
        }
    }
    return max
}
