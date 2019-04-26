/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */
func abs(a byte, b byte) byte {
    if a > b {
        return a - b
    } else {
        return b - a
    }
}

func isDigit(x byte) bool {
    return x > 47 && x < 58
}

func isLetter(x byte) bool {
    return (x > 64 && x < 91) || (x > 96 && x < 123)
}

func isPalindrome(s string) bool {
    for i, j := 0, len(s) - 1; i < j; {
        if !isDigit(s[i]) && !isLetter(s[i]) {
            i++
            continue
        }
        if !isDigit(s[j]) && !isLetter(s[j]) {
            j--
            continue
        }
        if i < j {
            if (isDigit(s[i]) || isDigit(s[j])) && s[i] != s[j] {
                return false
            } else if s[i] != s[j] && abs(s[j], s[i]) != 32 {
                return false
            }
            i++
            j--
        }
    }
    return true
}
