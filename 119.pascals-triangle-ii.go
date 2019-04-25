/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (42.11%)
 * Total Accepted:    196.3K
 * Total Submissions: 456.8K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 * 
 * Note that the row index starts from 0.
 * 
 * 
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 * 
 * Example:
 * 
 * 
 * Input: 3
 * Output: [1,3,3,1]
 * 
 * 
 * Follow up:
 * 
 * Could you optimize your algorithm to use only O(k) extra space?
 * 
 */
func getRow(rowIndex int) []int {
    result := []int{1}
    for i := 1; i <= rowIndex; i++ {
        current := []int{1}
        for j := 0; j < len(result) - 1; j++ {
            current = append(current, result[j] + result[j+1])
        }
        current = append(current, 1)
        result = current
    }
    return result
}
