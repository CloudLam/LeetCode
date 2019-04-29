/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */
func pow(x int, n int) int {
    result := 1
    for i := 0; i < n; i++ {
        result *= x
    }
    return result
}

func titleToNumber(s string) int {
    result := 0
    for i := len(s) - 1; i >= 0; i-- {
        result += (int(s[i]) - 64) * pow(26, len(s) - 1 - i)
    }
    return result
}
