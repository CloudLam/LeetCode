/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (37.99%)
 * Total Accepted:    279.7K
 * Total Submissions: 734.4K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 * 
 * The input strings are both non-empty and contains only characters 1 or 0.
 * 
 * Example 1:
 * 
 * 
 * Input: a = "11", b = "1"
 * Output: "100"
 * 
 * Example 2:
 * 
 * 
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 * 
 */
func addBinary(a string, b string) string {
	lena, lenb := len(a), len(b)
	long, short := 0, 0
	temp, result := "", ""
	carry := 0
	if lena > lenb {
		long = lena
		short = lena - lenb
		temp = a[:short]
	} else {
		long = lenb
		short = lenb - lena
		temp = b[:short]
	}
	for i := 1; i <= long; i++ {
		if i > long - short {
			if carry == 1 {
				if string(temp[long-i]) == "1" {
					result = "0" + result
				} else {
					result = "1" + result
					if i < long {
						carry = 0
					}
				}
				temp = temp[:long-i]
			} else {
				result = temp + result
				break
			}
		} else {
			if carry == 1 {
				if a[lena-i] == b[lenb-i] {
					result = "1" + result
					if string(a[lena-i]) == "0" {
						carry = 0
					}
				} else {
					result = "0" + result
				}
			} else {
				if a[lena-i] == b[lenb-i] {
					result = "0" + result
					if string(a[lena-i]) == "1" {
						carry = 1
					}
				} else {
					result = "1" + result
				}
			}
		}
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}
