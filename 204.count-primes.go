/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

// @lc code=start
func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n % 6 != 1 && n % 6 != 5 {
		return false
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count += 1
		}
	}
	return count
}
// @lc code=end
