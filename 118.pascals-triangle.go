/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 *
 * https://leetcode.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (44.61%)
 * Total Accepted:    243.1K
 * Total Submissions: 535.2K
 * Testcase Example:  '5'
 *
 * Given a non-negative integer numRows, generate the first numRows of Pascal's
 * triangle.
 * 
 * 
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 * 
 * Example:
 * 
 * 
 * Input: 5
 * Output:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 * 
 * 
 */
func generate(numRows int) [][]int {
    var result [][]int
    for i := 0; i < numRows; i++ {
        var temp []int
        for j := 0; j < i + 1; j++ {
            if j == 0 || j == i {
                temp = append(temp, 1)
            } else {
                temp = append(temp, result[i-1][j-1] + result[i-1][j])
            }
        }
        result = append(result, temp)
    }
    return result
}
