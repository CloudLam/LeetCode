/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */
func convertToTitle(n int) string {
    if n == 0 {
        return ""
    }
    var result string
    for n > 0 {
        if n % 26 == 0 {
            result = "Z" + result
        } else {
            result = string(n % 26 + 64) + result
        }
        n = (n - 1) / 26
    }
    return result
}
