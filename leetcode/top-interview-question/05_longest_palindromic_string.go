package leetcode

//https://leetcode.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {

	start := 0
	end := 0
	// two cases:
	// Case 1: result length is even
	// i, i+1 is center point
	// Case 2: result length is odd
	// i center point
	for i := 0; i < len(s); i++ {

		// Check case 1
		if i+1 < len(s) {
			ii := i
			jj := i + 1
			for ii >= 0 && jj < len(s) && s[ii] == s[jj] {
				if jj-ii > end-start {
					start = ii
					end = jj
				}
				ii--
				jj++
			}
		}

		// Check case 2

		ii := i
		jj := i

		for ii >= 0 && jj < len(s) && s[ii] == s[jj] {
			if jj-ii > end-start {
				start = ii
				end = jj
			}
			ii--
			jj++
		}

	}

	return s[start : end+1]
}

// Complexity Analysis
// Time Complexity: O(n^2)
// Space Complexity: O(length(s))
