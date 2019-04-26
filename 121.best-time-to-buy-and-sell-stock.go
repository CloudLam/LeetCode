/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */
func maxProfit(prices []int) int {
    min := math.MaxInt32
    max := 0
    for i := 0; i < len(prices); i++ {
        if prices[i] < min {
            min = prices[i]
        }
        if prices[i] - min > max {
            max = prices[i] - min
        }
    }
    return max
}
