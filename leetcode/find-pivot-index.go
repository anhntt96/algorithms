package leetcode

// https://leetcode.com/problems/find-pivot-index/
func pivotIndex(nums []int) int {
	leftSum := make([]int, len(nums))
	n := len(nums)
	for i, v := range nums {
		if i > 0 {
			leftSum[i] = leftSum[i-1]
		}
		leftSum[i] += v

	}
	sum := leftSum[n-1]

	for i := 0; i < len(nums); i++ {
		if leftSum[i]-nums[i] == sum-leftSum[i] {
			return i
		}
	}
	return -1
}
