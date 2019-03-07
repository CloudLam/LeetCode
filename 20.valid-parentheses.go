/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (35.90%)
 * Total Accepted:    521.9K
 * Total Submissions: 1.5M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * 
 * 
 * Note that an empty string is also considered valid.
 * 
 * Example 1:
 * 
 * 
 * Input: "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "(]"
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: "([)]"
 * Output: false
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: "{[]}"
 * Output: true
 * 
 * 
 */
func isValid(s string) bool {
    if (len(s) == 0) {
        return true
    }
    if (len(s) % 2 == 1) {
        return false
    }
    brackets := map[string]string {
        "(": ")",
        "[": "]",
        "{": "}",
    }
    stack := []byte{}
    stack = append(stack, s[0])
    for i := 1; i < len(s); i++ {
        last := len(stack) - 1
        if last < 0 {
            stack = append(stack, s[i])
            continue
        }
        if (brackets[string(stack[last])] == string(s[i])) {
            stack = append(stack[:last])
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}
