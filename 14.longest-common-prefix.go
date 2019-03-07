/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (33.00%)
 * Total Accepted:    413.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 * 
 * If there is no common prefix, return an empty string "".
 * 
 * Example 1:
 * 
 * 
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 * 
 * 
 * Note:
 * 
 * All given inputs are in lowercase letters a-z.
 * 
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	} else if len(strs) > 1 {
		for i := len(strs[0]); i > 0; i-- {
			for j := 1; j < len(strs); j++ {
				if len(strs[j]) < i || strs[j][0:i] != strs[0][0:i] {
					break
				} else if j + 1 == len(strs) {
					return strs[0][0:i]
				}
			}
		}
	}
	return ""
}
