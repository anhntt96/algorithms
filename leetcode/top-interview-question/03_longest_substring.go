package leetcode

//https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	result := 0

	// position of char in array, default -1
	pos := make([]int, 256)
	for i := 0; i < 256; i++ {
		pos[i] = -1
	}
	// i j represent substring s[i:j]
	i := 0
	j := 0

	for j <= len(s) {
		if j-i > result {
			result = j - i
		}
		if j == len(s) {
			break
		}
		if pos[s[j]] == -1 {
			pos[s[j]] = j
		} else {
			if pos[s[j]]+1 > i {
				i = pos[s[j]] + 1
			}
			pos[s[j]] = j
		}
		j++
	}
	return result
}

// Complexity Analysis
// Time Complexity: O(n)
// Space Complexity: O(256)
