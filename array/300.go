package array

func lengthOfLIS(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(dp[i], res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
